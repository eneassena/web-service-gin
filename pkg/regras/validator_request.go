package regras

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"
	//model_products "web-service-gin/internal/products/model"
	"web-service-gin/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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
 

func ValidateErrorInRequest(context *gin.Context, data any) bool {
	var out []RequestError
	if err := context.ShouldBind(&data); err != nil {
		var validatorError validator.ValidationErrors
		var jsonError *json.UnmarshalTypeError
		// var jsonFieldError *json.UnmarshalFieldError
		switch {
		case errors.As(err, &jsonError):
			strin := strings.Split(jsonError.Error(), ":")[1]
			req := RequestError{ jsonError.Field, strin }
			context.JSON(http.StatusBadRequest, 
				web.NewResponse(http.StatusBadRequest, req ))
			return true
			case errors.As(err, &validatorError):
				out = make([]RequestError, len(validatorError))
				mapField := map[string]int{"Name": 0, "Type": 1, "Count": 2, "Price":3}
				typeAluno := reflect.TypeOf(data)
				for i, fe := range validatorError {
					indiceField := mapField[fe.Field()]
					field :=typeAluno.Field(indiceField)
					out[i] = RequestError{ field.Tag.Get("json") , msgForTag(fe.Tag())}
				} 
				context.JSON(http.StatusUnprocessableEntity, 
					web.NewResponse(http.StatusUnprocessableEntity, out ))
					return true
				default: 
				context.JSON(http.StatusUnprocessableEntity, 
					web.DecodeError(http.StatusUnprocessableEntity, err.Error() )) 
				return true
		}
	}
	return false
}

