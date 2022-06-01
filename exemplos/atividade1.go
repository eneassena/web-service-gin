package exemplos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Product struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
}

func showFile(c *gin.Context) {
	
	
	data, err := ioutil.ReadFile("./product.json")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return 
	}
	
	var prod []Product
	
	if err := json.Unmarshal(data, &prod); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
		
	}
	fmt.Println(prod)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": prod,
	})
}
