package exercicios


// thunder client é uma exetnsão para teste de api
// no json o rotulo omitempty ocutará o campo caso esta vazio
// ShopPlace3Santana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)


 



func Exercicio2RetoandoJSON(context *gin.Context) {
	nome := "Enéas"
	message := fmt.Sprintf("Olá, %s", nome)

	context.JSON(http.StatusOK, 
		gin.H{
			"message": message,
		},
	)
}


func Exercicio3ListagemProdutos(c *gin.Context) {	
	data, err := ioutil.ReadFile("./product.json")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err })
		return 
	}

	var produtos []Produto
	
	if err := json.Unmarshal(data, &produtos); err != nil {
		c.JSON(http.StatusOK, gin.H{ "error": err })
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{ "produtcts": produtos })
}
