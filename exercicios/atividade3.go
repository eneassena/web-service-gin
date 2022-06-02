package exercicios

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

 
var produtos []Produto

var lastID int


func CreateProduto(context *gin.Context) {
	
	// questão 3 validar token
	token := context.Request.Header.Get("token") 	
	if token != "123456" {
		context.JSON(http.StatusUnauthorized, gin.H{ "errir": "você não tem permissão para fazer esta solicitação" })
		return
	}

	var entites Produto

	// questão 1 criar Entidade
	err := context.ShouldBindJSON(&entites)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{ "error": err })
		return
	}

	// questão 2 validação de campo
	if entites.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{ "error": "o campo nome é obrigatório" })
		return 
	}

	if entites.Count == 0 {
		context.JSON(http.StatusBadRequest, gin.H{ "error": "o campo count é obrigatório" })
		return 
	}

	if entites.Type == "" {
		context.JSON(http.StatusBadRequest, gin.H{ "error": "o campo type é obrigatório" })
		return 
	}

	if entites.Price == 0 {
		context.JSON(http.StatusBadRequest, gin.H{ "error": "o campo price é obrigatório" })
		return 
	}

	lastID ++
	entites.ID = lastID
	produtos = append(produtos, entites)

	context.JSON(200, gin.H{ "last_id": lastID })

	fmt.Println(produtos)
}
