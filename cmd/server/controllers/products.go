package controllers

import (
	"net/http"
	"os"
	"strconv"
	"web-service-gin/internal/products"
	"web-service-gin/pkg/web"

	"github.com/gin-gonic/gin"
)

type ProdutoController struct {
	service products.Service
}

// reflete as regras de negocio da aplicação

/**
 *
 *  Responses of ProdutosController
 * 
 */


 type produtoRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Count int `json:"count" binding:"required"`
	Price float64 `json:"price" binding:"required"` 
}

type produtoName struct {
	Name string `json:"name" binding:"required"`
}


func NewProduto(produtoService products.Service ) *ProdutoController {
	return &ProdutoController{
		service: produtoService,
	}
}

func (controller *ProdutoController) GetAll() gin.HandlerFunc {
	return func (context *gin.Context)  {
		token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, 
				web.DecodeError(http.StatusUnauthorized, "você não tem permissão para fazer está solicitação"),
			)
			return 
		}

		produtos, err := controller.service.GetAll()
		if err != nil {
			context.JSON(http.StatusNotFound, 
				web.DecodeError(http.StatusNotFound, err.Error()),
			)
			return
		}

		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, produtos))
	}
}

func (controller *ProdutoController) Store() gin.HandlerFunc {
	return func (context *gin.Context)  {
		token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized,
				web.DecodeError(
					http.StatusUnauthorized, 
					"você não tem permissão para fazer está solicitação" ),
				)
			return 
		}

		p := produtoRequest{}
		if err := context.ShouldBindJSON(&p); err != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error() ))
			return 
		}

		if p.Name == "" {
			context.JSON(http.StatusBadRequest,  web.DecodeError(http.StatusBadRequest, "campo name é obrigatório" ))
			return 
		}

		if p.Type == "" {
			context.JSON(http.StatusBadRequest,  web.DecodeError(http.StatusBadRequest, "o campo type é obrigatório" ))
			return 
		}

		if p.Count == 0 {
			context.JSON(http.StatusBadRequest,  web.DecodeError(http.StatusBadRequest, "O campo count é obrigatório" ))
			return 
		}

		if p.Price == 0 {
			context.JSON(http.StatusBadRequest,  web.DecodeError(http.StatusBadRequest,"O campo price é obrigatório" ))
			return
		}

		newproduto, err := controller.service.Store(0, p.Name, p.Type, p.Count, p.Price)
		if err != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error() ))
		}
		
		context.JSON(http.StatusOK,  web.NewResponse(http.StatusOK, newproduto ))
	}
}

func (controller *ProdutoController) Update () gin.HandlerFunc {
	return func (context *gin.Context)  {
		token := context.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized,
				web.DecodeError(http.StatusUnauthorized, "você não tem permissão para fazer está solicitação" ))
			return 
		}

		id := context.Param("id")
		numberId, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error() ))
			return 
		}
		
		produto := produtoRequest{}
		if err := context.ShouldBindJSON(&produto); err != nil { 
			context.JSON(http.StatusBadRequest, 
				web.DecodeError(http.StatusBadRequest, err.Error() ))
			return 
		}
		
		produtoUpdated, errUpdated := controller.service.Update(numberId, produto.Name, produto.Type, produto.Count, produto.Price)
		if errUpdated != nil {
			context.JSON(http.StatusNotFound, 
				web.DecodeError(http.StatusBadRequest, errUpdated.Error()))
			return 
		}
		 
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, produtoUpdated ))
	}
}

func (controller *ProdutoController) UpdateName() gin.HandlerFunc {
return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, 
				web.DecodeError(http.StatusUnauthorized, 
					"você não tem permissão para fazer está solicitação" ))
			return 
		}

		id := context.Param("id")
		numberId, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error() ))
			return
		}

		var produto produtoName
		if err := context.ShouldBindJSON(&produto); err != nil {
			context.JSON(http.StatusBadRequest,web.DecodeError(http.StatusBadRequest, err.Error() ))
			return 
		}

		produtoUpdated, falho := controller.service.UpdateName(numberId, produto.Name)
		if falho != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, falho.Error() ))
			return 
		}

		context.JSON(200, web.NewResponse(http.StatusOK, produtoUpdated ))
	}
}

func (controller *ProdutoController) Delete() gin.HandlerFunc {
	return func (context *gin.Context)  {
		token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, 
				web.NewResponse(http.StatusUnauthorized , 
					"você não tem permissão para fazer está solicitação" ))
			return 
		}

		id := context.Param("id")
		numberId, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error() ))
			return 
		}

		if err = controller.service.Delete(numberId); err != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest , err.Error() ))
			return 
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK,  "OK" ))
	}
}

 