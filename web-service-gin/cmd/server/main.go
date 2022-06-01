package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
  
type Produtos struct {
	Id int 
	Nome string
	Preco float64
}

var produtos []Produtos = []Produtos{}

var lastID int

const tokenKey = "TOKEN"

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
	token := os.Getenv(tokenKey)
	if context.GetHeader("token") !=  token {
		context.JSON(401, gin.H{ "error": "token invalido" })
		return
	}

	context.JSON(200, gin.H{ 
		"data": produtos,
	})
}

func filterProdutos(context *gin.Context){ 
	token := os.Getenv(tokenKey)
	if context.GetHeader("token") !=  token {
		context.JSON(401, gin.H{ "error": "token invalido" })
		return
	}

	nome := context.Query("nome")
	
	if strings.TrimSpace(nome) == "" {
		context.JSON(http.StatusNotFound, gin.H{ "error": "paramento nome é obrigatório" })
		return
	}

	detalhes := ""

	for _, value := range produtos {	
		if value.Nome == nome {
			precoStr := fmt.Sprintf("%.2f", value.Preco)
			idStr := string(rune(value.Id))
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
	token := os.Getenv(tokenKey)
	if context.GetHeader("token") !=  token {
		context.JSON(401, gin.H{ "error": "token invalido" })
		return
	}

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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
		
	router := gin.Default()
	router.GET("/produtos", listProdutos)
	router.POST("/produtos", add)
	router.GET("/produtos-filter", filterProdutos)
	router.GET("/produtos/:id", filterProdutosById) 

	router.Run() 
}

