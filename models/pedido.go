package models

type Pedido struct {
	ID         int32   `json:"id"`
	ValorTotal float64 `json:"valor_total"`
}

type ProdutosPedido struct {
	ID        int32 `json:"id"`
	PedidoID  int32 `json:"pedido_id" gorm:"foreignKey:ID"`
	ProdutoID int32 `json:"produto_id"`
	Qtde      int64 `json:"qtde"`
}

type NovoPedido struct {
	Produtos *[]ProdutoRequest `json:"produtos"`
}

type ProdutoRequest struct {
	ID   int32   `json:"id"`
	Qtde float64 `json:"quantidade"`
}

type ProdutoResponse struct {
	Nome string  `json:"nome"`
	Qtde float64 `json:"quantidade"`
}

type PedidoResponse struct {
	Produtos   []ProdutoResponse `json:"produtos"`
	ValorTotal float64           `json:"valor_total"`
}
