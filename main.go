package main

import (
	"log"

	"github.com/WeslleyRibeiro-1999/controle-de-estoque/database"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/api"
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/src/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/adega?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := database.NewDatabase(dsn, []interface{}{&models.Produto{}})
	if err != nil {
		log.Fatalf("failed to connect database: %+v", err)
	}

	repo := repository.NewRepository(db)

	e := echo.New()
	e.Use(middleware.CORS())

	api := api.NewHandler(repo)

	e.POST("/produto", api.CreateProduct)
	e.GET("/produtos", api.GetAllProducts)
	e.GET("/produto/:id", api.GetOne)

	e.Logger.Fatal(e.Start(":8080"))
}
