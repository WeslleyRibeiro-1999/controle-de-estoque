package models

type Produto struct {
	ID        int32   `json:"id"`
	Name      string  `json:"nome" gorm:"unique"`
	Descricao string  `json:"descricao"`
	Qtde      float64 `json:"quantidade"`
	Value     float64 `json:"valor"`
}
