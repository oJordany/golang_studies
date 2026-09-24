package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Pessoa struct {
	Id   *string `json:"id" xml:"id"`
	Nome string  `json:"nome" xml:"nome"`
	Cpf  string  `json:"cpf" xml:"cpf"`
}

var Pessoas []Pessoa

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type OAuth2LoginRequest struct {
	ClientID     string `json:"client_id`
	ClientSecret string `json:"client_secret`
	GrantType    string `json:grant_type`
}

type OAuthLoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

const (
	key      = "my-secret-api-key"
	username = "admin"
	password = "password"

	jwtSecret = "my-secret-jwt-secret"

	clientID     = "meu-client-id"
	clientSecret = "meu-client-secret"
)

func Guid() string {
	id := uuid.New().String()
	return id
}

func GetDatasFromFile() {
	file, err := os.Open("pessoas.json")
	if err != nil {
		if os.IsNotExist(err) {
			Pessoas = []Pessoa{}
			return
		}
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&Pessoas)
	if err != nil {
		panic(err)
	}
}

func SaveDatasToFile() {
	file, err := os.Create("pessoas.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(Pessoas)
	if err != nil {
		panic(err)
	}
}

func GetPessoaById(id string) *Pessoa {
	for _, pessoa := range Pessoas {
		if *pessoa.Id == id {
			return &pessoa
		}
	}
	return nil
}

func GetPessoaByCpf(cpf string) *Pessoa {
	for _, pessoa := range Pessoas {
		if pessoa.Cpf == cpf {
			return &pessoa
		}
	}
	return nil
}

func GetPessoaByCpfNotId(cpf string, id string) *Pessoa {
	for _, pessoa := range Pessoas {
		if pessoa.Cpf == cpf && *pessoa.Id != id {
			return &pessoa
		}
	}
	return nil
}

// Autenticação api-key
func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != key {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func saveMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		SaveDatasToFile()
	})
}

// Autenticação JWT
func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "token ausente", http.StatusUnauthorized)
			return
		}
		tokenString, ok := strings.CutPrefix(authHeader, "Bearer ")
		if !ok {
			http.Error(w, "formato de token inválido", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Metodo de assinatura inválido")
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Autenticação OAuth2
func oauthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer") {
			http.Error(w, "Token ausente", http.StatusUnauthorized)
			return
		}
		tokenString, _ := strings.CutPrefix(auth, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Método de assinatura inválido")
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// login jwt
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	claims := jwt.MapClaims{
		"username": loginReq.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(LoginResponse{Token: tokenString})
}

// login OAuth2
func oauthLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar form-data", http.StatusBadRequest)
		return
	}
	clientIDForm := r.FormValue("client_id")
	clientSecretForm := r.FormValue("client_secret")
	grantTypeForm := r.FormValue("grant_type")

	if grantTypeForm != "client_credentials" {
		http.Error(w, "grant_type inválido ou ausente", http.StatusBadRequest)
		return
	}

	if clientIDForm != clientID || clientSecretForm != clientSecret {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"client_id": clientIDForm,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}
	resp := OAuthLoginResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	GetDatasFromFile()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// r.Post("/login", loginHandler)
	r.Post("/login", oauthLoginHandler)

	r.With(oauthMiddleware).Get("/pessoas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if Pessoas == nil || len(Pessoas) == 0 {
			http.Error(w, "No people found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(Pessoas)
	})

	r.Group(func(r chi.Router) {
		// r.Use(apiKeyMiddleware)
		r.Use(oauthMiddleware)
		r.Use(saveMiddleware)

		r.Post("/pessoa", func(w http.ResponseWriter, r *http.Request) {
			var p Pessoa
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if GetPessoaByCpf(p.Cpf) != nil {
				http.Error(w, "CPF already exists", http.StatusConflict)
				return
			}

			id := Guid()
			p.Id = &id

			Pessoas = append(Pessoas, p)
			w.Header().Set("Content-Type", "application/json")
			SaveDatasToFile()
			json.NewEncoder(w).Encode(p)
		})

		r.Put("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			id := chi.URLParam(r, "id")

			var p Pessoa
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			existingPessoa := GetPessoaById(id)
			if existingPessoa == nil {
				http.Error(w, "Pessoa not found", http.StatusNotFound)
				return
			}

			pessoaWithSameCpfModified := GetPessoaByCpfNotId(p.Cpf, id)
			if p.Cpf != existingPessoa.Cpf && pessoaWithSameCpfModified != nil {
				http.Error(w, "CPF already exists", http.StatusConflict)
				return
			}

			existingPessoa.Nome = p.Nome
			existingPessoa.Cpf = p.Cpf

			json.NewEncoder(w).Encode(existingPessoa)
		})

		r.Delete("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			id := chi.URLParam(r, "id")

			existingPessoa := GetPessoaById(id)
			if existingPessoa == nil {
				http.Error(w, "Pessoa not found", http.StatusNotFound)
				return
			}

			for i, pessoa := range Pessoas {
				if *pessoa.Id == id {
					Pessoas = append(Pessoas[:i], Pessoas[i+1:]...)
					break
				}
			}

			w.WriteHeader(http.StatusNoContent)
		})
	})

	r.With(oauthMiddleware).Get("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")

		existingPessoa := GetPessoaById(id)
		if existingPessoa == nil {
			http.Error(w, "Pessoa not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(existingPessoa)
	})

	http.ListenAndServe(":7001", r)
}
