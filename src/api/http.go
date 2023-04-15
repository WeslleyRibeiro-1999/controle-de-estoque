package api

import (
	"net/http"
	"strconv"

	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/repository"
	"github.com/labstack/echo/v4"
)

type HttpProduto interface {
	CreateProduct(c echo.Context) error
	GetAllProducts(c echo.Context) error
	GetOne(c echo.Context) error
}

type httpProduto struct {
	repository repository.Repository
}

var _ HttpProduto = (*httpProduto)(nil)

func NewHandler(repo repository.Repository) HttpProduto {
	return &httpProduto{
		repository: repo,
	}
}

func (h *httpProduto) CreateProduct(c echo.Context) error {
	var req *models.Produto

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	_, err := h.repository.Create(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (h *httpProduto) GetAllProducts(c echo.Context) error {
	products, err := h.repository.FindAll()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, products)
}

func (h *httpProduto) GetOne(c echo.Context) error {
	var product models.Produto

	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	product.ID = int32(id)

	res, err := h.repository.FindOne(&product)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
