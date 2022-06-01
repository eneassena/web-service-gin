package main

import "github.com/gin-gonic/gin"




 
func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(200, "exercicio 1")
	})

	router.Run()
	
}

