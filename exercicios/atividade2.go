package exercicios

import (
	"net/http"

	"github.com/gin-gonic/gin"
)




type Auth struct{
	Usuario string
	Password string 
}
type User struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	auth Auth `json:"authentication"`
}


func show(c *gin.Context) {
	us := User{
		FirstName: "Ana" ,
		LastName: "Lima" , 
	}
	us.auth.Usuario = "ana132"
	us.auth.Password = "123456"

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": us,
	})
}
func auth(context *gin.Context) {
	us := Auth{
			Usuario: "ana132",
			Password: "123456",
	}
	context.IndentedJSON(http.StatusOK, 
	 gin.H{
		 "data": us,
	 },
	)
}

func showAll(context *gin.Context) {
	us := []User{
		{
			FirstName: "Ana",
			LastName: "Silva",
		},
		{
			FirstName: "Carlos",
			LastName: "Pereira",
		},
		{
			FirstName: "Mateus",
			LastName: "Oliveira",
		},
		{
			FirstName: "Lucia",
			LastName: "Lima",
		},
		{
			FirstName: "Cremilda",
			LastName: "Santos",
		},
	}

	context.IndentedJSON(http.StatusOK, gin.H{
		"data": us,
	})

}






