package main

import (
	"log"

	"github.com/WeslleyRibeiro-1999/controle-de-estoque/database"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	apiFornecedor "github.com/WeslleyRibeiro-1999/controle-de-estoque/src/fornecedor/api"
	repoFornecedor "github.com/WeslleyRibeiro-1999/controle-de-estoque/src/fornecedor/repository"
	apiPedido "github.com/WeslleyRibeiro-1999/controle-de-estoque/src/pedido/api"
	repoPedido "github.com/WeslleyRibeiro-1999/controle-de-estoque/src/pedido/repository"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/pedido/usecase"
	apiProdutos "github.com/WeslleyRibeiro-1999/controle-de-estoque/src/produto/api"
	repoProduto "github.com/WeslleyRibeiro-1999/controle-de-estoque/src/produto/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/adega?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := database.NewDatabase(dsn, []interface{}{&models.Produto{}, &models.Fornecedor{}, &models.Pedido{}, &models.ProdutosPedido{}})
	if err != nil {
		log.Fatalf("failed to connect database: %+v", err)
	}

	repoProd := repoProduto.NewRepository(db)
	repoForn := repoFornecedor.NewRepository(db)
	repoPed := repoPedido.NewRepository(db)

	usecase := usecase.NewUsecase(repoPed, repoProd)

	e := echo.New()
	e.Use(middleware.CORS())

	produto := apiProdutos.NewHandler(repoProd)
	fornecedor := apiFornecedor.NewHandler(repoForn)
	pedido := apiPedido.NewHandler(usecase)

	e.POST("/produto", produto.CreateProduct)
	e.GET("/produtos", produto.GetAllProducts)
	e.GET("/produto/:id", produto.GetOne)

	e.POST("/fornecedor", fornecedor.CreateFornecedor)
	e.GET("/fornecedores", fornecedor.GetAllFornecedor)
	e.GET("/fornecedor/:id", fornecedor.GetOne)

	e.POST("/pedidos", pedido.NewOrder)

	e.Logger.Fatal(e.Start(":8080"))
}
