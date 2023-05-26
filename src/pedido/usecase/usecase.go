package usecase

import (
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/pedido/repository"
)

type Usecase interface {
}

type usecasePedido struct {
	repository repository.Repository
}

var _ Usecase = (*usecasePedido)(nil)

func NewUsecase(repo repository.Repository) Usecase {
	return &usecasePedido{
		repository: repo,
	}
}

func (u *usecasePedido) NewOrder(itens *models.NovoPedido) (*models.PedidoResponse, error) {
	var valorTotal float64

	for _, produto := range *itens.Produtos {

	}
	return nil, nil
}
