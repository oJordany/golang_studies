package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	dataset "github.com/fitlcarlos/godataset"
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

func tableExists(ds *dataset.DataSet, tableName string) bool {
	err := ds.AddSql("SELECT COUNT(*) FROM user_tables WHERE table_name = '" + tableName + "'").Open()
	if err != nil {
		panic(err)
	}
	defer ds.Close()

	return ds.FieldByName("count(*)").AsInt() > 0
}

func hasClientes(ds *dataset.DataSet) bool {
	err := ds.AddSql(`SELECT COUNT(*) FROM cliente`).Open()
	if err != nil {
		panic(err)
	}
	defer ds.Close()

	result := ds.FieldByName("count(*)").AsInt()
	fmt.Println(result)
	return result != 0
}

func showClientes(ds *dataset.DataSet) {
	err := ds.AddSql("SELECT RAWTOHEX(id) AS id, nome_razao AS nome FROM cliente").Open()
	if err != nil {
		panic(err)
	}
	defer ds.Close()

	for !ds.Eof() {
		fmt.Println("ID:", ds.FieldByName("id").AsString(),
			"Nome", ds.FieldByName("nome").AsString())
		ds.Next()
	}
}

func main() {
	DSN := fmt.Sprintf(Url, OraUser, OraPass, OraServer, OraPort, OraDb)

	conn, err := dataset.NewConnection(dataset.DialectType(dataset.ORACLE), DSN)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	Ds := dataset.NewDataSet(conn)
	defer Ds.Free()

	if !tableExists(Ds, "CLIENTE") {
		query := `
		CREATE TABLE cliente (
			id RAW(16) DEFAULT SYS_GUID() PRIMARY KEY,
			nome_razao VARCHAR2(255) NOT NULL,
			criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

		_, err = Ds.AddSql(query).Exec()
		if err != nil {
			panic(err)
		}
		Ds.Close()
	}
	if !hasClientes(Ds) {
		r, err := Ds.AddSql("INSERT INTO cliente (nome_razao) VALUES ('jordany')").Exec()
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
		Ds.Close()
	}
	showClientes(Ds)

	err = Ds.AddSql(`SELECT RAWTOHEX(id) AS id, nome_razao AS nome FROM cliente`).Open()
	if err != nil {
		panic(err)
	}

	DsUpdate := dataset.NewDataSet(conn)
	defer DsUpdate.Free()

	fmt.Println("Após alteração dos nomes: ")
	for !Ds.Eof() {
		var c Cliente
		err := Ds.ToStruct(&c)
		if err != nil {
			panic(err)
		}
		gofakeit.Struct(&c)

		_, err = DsUpdate.AddSql(`UPDATE cliente SET nome_razao = :nome WHERE id = HEXTORAW(:id)`).
			SetInputParam("nome", c.Nome).
			SetInputParam("id", c.Id).
			Exec()
		if err != nil {
			panic(err)
		}
		DsUpdate.Close()

		Ds.Next()
	}
	Ds.Close()

	showClientes(Ds)
}
