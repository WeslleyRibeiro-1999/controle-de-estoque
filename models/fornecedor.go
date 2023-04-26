package models

type Fornecedor struct {
	ID             int32  `json:"id"`
	NomeFornecedor string `json:"nome_fornecedor"`
	Endereco       string `json:"endereco"`
	Numero         string `json:"numero"`
}
