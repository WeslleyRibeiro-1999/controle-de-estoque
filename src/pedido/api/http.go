package api

import (
	"net/http"

	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/pedido/repository"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/pedido/usecase"
	"github.com/labstack/echo/v4"
)

type HttpPedido interface {
	NewOrder(c echo.Context) error
	GetAll(c echo.Context) error
}

type httpPedido struct {
	usecase    usecase.Usecase
	repository repository.Repository
}

var _ HttpPedido = (*httpPedido)(nil)

func NewHandler(usecase usecase.Usecase, repository repository.Repository) HttpPedido {
	return &httpPedido{
		usecase:    usecase,
		repository: repository,
	}
}

func (h *httpPedido) NewOrder(c echo.Context) error {
	var req *models.NovoPedido

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := h.usecase.NewOrder(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *httpPedido) GetAll(c echo.Context) error {
	pedidos, err := h.repository.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, pedidos)
}
