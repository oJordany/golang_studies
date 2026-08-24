package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

func Ptr[T any](v T) *T {
	return &v
}

func main() {
	var notas Notas
	var nota Nota
	nota.DtaCompra = time.Now()
	nota.TotalCompra = 150.75
	nota.Numero = 12345
	nota.XmlAttr = "entrada"
	nota.Cliente = Ptr("XYZ")
	nota.Itens = []Item{
		{Descricao: "Item A", Quantidade: 2, ValorUnitario: 25.00, Observacao: Ptr("teste")},
		{Descricao: "Item B", Quantidade: 1, ValorUnitario: 100.75},
	}

	var item Item
	item.Descricao = "Item C"
	item.Quantidade = 3
	item.ValorUnitario = 50.00

	nota.Itens = append(nota.Itens, item)
	nota.refresh()
	notas.Notas = append(notas.Notas, nota)
	notas.Notas = append(notas.Notas, nota)

	// ==================== JSON =====================

	fmt.Println("Modelo 1:")
	BodyJson1, err1 := json.Marshal(nota)
	if err1 != nil {
		panic(err1)
	}
	println("JSON Byte Array:", string(BodyJson1))
	fmt.Println("JSON len:", len(BodyJson1))

	fmt.Println("Modelo 2:")
	BodyJson2, err := json.MarshalIndent(nota, "", " ")
	if err != nil {
		panic(err)
	}
	println("JSON Byte Array:", string(BodyJson2))
	fmt.Println("JSON len:", len(BodyJson2))

	fmt.Println("Modelo 3:")
	BodyJson3, err1 := json.MarshalIndent(notas, "", "	")
	if err1 != nil {
		panic(err1)
	}
	println("JSON Byte Array:", string(BodyJson3))
	fmt.Println("JSON len:", len(BodyJson3))

	// ==================== XML =====================

	fmt.Println("Modelo 1:")
	BodyXML1, err1 := xml.Marshal(nota)
	if err1 != nil {
		panic(err1)
	}
	println("JSON Byte Array:", string(BodyXML1))
	fmt.Println("JSON len:", len(BodyXML1))

	fmt.Println("Modelo 2:")
	BodyXML2, err := xml.MarshalIndent(nota, "", " ")
	if err != nil {
		panic(err)
	}
	println("JSON Byte Array:", string(BodyXML2))
	fmt.Println("JSON len:", len(BodyXML2))

	fmt.Println("Modelo 3:")
	BodyXML3, err1 := xml.MarshalIndent(notas, "", "	")
	if err1 != nil {
		panic(err1)
	}
	println("JSON Byte Array:", string(BodyXML3))
	fmt.Println("JSON len:", len(BodyXML3))
	// var notas Notas
	// notas.Notas = append(notas.Notas, nota)

}
