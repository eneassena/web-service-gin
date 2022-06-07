package products

import (
	"fmt"
	"web-service-gin/pkg/store"
)

/* 
	manutenção o código prescisa buscar o ultimo id inserido para salvar um novo registro no arquivo.json
	caso feche o server e volte o id sempre começará do valor 1
*/ 

// arquivo implementa as regras de negocio da aplicação
//var ps []Produtos = []Produtos{}


type Repository interface {
	GetAll() ([]Produtos, error)
	GetOne(id int) (Produtos, error)
	Store(id int, name string, produtoType string, count int, price float64) (Produtos, error)
	LastID() (int, error)
	Update(id int, name string, produtoType string, count int, price float64) (Produtos, error)
	UpdateName(id int, name string) (Produtos, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}


func (rep repository) GetAll() ([]Produtos, error) {
	var produtosList []Produtos = []Produtos{}

	return produtosList, rep.db.Read(&produtosList) 
}

func (rep repository) GetOne(id int) (Produtos, error) {
	var (
		produtoList []Produtos
		produto Produtos
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


func (rep repository) Store(id int, name string, produtoType string, count int , price float64) (Produtos, error) {	
	var (
		produtosList 	[]Produtos
		erro 			error
		lastID 			int
	)
	
	if erro = rep.db.Read(&produtosList); erro != nil {
		return Produtos{}, erro
	}
	
	lastID, erro = rep.LastID()
	if erro != nil {
		return Produtos{}, erro
	}
	lastID ++

	p := Produtos{ID: lastID, Name: name, Type: produtoType, Count: count, Price: price}
	produtosList = append(produtosList, p)
	 
	if erro = rep.db.Write(produtosList); erro != nil {
		return Produtos{}, erro
	} 

	return p, nil
}

func (rep repository) LastID() (int, error) {
	var produtoList []Produtos

	if err := rep.db.Read(&produtoList); err != nil {
		return 0, err
	}
	totalProdutos := len(produtoList)

	if totalProdutos > 0 {
		return produtoList[totalProdutos-1].ID, nil
	}
	return 0, nil
}

func (rep repository) Update(id int, name string, produtoType string, count int , price float64) (Produtos, error) {
	var produtos []Produtos

	updateProduto := Produtos{Name: name, Type: produtoType, Count: count, Price: price}

	if err := rep.db.Read(&produtos); err != nil {
		return Produtos{}, err
	}
 
	for k := range produtos {
		if produtos[k].ID == id {
			updateProduto.ID = produtos[k].ID
			produtos[k] = updateProduto
			if err := rep.db.Write(produtos); err != nil {
				return Produtos{}, nil
			} 
			return updateProduto, nil
		}
	}
	return Produtos{}, fmt.Errorf("produto nao esta registrado") 
}

func (rep repository) UpdateName(id int, name string) (Produtos, error) {
	 
	var produtoList []Produtos

	if err := rep.db.Read(&produtoList); err != nil {
		return Produtos{}, err
	}

 	for index := range produtoList {
		if produtoList[index].ID == id {
			produtoList[index].Name = name 
			if err := rep.db.Write(produtoList); err != nil {
				return Produtos{}, err
			}
			return produtoList[index], nil
		}
	} 
	return Produtos{}, fmt.Errorf("produto não esta registrado")
}

func (rep repository) Delete(id int) error { 

	produtosUpdated := []Produtos{}

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

func deleteItem(rep repository, lista []Produtos, id int) ([]Produtos, error) {
	for index := range lista {
		if lista[index].ID == id {
			if len(lista)-1 == index {
				lista = append([]Produtos{}, lista[:index]... )
			} else {
				lista = append(lista[:index], lista[index+1:]... )
			} 
			return lista, nil
		}
	}
	return []Produtos{}, fmt.Errorf("error: ao remover o recurso %d", id)
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}