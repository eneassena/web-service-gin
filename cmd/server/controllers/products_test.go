package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	productController "web-service-gin/cmd/server/controllers"
	model_products "web-service-gin/internal/products/model"
	productRepository "web-service-gin/internal/products/repository/file"
	service_products "web-service-gin/internal/products/service"
	"web-service-gin/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ProductTest struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

func CreateServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	_ = os.Setenv("TOKEN", "123456")

	dbStore := store.New(store.FileType, "../../../internal/repositories/produtos.json")
	repo := productRepository.NewRepository(dbStore)
	service := service_products.NewService(repo)
	productController.NewProduto(router, service)

	return router
}

func CreateRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")
	return req, httptest.NewRecorder()
}

func Test_GetProductsAll(t *testing.T) {
	// criar um servidor e define suas rotas
	r := CreateServer()

	t.Run("GetAll", func(t *testing.T) {
		req, rr := CreateRequestTest(http.MethodGet, "/api/v1/products/", "")

		defer req.Body.Close()

		// diz ao servidor que ele pode atender a solicitação
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		objRes := struct {
			Code int
			Data []model_products.Produtos
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.True(t, len(objRes.Data) > 0)
	})
}

func Test_PostProducts(t *testing.T) {
	t.Run("create products, success", func(t *testing.T) {
		r := CreateServer()

		product := ProductTest{
			Name:  "Monito",
			Type:  "Informatica",
			Count: 1,
			Price: 100.00,
		}
		dt, _ := json.Marshal(product)

		req, rr := CreateRequestTest(http.MethodPost, "/api/v1/products/", string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("create products, error", func(t *testing.T) {
		r := CreateServer()

		product := ProductTest{
			Name:  "Monito",
			Count: 1,
			Price: 100.00,
		}
		dt, _ := json.Marshal(product)

		req, rr := CreateRequestTest(http.MethodPost, "/api/v1/products/", string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})
}

func Test_PutProducts(t *testing.T) {
	r := CreateServer()
	t.Run("update products, success", func(t *testing.T) {
		product := ProductTest{
			Name:  "Monito",
			Type:  "Informatica",
			Count: 10,
			Price: 100.00,
		}
		dt, _ := json.Marshal(product)

		req, rr := CreateRequestTest(http.MethodPut, "/api/v1/products/14", string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})	
	t.Run("update products, error, not found", func(t *testing.T) {
		product := ProductTest{
			Name:  "Monito",
			Type:  "Informatica",
			Count: 10,
			Price: 100.00,
		}
		dt, _ := json.Marshal(product)

		req, rr := CreateRequestTest(http.MethodPut, "/api/v1/products/1", string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})	
	t.Run("update products, error, no paramentro", func(t *testing.T) {
		product := ProductTest{
			Name:  "Monito",
			Type:  "Informatica",
			Count: 10,
			Price: 100.00,
		}
		dt, _ := json.Marshal(product)

		req, rr := CreateRequestTest(http.MethodPut, "/api/v1/products/abc", string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})	
	t.Run("update products, error, no body", func(t *testing.T) {
		product := ProductTest{
			Name:  "Monito",
			Count: 10,
			Price: 100.00,
		}
		dt, _ := json.Marshal(product)

		req, rr := CreateRequestTest(http.MethodPut, "/api/v1/products/14", string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})	
}

func Test_PutProductsName(t *testing.T) {
	r := CreateServer()

	t.Run("update Name products, success", func(t *testing.T) {
		productName := struct {
			Name string `json:"name"`
		}{}
		productName.Name = "Fone Game"
		dt, _ := json.Marshal(productName)

		req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/products/11", string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
	t.Run("update Name products, error", func(t *testing.T) {
		productName := struct {
			Name string `json:"name"`
		}{}
		productName.Name = ""
		dt, _ := json.Marshal(productName)

		req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/products/11", string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})
	t.Run("update Name products, error no paramentro", func(t *testing.T) {
		productName := struct {
			Name string `json:"name"`
		}{}
		productName.Name = "Fone"
		dt, _ := json.Marshal(productName)
		
		urlUpdate := fmt.Sprintf("/api/v1/products/%d", 100)
		req, rr := CreateRequestTest(http.MethodPatch, urlUpdate, string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	t.Run("update Name products, error no paramentro", func(t *testing.T) {
		productName := struct {
			Name string `json:"name"`
		}{}
		productName.Name = "Fone"
		dt, _ := json.Marshal(productName)
		
		urlUpdate := fmt.Sprintf("/api/v1/products/%s", "100s")
		req, rr := CreateRequestTest(http.MethodPatch, urlUpdate, string(dt))

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	 
}

func Test_DeleteProducts(t *testing.T) {
	
	r := CreateServer()

	t.Run("delete products, success", func(t *testing.T) {
		req, rr := newRequest(http.MethodDelete, fmt.Sprintf("/api/v1/products/%v", 12), "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 204, rr.Code)
	})
	t.Run("delete products, error", func(t *testing.T) {		
		req, rr := newRequest(http.MethodDelete, "/api/v1/products/9", "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)
	})
	t.Run("delete products, error", func(t *testing.T) {		
		req, rr := newRequest(http.MethodDelete, "/api/v1/products/10s", "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 400, rr.Code)
	})
}

func newRequest(m string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	return CreateRequestTest(m, url, body)
}