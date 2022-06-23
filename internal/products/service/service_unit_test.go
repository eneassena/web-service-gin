package service_products

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	model_products "web-service-gin/internal/products/model"
	products_repository "web-service-gin/internal/products/repository"
	"web-service-gin/pkg/store"

	"github.com/stretchr/testify/assert"
)

var pList []model_products.Produtos = []model_products.Produtos{
	{
		ID: 1,
		Name: "Iphone 7",
		Type: "Eletronico",
		Count: 1,
		Price: 80000,
	}, {
		ID: 2,
		Name: "iphone 8",
		Type: "Eletronico",
		Count: 1,
		Price: 95231,
	},
}


func TestUnitServiceGetAll(t *testing.T) {

	pListInByte, _ := json.Marshal(pList)

	fileStore := store.StoreMock{
		ReadMock: func(data interface{}) error {
			return json.Unmarshal(pListInByte, data)
		},
		WriteMock: func(data interface{}) error {
			return nil
		},
	}

	repository := products_repository.NewRepository(&fileStore)
	service := NewService(repository)
	listaProducts, _ := service.GetAll()

	assert.True(t, len(listaProducts) > 0)
}

func TestUnitServiceStore(t *testing.T) {
	 
	pListInByte, _ := json.Marshal(pList)

	newProduct := model_products.Produtos{
		ID: 1,
		Name: "Tenis",
		Type: "Calçados",
		Count: 1 ,
		Price: 10000,
	}

	fileStore := store.StoreMock{
		ReadMock: func(data interface{}) error {
			return json.Unmarshal(pListInByte, data)
		},
		WriteMock: func(data interface{}) error {
			_, err := json.Marshal(data)
			return err
		},
	}

	repository := products_repository.NewRepository(&fileStore)
	
	service := NewService(repository)

	storeProducts, err := service.Store(
		newProduct.Name, newProduct.Type, 
		newProduct.Count, newProduct.Price,
	)

	assert.Nil(t, err)
	assert.ObjectsAreEqual( newProduct, storeProducts)
}

func TestUnitGetOne(t *testing.T) {
	pListInByte, _ := json.Marshal(pList)

	newProduct := model_products.Produtos{
		ID: 1,
		Name: "Tenis",
		Type: "Calçados",
		Count: 1 ,
		Price: 10000,
	}

	fileStore := store.StoreMock{
		ReadMock: func(data interface{}) error {
			return json.Unmarshal(pListInByte, data)
		},
		WriteMock: func(data interface{}) error {
			return nil
		},
	}

	repository := products_repository.NewRepository(&fileStore)
	
	service := NewService(repository)
	
	searchProducts, err := service.GetOne(newProduct.ID)
	
	assert.Nil(t, err)
	assert.ObjectsAreEqual( newProduct, searchProducts)
}

func TestUnitServiceUpdate(t *testing.T) {
	t.Run("service test no metodo Update, caso de sucesso", func(t *testing.T) {
		pListInByte, _ := json.Marshal(pList)

		newProduct := model_products.Produtos{
			ID: 1,
			Name: "Tenis",
			Type: "Calçados",
			Count: 1 ,
			Price: 10000,
		}

		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(pListInByte, data)
			},
			WriteMock: func(data interface{}) error {
				_, err := json.Marshal(data)
				return err
			},
		}

		repository := products_repository.NewRepository(&fileStore)
		
		service := NewService(repository)

		updateProducts, err := service.Update(
		newProduct.ID, "iphone 88", 
			newProduct.Type, newProduct.Count, 
			newProduct.Price,
		)
		
		assert.Nil(t, err)
		assert.ObjectsAreEqual(newProduct, updateProducts)
		assert.Equal(t, "iphone 88", updateProducts.Name)
	})
	t.Run("service test no metodo Update, caso de error", func(t *testing.T) {
		pListInByte, _ := json.Marshal(pList)

		newProduct := model_products.Produtos{
			ID: 1,
			Name: "Tenis",
			Type: "Calçados",
			Count: 1 ,
			Price: 10000,
		}

		expectedError := errors.New("produto nao esta registrado")

		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(pListInByte, data)
			},
			WriteMock: func(data interface{}) error {
				_, err := json.Marshal(data)
				return err
			},
		}

		repository := products_repository.NewRepository(&fileStore)
		
		service := NewService(repository)

		_, err := service.Update(
		20, "iphone 88", 
			newProduct.Type, newProduct.Count, 
			newProduct.Price,
		)
		  
		assert.Equal(t, expectedError,err)
	})
}

func TestUnitServiceUpdateName(t *testing.T) {
	t.Run("service test no metodo UpdateName", func(t *testing.T) {
		pListInByte, _ := json.Marshal(pList)

		newProduct := model_products.Produtos{
			ID: 1,
			Name: "Tenis",
			Type: "Calçados",
			Count: 1 ,
			Price: 10000,
		}
		newNameProduct := "Tenis Nike"

		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(pListInByte, data)
			},
			WriteMock: func(data interface{}) error {
				_, err := json.Marshal(data)
				return err
			},
		}

		repository := products_repository.NewRepository(&fileStore)
		
		service := NewService(repository)

		storeProducts, err := service.UpdateName(newProduct.ID, newNameProduct)
		
		assert.Nil(t, err)
		assert.Equal(t, newNameProduct, storeProducts.Name)
	})

	t.Run("service test no metodo UpdateName, caso de error", func(t *testing.T) {
		pListInByte, _ := json.Marshal(pList)

		newNameProduct := "Tenis Nike"

		expectedError := errors.New("produto não esta registrado")

		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(pListInByte, data)
			},
			WriteMock: func(data interface{}) error {
				_, err := json.Marshal(data)
				return err
			},
		}

		repository := products_repository.NewRepository(&fileStore)
		
		service := NewService(repository)

		_, err := service.UpdateName(20, newNameProduct)
		
		assert.Equal(t, expectedError, err)
	})

}

func TestUnitServiceDelete(t *testing.T) {
	t.Run("service test no metodo Delete, caso de sucesso", func(t *testing.T) {
		pListInByte, _ := json.Marshal(pList)

		id := 1

		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(pListInByte, data)
			},
			WriteMock: func(data interface{}) error {
				_, err := json.Marshal(data)
				return err
			},
		}

		repository := products_repository.NewRepository(&fileStore)
		
		service := NewService(repository)

		err := service.Delete(id)
		
		assert.Nil(t, err)
	})
	t.Run("service test no metodo Delete, caso de error", func(t *testing.T) {
		
		id := 10

		expectedError := fmt.Errorf("error: ao remover o recurso %d", id)

		pListInByte, _ := json.Marshal(pList)
 
		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(pListInByte, data)
			},
			WriteMock: func(data interface{}) error {
				_, err := json.Marshal(data)
				return err
			},
		}

		repository := products_repository.NewRepository(&fileStore)
		
		service := NewService(repository)

		err := service.Delete(id)
		
		assert.NotNil(t, err)

		assert.Equal(t, expectedError, err)

	})
}