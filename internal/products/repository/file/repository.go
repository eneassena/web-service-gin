package products_repository

import (
	"errors"
	"fmt"
	"web-service-gin/internal/products/model"
	"web-service-gin/pkg/store"
)
 
type repository struct {
	db store.Store
}

func (rep repository) GetAll() ([]model_products.Produtos, error) {
	var produtosList []model_products.Produtos = []model_products.Produtos{}

	return produtosList, rep.db.Read(&produtosList) 
}

func (rep repository) GetOne(id int) (model_products.Produtos, error) {
	var (
		produtoList []model_products.Produtos
		produto model_products.Produtos
	)

	if err := rep.db.Read(&produtoList); err != nil {
		return produto, err
	}

	if len(produtoList) > 0 {
		for indice := range produtoList {
			if produtoList[indice].ID==id{
				return produtoList[indice], nil
			}
		}
	} 
	return produto, fmt.Errorf("produto não esta registrado")
}

func (rep *repository) Store(
	id int,
	name string, 
	produtoType string, 
	count int , price float64,
	) (model_products.Produtos, error) {	
	var (
		produtosList 	[]model_products.Produtos
		erro 			error 
	)
	
	if erro = rep.db.Read(&produtosList); erro != nil {
		return model_products.Produtos{}, erro
	} 

	product := model_products.Produtos{
		ID: id,
		Name: name, 
		Type: produtoType, 
		Count: count, 
		Price: price,
	}
	produtosList = append(produtosList, product)
	 
	if erro = rep.db.Write(produtosList); erro != nil {
		return model_products.Produtos{}, erro
	} 

	return product, nil
}

func (rep *repository) LastID() (int, error) {
	produtoList, err := rep.GetAll()
	if err != nil {
		return 0, err
	}
 
	totalProdutos := len(produtoList)

	if totalProdutos > 0 {
		return produtoList[totalProdutos-1].ID, nil
	}
	return 0, errors.New("não há sections registrados")
}

func (rep *repository) Update(id int, name string, produtoType string, count int , price float64) (model_products.Produtos, error) {
	var produtos []model_products.Produtos

	updateProduto := model_products.Produtos{Name: name, Type: produtoType, Count: count, Price: price}

	if err := rep.db.Read(&produtos); err != nil {
		return model_products.Produtos{}, err
	}
 
	for k := range produtos {
		if produtos[k].ID == id {
			updateProduto.ID = produtos[k].ID
			produtos[k] = updateProduto
			if err := rep.db.Write(produtos); err != nil {
				return model_products.Produtos{}, nil
			} 
			return updateProduto, nil
		}
	}
	return model_products.Produtos{}, fmt.Errorf("produto nao esta registrado") 
}

func (rep *repository) UpdateName(id int, name string) (model_products.Produtos, error) {
	 
	var produtoList []model_products.Produtos

	if err := rep.db.Read(&produtoList); err != nil {
		return model_products.Produtos{}, err
	}

 	for index := range produtoList {
		if produtoList[index].ID == id {
			produtoList[index].Name = name 
			if err := rep.db.Write(produtoList); err != nil {
				return model_products.Produtos{}, err
			}
			return produtoList[index], nil
		}
	} 
	return model_products.Produtos{}, fmt.Errorf("produto não esta registrado")
}

func (rep *repository) Delete(id int) error { 

	produtosUpdated := []model_products.Produtos{}

	if err := rep.db.Read(&produtosUpdated); err != nil {
		return err
	} 
	produtosUpdated, err := deleteItem(rep, produtosUpdated, id)


	if err != nil {
		return err
	}

	if err = rep.db.Write(produtosUpdated); err != nil {
		return err
	}

	return nil
}

func deleteItem(rep *repository, lista []model_products.Produtos, id int) ([]model_products.Produtos, error) {
	for index := range lista {
		if lista[index].ID == id {
			if len(lista)-1 == index {
				lista = append([]model_products.Produtos{}, lista[:index]... )
			} else {
				lista = append(lista[:index], lista[index+1:]... )
			} 
			return lista, nil
		}
	}
	return []model_products.Produtos{}, fmt.Errorf("error: ao remover o recurso %d", id)
}

func NewRepository(db store.Store) model_products.Repository {
	return &repository{db: db}
}