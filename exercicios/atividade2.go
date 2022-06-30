package exercicios

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)


func FilterProducts(context *gin.Context) {

	name := context.Query("name")
	if name == "" {
		context.JSON(http.StatusBadRequest, gin.H{ "error": "name do produto é obrigatório" })
		return 
	}

	produto, err := ioutil.ReadFile("./product.json")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "arquivo_error": err })
		return
	}

	var produtos []Produto

	err  = json.Unmarshal(produto, &produtos)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "json_error": err })
		return 
	}
	p := Produto{}
	ok := false

	for _, value := range produtos {
		if value.Name == name {
			p = value
			ok = true
		}
	}

	if !ok {
		context.JSON(http.StatusOK, gin.H{ "produto": "produto nao foi encontrado" })
		return 
	}
	context.JSON(http.StatusOK, gin.H{ "produto": p })
	
}

