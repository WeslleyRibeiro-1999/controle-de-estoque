package api

import (
	"net/http"

	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/pedido/repository"
	"github.com/labstack/echo/v4"
)

type HttpPedido interface {
	NewOrder(c echo.Context) error
}

type httpPedido struct {
	repository repository.Repository
}

var _ HttpPedido = (*httpPedido)(nil)

func NewHandler(repo repository.Repository) HttpPedido {
	return &httpPedido{
		repository: repo,
	}
}

func (h *httpPedido) NewOrder(c echo.Context) error {
	var req *models.Pedido

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	pedido, err := h.repository.NewOrder(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	// produtos, err := h.repository.NewProductOrder(&models.ProdutosPedido{})

	return c.JSON(http.StatusCreated, pedido)
}
