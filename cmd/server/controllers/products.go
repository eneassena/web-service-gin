package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	service_products "web-service-gin/internal/products/service"
	web "web-service-gin/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProdutoController struct {
	service service_products.Service
}

 type produtoRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Count int `json:"count" binding:"required"`
	Price float64 `json:"price" binding:"required"` 
}

type produtoName struct {
	Name string `json:"name" binding:"required"`
}


type RequestError struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

type ResponseError struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}


func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "numeric":
		return "This field only accepts numbers"
	}
	return ""
}
 



func NewProduto(produtoService service_products.Service ) *ProdutoController {
	return &ProdutoController{
		service: produtoService,
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
		token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			context.JSON(http.StatusUnauthorized,
				web.DecodeError(
					http.StatusUnauthorized, 
					"você não tem permissão para fazer está solicitação" ),
				)
			return 
		}
		// cria uma estrutura que recebe a request body method post
		p := produtoRequest{}
		// faz o bind de um json recebido
		err := context.ShouldBindJSON(&p)
		var out []RequestError
		if err != nil {
			var jsonError *json.UnmarshalTypeError
			var validatorError validator.ValidationErrors
			switch {
			case errors.As(err, &jsonError):
				strin := strings.Split(jsonError.Error(), ":")[1]
				req := RequestError{ jsonError.Field, strin }
				context.JSON(400, gin.H{  
					"error": req,
				})
				return
			 
			case errors.As(err, &validatorError):
				out = make([]RequestError, len(validatorError))
				mapField := map[string]int{"Name": 0, "Type": 1, "Count": 2, "Price":3}
				typeAluno := reflect.TypeOf(p)

				for i, fe := range validatorError {
					indiceField := mapField[fe.Field()] // index of field
					field :=typeAluno.Field(indiceField)

					out[i] = RequestError{ field.Tag.Get("json") , msgForTag(fe.Tag())}
				}
				context.JSON(400, gin.H{ "error": out })
				return
				default: 
				context.JSON(400, ResponseError{Code: 400, Data: err.Error()})
				return 
			} 			 
 			// context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error() ))
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

 