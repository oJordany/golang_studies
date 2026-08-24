package main

import (
	"encoding/json"
	"encoding/xml"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Pessoa struct {
	Nome  string `json:nome xml:"nome"`
	Idade int    `json:idade xml:"idade"`
}

type PublicPessoa struct {
	Id    int    `json:id xml:id`
	Nome  string `json:nome xml:"nome"`
	Idade int    `json:idade xml:"idade"`
}

func main() {

	r := chi.NewRouter()

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		pessoa := Pessoa{Nome: "João", Idade: 30}
		jsonData, _ := json.Marshal(pessoa)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	r.Get("/xml", func(w http.ResponseWriter, r *http.Request) {
		pessoa := Pessoa{Nome: "João", Idade: 30}
		xmlData, _ := xml.Marshal(pessoa)
		w.Header().Set("Content-Type", "application/xml")
		w.Write(xmlData)
	})

	r.Post("/pessoa", func(w http.ResponseWriter, r *http.Request) {
		var p Pessoa
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var pp PublicPessoa
		pp.Id = rand.Int()
		pp.Idade = p.Idade
		pp.Nome = p.Nome
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pp)
	})

	r.Put("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
		IdStr := chi.URLParam(r, "id")
		Id, strErr := strconv.Atoi(IdStr)
		if strErr != nil {
			http.Error(w, strErr.Error(), http.StatusBadRequest)
			return
		}
		var p Pessoa
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		demitido := r.URL.Query().Get("demitido") // Para Query params
		if demitido == "" {
			http.Error(w, "Query param 'demitido' is required", http.StatusBadRequest)
			return
		}

		HeaderApiKey := r.Header.Get("x-api-key")
		if HeaderApiKey == "" {
			http.Error(w, "Header 'x-api-key' is required", http.StatusBadRequest)
			return
		}

		if HeaderApiKey != "cursogolang" {
			http.Error(w, "Invalid 'x-api-key'", http.StatusUnauthorized)
			return
		}

		var pp PublicPessoa
		pp.Id = Id
		pp.Idade = p.Idade + 1
		pp.Nome = p.Nome + " Atualizado d:" + demitido
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pp)
	})

	http.ListenAndServe(":7001", r)
}
