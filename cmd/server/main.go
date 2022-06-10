package main

import (
	"fmt"
	"log"
	"os"
	"web-service-gin/cmd/server/controllers"
	"web-service-gin/cmd/server/docs"
	"web-service-gin/internal/products/repository"
	products_service "web-service-gin/internal/products/service"
	"web-service-gin/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{ "code": code, "error": message })
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == ""{
		log.Fatal("Please set API_TOKEN enviroment variable")
	}

	return func (context *gin.Context)  {
		token := context.Request.Header.Get("token")

		if token == "" {
			responseWithError(context, 401, "API token required")
			return
		}

		if token != requiredToken {
			responseWithError(context, 401, "invalid API token")
			return
		}
		context.Next()
	}	
}

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")

	c.Next()
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
	// ler arquivo env com variáveis de ambiente
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	store := store.New("file", "./internal/repositories/produtos.json")	
	rep := repository_products.NewRepository(store)
	service := products_service.NewService(rep)
	p := controllers.NewProduto(service)
	
	docs.SwaggerInfo.Host= fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	docs.SwaggerInfo.BasePath = "/api/v1"
	
	router := gin.Default()
	//router.Use(DummyMiddleware)
	router.Use(TokenAuthMiddleware())
	group := router.Group("/api/v1/products")
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

 