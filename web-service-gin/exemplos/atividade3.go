package exemplos

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Hello(context *gin.Context) {
	nome := context.Query("nome")
	sobrenome := context.Query("sobrenome")
	message := fmt.Sprintf("Olá, %s %s", nome, sobrenome)

	context.JSON(http.StatusOK, 
		gin.H{
			"message": message,
		},
	)
}

func GetAll(context *gin.Context) {
	var produtos []Product = []Product{
		{
			Id: 1,
			Nome: "Ventilador",
			Preco: 150.56,
		},
		{
			Id: 1,
			Nome: "Sofa",
			Preco: 15000.00,
		},
		{
			Id: 1,
			Nome: "Geladeira",
			Preco: 3000,
		},
		{
			Id: 1,
			Nome: "Pneus Pirelli",
			Preco: 650.99,
		},
		{
			Id: 1,
			Nome: "Lampada",
			Preco: 18.99,
		},
		{
			Id: 1,
			Nome: "Baki Coral",
			Preco: 2856.45,
		},
	}

	produtosJson, err := json.Marshal(produtos)
	if err != nil {
		context.JSON(
			http.StatusOK,
			gin.H{
				"error": fmt.Sprintf("%s", err),
			},
		)
		return 	
	}
	
	fmt.Println(produtosJson)

	context.IndentedJSON(http.StatusOK, gin.H{ "data": produtos })
}