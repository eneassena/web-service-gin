package exemplos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var employee = map[string]string{
	"644": "Employee A",
	"755": "Employee B",
	"777": "Employee C",
}

func filterRequest(context *gin.Context) {
	body := context.Request.Body
	header := context.Request.Header
	metodo := context.Request.Method
	param_login := context.Query("login")
	param_pass := context.Query("senha")

	fmt.Println("Eu recebi algo!")
	fmt.Printf("\tMétodo: %s\n", metodo)
	fmt.Println("Conteúdo do cabeçario")

	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}

	fmt.Println("\to body é um io.ReadCloser:(%s), e para trabalhar com ele teremos que leia depois\n", body)
	fmt.Println("yay")

	fmt.Printf("Queries recebida: [%s,%s]\n", param_login, param_pass)
	context.String(http.StatusOK, "Eu recebi")

}

func paramsEndpoint(context *gin.Context)  {
	token := context.GetHeader("token")
	msg := context.Param("message")

	if token != "123456" {
		context.IndentedJSON(http.StatusForbidden, gin.H{
			"error": "não autorization",
		})
		return 
	}

	numberString := "20"
	fmt.Printf("\ntipo do numero antes do parce %T", numberString)
	if number, err := strconv.ParseInt(numberString, 10, 6); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("\ntipo do numero após o parce %T\n", number)
	}
	
	

	context.JSON(http.StatusOK , gin.H{
		"mensagem": msg,
		"head": token,
	})
}

func SearchEmployee(context * gin.Context) {
	var id string = context.Param("id")
	p := context.Params
	if len(p) > 0 {
		fmt.Println("Total de paramentros recebidos = ", len(p), p)
		for key, value := range p {
			fmt.Printf("\nkey %d, value %s -> %s\n", key, value.Key, value.Value)
		}
	}
	fmt.Printf("tipo do paramentro: %T\n", id)
	

	e, err := strconv.ParseInt(id, 32, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return 
	} else {
		emp := employee[id]
		if emp != "" {
			message := fmt.Sprintf("Informações do employee %s, Nome: %s", id, emp)
			
			context.JSON(http.StatusOK, gin.H{ 
				"employee": message,
			})
			
			fmt.Println("valor de retorno da converção para inteiro: ", e)

		}else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "employee não encontrado",
			})
			return 
		}
	}
}