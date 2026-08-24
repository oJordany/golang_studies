package main

import "time"

type Item struct {
	Descricao     string  `json:"descricao" xml:"descricao"`
	Quantidade    int     `json:"quantidade" xml:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario" xml:"valor_unitario"`
	Observacao    *string `json:"observacao,omitempty" xml:"observacao,omitempty"`
}

type Nota struct {
	XmlAttr     string    `json:"-" xml:"tipo,attr"`
	DtaCompra   time.Time `json:"dta_compra" xml:"dta_compra"`
	TotalCompra float64   `json:"total_nota" xml:"total_nota"`
	Numero      int       `json:"numero" xml:"numero"`
	Cliente     *string   `json:"cliente" xml:"cliente"`
	Logradouro  *string   `json:"logradouro,omitempty" xml:"logradouro,omitempty"`
	Itens       []Item    `json:"itens" xml:"itens"`
}

func (nota *Nota) refresh() {
	nota.TotalCompra = 0
	for _, item := range nota.Itens {
		nota.TotalCompra += float64(item.Quantidade) * item.ValorUnitario
	}
}

type Notas struct {
	Notas []Nota `json:"notas" xml:"notas"`
}
