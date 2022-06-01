package main

// thunder client é uma exetnsão para teste de api

import (
	 
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)
  
type Produtos struct {
	Id int 
	Nome string
	Preco float64
}
var produtos []Produtos = []Produtos{}

var lastID int

 

func add(context *gin.Context) {
	
	var produto Produtos 
	
	err := context.ShouldBindJSON(&produto) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ 
			"error": err,
		})
		return 
	}
	
	lastID++

	
	produto.Id = lastID
	produtos = append(produtos, produto)	
	fmt.Println(produto)
	/*p := Produtos{
		Id: lastID,
		Nome: "Teste",
		Preco: 100,
	}*/
	//produtos = append(produtos, p)

	context.IndentedJSON(200, gin.H{
		"msg": "add produtos",
		"produto": produto,
	})
}

func listProdutos(context *gin.Context) {

	context.JSON(200, gin.H{ 
		"data": produtos,
	})
}


func filterProdutos(context *gin.Context){ 
	nome := context.Query("nome")
	
	if strings.TrimSpace(nome) == "" {
		context.JSON(http.StatusNotFound, gin.H{ "error": "paramento nome é obrigatório" })
		return
	}

	detalhes := ""

	for _, value := range produtos {	
		if value.Nome == nome {
			precoStr := fmt.Sprintf("%.2f", value.Preco)
			idStr := string(value.Id)
			detalhes = strings.Join([]string{idStr, value.Nome, precoStr}, " - ") 
		}
	}

	if detalhes != "" {
		context.JSON(http.StatusOK, gin.H{"produto": detalhes})
		return
	} else {
		context.JSON(http.StatusNotFound, gin.H{ "error": "produto não foi encontrado" })
		return
	}	
}

func filterProdutosById(context *gin.Context) {
	fieldid := context.Param("id")

	if strings.TrimSpace(fieldid) == "" {
		context.JSON(http.StatusNotFound, gin.H{ "error": "campo id é obrigatório" })
		return 
	}
	
	field, err := strconv.Atoi(fieldid)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{ "error": "informe um valor numerico para busca de um produto valido" })
		return
	}

	var (
		produtoEncontrado Produtos
		encontrado bool
	)
	fmt.Println(field)

	for _, value := range produtos {
		if value.Id == field {
			produtoEncontrado = value
			encontrado = true
		}
	}

	if encontrado {
		context.JSON(http.StatusOK, gin.H{ "produto": produtoEncontrado })  
	} else {
		context.JSON(http.StatusNotFound, gin.H{ "message": "produto não existe no banco"})  
		return 
	}
	  
}

// ShopPlace3Santana
func main() {
	
	router := gin.Default()
	router.GET("/produtos", listProdutos)
	router.POST("/produtos", add)
	router.GET("/produtos-filter", filterProdutos)
	router.GET("/produtos/:id", filterProdutosById)

	router.Run() 
}

