package usecase

import (
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/pedido/repository"
	repoProduto "github.com/WeslleyRibeiro-1999/controle-de-estoque/src/produto/repository"
)

type Usecase interface {
	NewOrder(itens *models.NovoPedido) (*models.PedidoResponse, error)
}

type usecasePedido struct {
	repository  repository.Repository
	repoProduto repoProduto.Repository
}

var _ Usecase = (*usecasePedido)(nil)

func NewUsecase(repo repository.Repository, repoProd repoProduto.Repository) Usecase {
	return &usecasePedido{
		repository:  repo,
		repoProduto: repoProd,
	}
}

func (u *usecasePedido) NewOrder(itens *models.NovoPedido) (*models.PedidoResponse, error) {
	var valorTotal float64
	pedidoResponse := models.PedidoResponse{
		Produtos:   []models.ProdutoResponse{},
		ValorTotal: 0,
	}

	pedido, err := u.repository.NewOrder(&models.Pedido{})
	if err != nil {
		return nil, err
	}

	for _, produtoRequest := range *itens.Produtos {
		produto, err := u.repoProduto.FindOne(&models.Produto{ID: produtoRequest.ID})
		if err != nil {
			return nil, err
		}

		_, err = u.repository.NewProductOrder(&models.ProdutosPedido{
			PedidoID:  pedido.ID,
			ProdutoID: produto.ID,
			Qtde:      int64(produtoRequest.Qtde),
		})
		if err != nil {
			return nil, err
		}

		response := models.ProdutoResponse{Nome: produto.Name, Qtde: produtoRequest.Qtde}

		pedidoResponse.Produtos = append(pedidoResponse.Produtos, response)

		valorTotal += produto.Value * produtoRequest.Qtde
	}

	_, err = u.repository.UpdateOrder(&models.Pedido{ID: pedido.ID, ValorTotal: valorTotal})
	if err != nil {
		return nil, err
	}

	pedidoResponse.ValorTotal = valorTotal

	return &pedidoResponse, nil
}
