package main

import (
	"log"

	"web-service-gin/cmd/server/controllers"
	"web-service-gin/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	 
)


func TokenAuthMiddleware() gin.HandlerFunc {
	return func (context *gin.Context)  {
		context.JSON(200, gin.H{ "message" : "oi" })
	}
}


// ShopPlace3Santana
func main() {
	 
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}	
	rep := products.NewRepository()

	service := products.NewService(rep)

	p := controllers.NewProduto(service)
	
	router := gin.Default()
	group := router.Group("/produtos")
	{
		group.GET("/", p.GetAll())
		group.POST("/", p.Store())
		group.PUT("/:id", p.Update())
		group.PATCH("/:id", p.UpdateName())
		group.DELETE("/:id", p.Delete())
	}
	  
	router.Run() 
}

