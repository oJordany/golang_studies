package main

import (
	"database/sql"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const PgUser = "postgres"
const PgPass = "master"
const PgServer = "localhost"
const PgPort = 5432
const PgDb = "postgres"

const Url = "postgres://%s:%s@%s:%d/%s"

func hasClientes(conn *sql.DB) bool {
	var result int
	err := conn.QueryRow(`SELECT COUNT(*) FROM public.cliente`).Scan(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	if result != 0 {
		return true
	}
	return false
}

func showClientes(conn *sql.DB) {
	var clientes []Cliente
	rows, err := conn.Query("SELECT id, nome_razao FROM public.cliente")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Cliente
		err := rows.Scan(&c.Id, &c.NomeRazao)
		if err != nil {
			panic(err)
		}
		clientes = append(clientes, c)
	}

	for _, cliente := range clientes {
		fmt.Printf("ID: %s, Nome/Razao: %s\n", cliente.Id, cliente.NomeRazao)
	}
}

func main() {
	DSN := fmt.Sprintf(Url, PgUser, PgPass, PgServer, PgPort, PgDb)
	conn, err := sql.Open("pgx", DSN)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	query := `
	CREATE TABLE IF NOT EXISTS cliente (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		nome_razao VARCHAR(255) NOT NULL,
		criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = conn.Exec(query)
	if err != nil {
		panic(err)
	}
	if !hasClientes(conn) {
		var r sql.Result
		r, err = conn.Exec("INSERT INTO cliente (nome_razao) VALUES ('jordany')")
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	}

	showClientes(conn)
	rows, err := conn.Query(`SELECT id, nome_razao FROM public.cliente`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Após alteração dos nomes: ")
	for rows.Next() {
		var c Cliente
		err := rows.Scan(&c.Id, &c.NomeRazao)
		if err != nil {
			panic(err)
		}
		gofakeit.Struct(&c)
		query := "UPDATE cliente SET nome_razao=$1 WHERE id=$2"
		_, err = conn.Exec(query, c.NomeRazao, c.Id)
		if err != nil {
			panic(err)
		}
	}
	showClientes(conn)
}
