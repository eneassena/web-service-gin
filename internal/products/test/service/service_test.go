package service_test

import (
	"errors"
	"testing"
	"web-service-gin/internal/products/mocks"
	"web-service-gin/internal/products/model"
	products_service "web-service-gin/internal/products/service"

	"github.com/stretchr/testify/assert"
)


func TestGetAll(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)

	product := model_products.Produtos{
		ID: 1,
		Name: "iphone",
		Type: "Eletronico",
		Price: 5000,
	}

	var productList = make([]model_products.Produtos, 0)
	productList = append(productList, product)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetAll").Return(productList, nil).Once()

		s := products_service.NewService(mockRepo)
		lista, err := s.GetAll()
		assert.NoError(t, err)

		assert.Equal(t, "iphone", lista[0].Name)

		mockRepo.AssertExpectations(t)
	})


	t.Run("error", func (t *testing.T)  {
		expected := errors.New("failed to retrive products")

		mockRepo.On("GetAll").Return(nil, expected).
		Once()

		s := products_service.NewService(mockRepo)

		_, err := s.GetAll()

		assert.NotNil(t, err)
		
		assert.Equal(t, err, expected)

		mockRepo.AssertExpectations(t)
	})
}