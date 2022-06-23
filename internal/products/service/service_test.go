package service_products

import (
	"errors"
	"fmt"
	"testing"
	"web-service-gin/internal/products/mocks"
	model_products "web-service-gin/internal/products/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productsList []model_products.Produtos = []model_products.Produtos{
	{
		ID: 1,
		Name: "Tenis",
		Type: "Calçados",
		Count: 1,
		Price: 1000.01,
	},
	{
		ID: 2,
		Name: "Sandalha",
		Type: "Calçados",
		Count: 1,
		Price: 1000.01,
	},
}


func TestServiceGetAll(t *testing.T)() {

	mockRep := new(mocks.ProductRepository)

	t.Run("test de integração service e repository, metodo GetAll, caso de sucesso", func(t *testing.T) {
		mockRep.On("GetAll"). 
			Return(productsList, nil).
			Once()

		service := NewService(mockRep)

		pList, err := service.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, productsList, pList)

		mockRep.AssertExpectations(t)
	})
	t.Run("test de integração service e repository, metodo GetAll, caso de error", func(t *testing.T) {
		var (
			products []model_products.Produtos = []model_products.Produtos{}
			expectedErr error = errors.New("não há produtos registrados")
		)

		mockRep.On("GetAll"). 
			Return(products, expectedErr). 
			Once()
		
		service := NewService(mockRep)

		pList, err := service.GetAll()

		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, products, pList)

		mockRep.AssertExpectations(t)
	})


}
func TestServiceGetOne(t *testing.T)() {

	mockRep := new(mocks.ProductRepository)

	searchProduct := model_products.Produtos{
		ID: 1,
		Name: "Tenis",
		Type: "Calçados",
		Count: 1,
		Price: 100.12,
	}

	t.Run("test de integração em repository e service no metodo GetOne, caso de sucesso", func(t *testing.T) {
		mockRep.On("GetOne", mock.AnythingOfType("int")).
			Return(searchProduct, nil).
			Once()
		
		service := NewService(mockRep)

		productEncontrado, err := service.GetOne(searchProduct.ID)

		assert.Nil(t, err)
		assert.ObjectsAreEqual(searchProduct, productEncontrado)

		mockRep.AssertExpectations(t)
	})
	t.Run("test de integração em repository e service no metodo GetOne, caso de error", func(t *testing.T) {

		expectedError := fmt.Errorf("produto não esta registrado")
		expectedProduct := model_products.Produtos{}

		mockRep.On("GetOne", mock.AnythingOfType("int")).
			Return(expectedProduct, expectedError).
			Once()
		
		service := NewService(mockRep)

		productEncontrado, err := service.GetOne(30)

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
		assert.ObjectsAreEqual(expectedProduct, productEncontrado)

		mockRep.AssertExpectations(t)
	})

}
func TestServiceStore(t *testing.T)() {

	mockRep := new(mocks.ProductRepository)

	t.Run("test de integração service e repository, caso de sucesso", func(t *testing.T) {
		newProduct := model_products.Produtos{
			ID: 1,
			Name: "Mause",
			Type: "Informatica",
			Count: 1,
			Price: 645.445,
		}
		mockRep.On("LastID").Return(1, nil).Once()
		mockRep.On("Store",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int"),
			mock.AnythingOfType("float64"),
			).
			Return(newProduct, nil).
			Once()
		
		service := NewService(mockRep)

		productCriado, err := service.Store(
			newProduct.Name,
			newProduct.Type,
			newProduct.Count,
			newProduct.Price,
		)

		assert.Nil(t, err)
		assert.ObjectsAreEqual(newProduct, productCriado)

		mockRep.AssertExpectations(t)
	})
	t.Run("test de integração service e repository caso de error", func(t *testing.T) {
		
		productVazio := model_products.Produtos{}

		
		expectedErr := fmt.Errorf("falha ao criar um novo produto") 
		mockRep.On("LastID").Return(1, nil).Once()
		mockRep.On("Store",  
				mock.AnythingOfType("int"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("int"),
				mock.AnythingOfType("float64"),		
			).
			Return(productVazio, expectedErr).
			Once()
		
		service := NewService(mockRep)

		_, err := service.Store(
			"Notbook",
			"informatica",
			1,
			56545.45,
		)

		assert.NotNil(t, err)
		assert.ObjectsAreEqual(expectedErr, err)

		mockRep.AssertExpectations(t)
	})
	
	t.Run("test de integração service e repository caso de error no LastID", func(t *testing.T) {
		
		productVazio := model_products.Produtos{}
	
		expectedErr := fmt.Errorf("falha ao criar um novo produto") 
		mockRep.On("LastID").Return(0, errors.New("id invalid")).Once()
		mockRep.On("Store",  
				mock.AnythingOfType("int"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("int"),
				mock.AnythingOfType("float64"),		
			).
			Return(productVazio, expectedErr).
			Once()
		
		service := NewService(mockRep)

		_, err := service.Store(
			"Notbook",
			"informatica",
			1,
			56545.45,
		)

		assert.NotNil(t, err)
		assert.ObjectsAreEqual(expectedErr, err)

		mockRep.AssertExpectations(t)
	})
	
}
func TestServiceUpdate(t *testing.T)() {

}
func TestServiceUpdateName(t *testing.T)() {

}
func TestServiceDelete(t *testing.T)() {

}

