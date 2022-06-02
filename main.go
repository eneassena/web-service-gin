package main

import (
 
	"web-service-gin/exercicios"

	"github.com/gin-gonic/gin"
)




 
func main() {
	router := gin.Default()
	
	group := router.Group("/exercicios/produtos")
	group.GET("/", exercicios.Exercicio2RetoandoJSON)
	group.GET("/all", exercicios.Exercicio3ListagemProdutos)
	group.GET("/filter", exercicios.FilterProducts)
	group.POST("/create", exercicios.CreateProduto)

	router.Run()
	
}

