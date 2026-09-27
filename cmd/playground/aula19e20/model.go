package main

type Cliente struct {
	Id        string `db:"id" fake:"skip"`
	NomeRazao string `db:"nome_razao" fake:"{firstname}"`
}
