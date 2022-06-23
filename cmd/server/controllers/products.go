package controllers

import (
 
	"net/http"
	"strconv"
	model_products "web-service-gin/internal/products/model"
	service_products "web-service-gin/internal/products/service"
	"web-service-gin/pkg/regras"
	"web-service-gin/pkg/web"

	"github.com/gin-gonic/gin"
)

type ProdutoController struct {
	service service_products.Service
}

type produtoName struct {
	Name string `json:"name" binding:"required"`
}
 
func NewProduto(r *gin.Engine, produtoService service_products.Service ) {
	pc := &ProdutoController{service: produtoService}

	group := r.Group("/products")
	{ 
		group.GET("/", pc.GetAll())
		group.GET("/:id", pc.GetOne())
		group.POST("/", pc.Store())
		group.PUT("/:id", pc.Update())
		group.PATCH("/:id", pc.UpdateName())
		group.DELETE("/:id", pc.Delete())
	
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 string data
// @Failure 401 error Error
// @Router /products [get]
func (controller *ProdutoController) GetAll() gin.HandlerFunc {
	return func (context *gin.Context)  { 
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
// @ ListOneProduct godoc
// @ Summary Search products
// @Tags Products
// @Description search one product
// @Accept json
// @Produce json
// @Param token path string true "token"
// @Param id path int true "Account ID"
// @Success 200 string data
// @Router /products/{id} [get]
func (controller *ProdutoController) GetOne() gin.HandlerFunc {
	return func (context *gin.Context) {
		id := context.Param("id")
		paramId, erro := strconv.Atoi(id)
		if erro != nil {
			context.JSON(
				http.StatusBadRequest,
				web.DecodeError(http.StatusBadRequest, erro.Error()),
			)
			return
		}	
		
		produto, erro := controller.service.GetOne(paramId)
		if erro != nil {
			context.JSON(
				http.StatusNotFound,
				web.DecodeError(http.StatusNotFound, erro.Error()),
			)
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, produto))
	}
}



func (controller *ProdutoController) Store() gin.HandlerFunc {
	return func (context *gin.Context)  {
		/*token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized,
				web.DecodeError(
					http.StatusUnauthorized, 
					"você não tem permissão para fazer está solicitação" ),
				)
			return 
		}*/
	// cria uma estrutura que recebe a request body method post
	var data model_products.ProdutoRequest	
	ok := regras.ValidateErrorInRequest(context, &data)
	if ok {	
		return
	}
	   
	newproduto, err := controller.service.Store(
		data.Name, 
		data.Type, 
		data.Count, 
		data.Price)

	if err != nil {
		context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error() ))
	}	
	context.JSON(http.StatusOK,  web.NewResponse(http.StatusOK, newproduto ))
	}
}

func (controller *ProdutoController) Update () gin.HandlerFunc {
	return func (context *gin.Context)  {
		/*token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized,
				web.DecodeError(http.StatusUnauthorized, "você não tem permissão para fazer está solicitação" ))
			return 
		}*/

		id := context.Param("id")
		numberId, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error() ))
			return 
		}
		
		var produto model_products.ProdutoRequest
		if regras.ValidateErrorInRequest(context, &produto) { 
			return 
		}
		
		produtoUpdated, errUpdated := controller.service.Update(numberId, produto.Name, produto.Type, produto.Count, produto.Price)
		if errUpdated != nil {
			context.JSON(http.StatusNotFound, 
				web.DecodeError(http.StatusNotFound, errUpdated.Error()))
			return 
		} 
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, produtoUpdated ))
	}
}

func (controller *ProdutoController) UpdateName() gin.HandlerFunc {
	return func(context *gin.Context) {
		/*token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, 
				web.DecodeError(http.StatusUnauthorized, 
					"você não tem permissão para fazer está solicitação" ))
			return 
		}*/

		id := context.Param("id")
		numberId, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error() ))
			return
		}

		var produto produtoName
		if regras.ValidateErrorInRequest(context, &produto) {
			return 
		}

		produtoUpdated, falho := controller.service.UpdateName(numberId, produto.Name)
		if falho != nil {
			context.JSON(http.StatusNotFound, web.DecodeError(http.StatusNotFound, falho.Error() ))
			return 
		}

		context.JSON(200, web.NewResponse(http.StatusOK, produtoUpdated ))
	}
}

func (controller *ProdutoController) Delete() gin.HandlerFunc {
	return func (context *gin.Context)  {
		/*token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized, 
				web.NewResponse(http.StatusUnauthorized , 
					"você não tem permissão para fazer está solicitação" ))
			return 
		}*/

		id := context.Param("id")
		numberId, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, 
				web.DecodeError(http.StatusBadRequest, err.Error() ))
			return 
		}

		if err = controller.service.Delete(numberId); err != nil {
			context.JSON(http.StatusNotFound, 
				web.DecodeError(http.StatusNotFound , err.Error() ))
			return 
		}
		context.JSON(http.StatusNoContent, 
			web.NewResponse(http.StatusNoContent,  "OK" ))
	}
}
 
 