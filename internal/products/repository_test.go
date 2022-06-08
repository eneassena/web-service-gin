package products

import (
	"testing"
	"web-service-gin/pkg/store"
)


func TestGetAll(t *testing.T){
	db := store.New("file", "product.json")
	rep := NewRepository(db)
	productList, err := rep.GetAll()
	if err != nil {
		t.Errorf("A função GetAll() retornou um error %w", err)
	}

	if len(productList) == 0 {
		t.Errorf("A função GetAll() não retorno nenhum produto")
	}
}