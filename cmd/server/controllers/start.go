package controllers




import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	products_repository "web-service-gin/internal/products/repository/file"
	service_products "web-service-gin/internal/products/service"
	"web-service-gin/pkg/store"

	"github.com/gin-gonic/gin"
)


func CreateServver() *gin.Engine {
	gin.SetMode(gin.TestMode)
	
	_ = os.Setenv("TOKEN", "123456")

	db := store.New(store.FileType, "../../../internal/repositories/produtos.json")
	
	router := gin.Default()
	
	repo := products_repository.NewRepository(db)

	service := service_products.NewService(repo)

	NewProduto(router, service)

	return router
}

func CreateRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder){
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")
	return req, httptest.NewRecorder()
}