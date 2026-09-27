package main

import (
	"database/sql"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/sijms/go-ora/v2"
)

type Cliente struct {
	Id   string `fake:"skip"`
	Nome string `fake:"{firstname}"`
}

const OraUser = "NEXUS"
const OraPass = "master"
const OraServer = "192.168.1.6"
const OraPort = 1521
const OraDb = "XEPDB1"
const Url = "oracle://%s:%s@%s:%d/%s"

func hasClientes(conn *sql.DB) bool {
	var result int
	err := conn.QueryRow(`SELECT COUNT(*) FROM cliente`).Scan(&result)
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
	rows, err := conn.Query("SELECT RAWTOHEX(id), nome_razao FROM cliente")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Cliente
		err := rows.Scan(&c.Id, &c.Nome)
		if err != nil {
			panic(err)
		}
		clientes = append(clientes, c)
	}

	for _, cliente := range clientes {
		fmt.Printf("ID: %s, Nome/Razao: %s\n", cliente.Id, cliente.Nome)
	}
}

func main() {
	DSN := fmt.Sprintf(Url, OraUser, OraPass, OraServer, OraPort, OraDb)

	conn, err := sql.Open("oracle", DSN)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	query := `
	BEGIN
		EXECUTE IMMEDIATE '
			CREATE TABLE cliente (
				id RAW(16) DEFAULT SYS_GUID() PRIMARY KEY,
				nome_razao VARCHAR2(255) NOT NULL,
				criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			)';
	EXCEPTION
		WHEN OTHERS THEN
			IF SQLCODE != -955 THEN
				RAISE;
			END IF;
	END;`

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
	rows, err := conn.Query(`SELECT RAWTOHEX(id), nome_razao FROM cliente`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Após alteração dos nomes: ")
	for rows.Next() {
		var c Cliente
		err := rows.Scan(&c.Id, &c.Nome)
		if err != nil {
			panic(err)
		}
		gofakeit.Struct(&c)
		query := "UPDATE cliente SET nome_razao=:1 WHERE id=HEXTORAW(:2)"
		_, err = conn.Exec(query, c.Nome, c.Id)
		if err != nil {
			panic(err)
		}
	}
	showClientes(conn)
}
