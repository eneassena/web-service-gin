package controller

import (
	"web-service-gin/internal/products/domain"

	"github.com/gin-gonic/gin"
)

func NewRouter(route *gin.Engine, productService domain.Service) {
	pc := &ProdutoController{service: productService}

	group := route.Group("/api/v1/products")
	{
		group.GET("/", pc.GetAll())
		group.GET("/:id", pc.GetOne())
		group.POST("/", pc.Store())
		group.PUT("/:id", pc.Update())
		group.PATCH("/:id", pc.UpdateName())
		group.DELETE("/:id", pc.Delete())
	}
}
