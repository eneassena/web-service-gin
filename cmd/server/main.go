package main

import (
	"fmt"
	"log" 
	"os"

	"web-service-gin/cmd/server/controllers"
	"web-service-gin/cmd/server/docs"
	"web-service-gin/internal/products/repository"
	"web-service-gin/internal/products/service"
	"web-service-gin/pkg/store"

	"github.com/gin-gonic/gin" 
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func TokenAuthMiddleware() gin.HandlerFunc {
	return func (context *gin.Context)  {
		context.JSON(200, gin.H{ "message" : "oi" })
	} 
} 

// @title MALI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termoOfService http://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url http://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/License-2.0.html
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	store := store.New("file", "./internal/repositories/produtos.json")	
	rep := repository_products.NewRepository(store)
	service := service_products.NewService(rep)
	p := controllers.NewProduto(service)
	
	docs.SwaggerInfo.Host= fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	docs.SwaggerInfo.BasePath = "/api/v1"
	
	router := gin.Default()
	group := router.Group("/api/v1")
	{ 
		group.GET("/", p.GetAll())
		group.GET("/:id", p.GetOne())
		group.POST("/", p.Store())
		group.PUT("/:id", p.Update())
		group.PATCH("/:id", p.UpdateName())
		group.DELETE("/:id", p.Delete())
	}
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + os.Getenv("PORT")) 
 
}

 