package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type Pessoa struct {
	Id   *string `json:"id" xml:"id"`
	Nome string  `json:"nome" xml:"nome"`
	Cpf  string  `json:"cpf" xml:"cpf"`
}

var Pessoas []Pessoa

const key = "my-secret-api-key"

func Guid() string {
	id := uuid.New().String()
	return id
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

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/pessoas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if Pessoas == nil || len(Pessoas) == 0 {
			http.Error(w, "No people found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(Pessoas)
	})

	r.Group(func(r chi.Router) {
		r.Use(apiKeyMiddleware)

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

	r.Get("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
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
