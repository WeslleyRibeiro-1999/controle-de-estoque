package api

import (
	"net/http"
	"strconv"

	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/fornecedor/repository"
	"github.com/labstack/echo/v4"
)

type HttpFornecedor interface {
	CreateFornecedor(c echo.Context) error
	GetAllFornecedor(c echo.Context) error
	GetOne(c echo.Context) error
}

type httpFornecedor struct {
	repository repository.Repository
}

var _ HttpFornecedor = (*httpFornecedor)(nil)

func NewHandler(repo repository.Repository) HttpFornecedor {
	return &httpFornecedor{
		repository: repo,
	}
}

func (h *httpFornecedor) CreateFornecedor(c echo.Context) error {
	var req *models.Fornecedor

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	_, err := h.repository.Create(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (h *httpFornecedor) GetAllFornecedor(c echo.Context) error {
	fornecedor, err := h.repository.FindAll()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, fornecedor)
}

func (h *httpFornecedor) GetOne(c echo.Context) error {
	var fornecedor models.Fornecedor

	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	fornecedor.ID = int32(id)

	res, err := h.repository.FindOne(&fornecedor)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
