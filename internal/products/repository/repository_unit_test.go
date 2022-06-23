package products_repository

import (
	"encoding/json"
	"testing"
	model_products "web-service-gin/internal/products/model"
	"web-service-gin/pkg/store"

	"github.com/stretchr/testify/assert"
)

var pList []model_products.Produtos = []model_products.Produtos{
	{
		ID: 1,
		Name: "Mause",
		Type: "informatica",
		Count: 5,
		Price: 50.0,
	},
	{
		ID: 2,
		Name: "Mause 2",
		Type: "informatica",
		Count: 5,
		Price: 50.0,
	},
}


func TestGetAll(t *testing.T) {
	t.Run("test no repository método GetAll", func(t *testing.T) {
		productListInByte, _ := json.Marshal(pList)
	
		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(productListInByte, data)
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		}

		repository := NewRepository(&fileStore)

		result, err := repository.GetAll()
		
		assert.Nil(t, err)
		assert.True(t, len(result) > 0)
		assert.Equal(t, pList,result)
	})
}

func TestRepositoryGetOne(t *testing.T) {
	t.Run("repository test no metodo getOne, caso de sucesso", func(t *testing.T) {
		productInByte, _ := json.Marshal(pList)

		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(productInByte, data)
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		}

		respository := NewRepository(&fileStore)

		productOne, err := respository.GetOne(1)
		assert.Nil(t, err)
		assert.ObjectsAreEqual(pList[0], productOne)
	})
}

func TestRepositoryStore(t *testing.T) {

	newProduct := model_products.Produtos{
		ID: 1,
		Name: "Notbook Dell",
		Type: "Informatica",
		Count: 1,
		Price: 51231,
	}

	t.Run("repository test no metodo Store, caso de sucesso", func(t *testing.T) {
		dataByte, _ := json.Marshal(pList)
		
		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(dataByte, data)
			},
			WriteMock: func(data interface{}) error {
				_, err := json.Marshal(data)				 
				return err
			},
		}

		repository := NewRepository(&fileStore)
		nProduct, err := repository.Store(
			newProduct.ID,
			newProduct.Name, 
			newProduct.Type, 
			newProduct.Count, 
			newProduct.Price,
		)
		assert.Nil(t, err)
		assert.ObjectsAreEqual(newProduct, nProduct)
	})
}


func TestRepositoryUpdate(t *testing.T) {
	t.Run("repositry test no metodo Update, caso de sucesso", func(t *testing.T) {

		productUpdate1 := model_products.Produtos{
			Name: "Notbook",
			Type: "Informatica",
			Count: 1,
			Price: 45312,
		}
		productInByte, _ := json.Marshal(pList)

		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(productInByte, data)
			},
			WriteMock: func(data interface{}) error {
				err := json.Unmarshal(productInByte, data)
				return err
			},
		}
		repository := NewRepository(&fileStore)

		productAlter, err := repository.Update(1, productUpdate1.Name,
			productUpdate1.Type, productUpdate1.Count, 
			productUpdate1.Price)

		assert.ObjectsAreEqual(productUpdate1, productAlter)
		assert.Nil(t, err)
	})
}

func TestRepositoryUpdateName(t *testing.T) {
	t.Run("repositry test no metodo Update, caso de sucesso", func(t *testing.T) {
 
		productInByte, _ := json.Marshal(pList)

		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(productInByte, data)
			},
			WriteMock: func(data interface{}) error {
				_, err := json.Marshal(data)
				return err
			},
		}
		repository := NewRepository(&fileStore)

		productAlter, err := repository.UpdateName(1, "Iphone 7")

		assert.Equal(t, "Iphone 7", productAlter.Name)
		assert.Nil(t, err)
	})
}

func TestRepositoryDelete(t *testing.T) {
	var id = 1

	dataByte, _ := json.Marshal(pList)

	t.Run("repository test no metodo delete, caso de sucesso", func(t *testing.T) {
		fileStore := store.StoreMock{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(dataByte, data)
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		}

		repository := NewRepository(&fileStore)

		err := repository.Delete(id)
		
		assert.Nil(t, err)
	})
}