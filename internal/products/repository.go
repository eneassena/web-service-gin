package products

import (
	"fmt"  
)

 
var ps []Produtos = []Produtos{}

var lastID int

// arquivo implementa as regras de negocio da aplicação


type Repository interface {
	GetAll() ([]Produtos, error)
	Store(id int, name string, produtoType string, count int, price float64) (Produtos, error)
	LastID() (int, error)
	Update(id int, name string, produtoType string, count int, price float64) (Produtos, error)
	UpdateName(id int, name string) (Produtos, error)
	Delete(id int) error
}

type repository struct {}


func (repository) GetAll() ([]Produtos, error) {
	if len(ps) == 0 {
		return []Produtos{}, fmt.Errorf("não há produto registrados")
	}
	return ps, nil
}

func (repository) Store(id int, name string, produtoType string, count int , price float64) (Produtos, error) {

	lastID ++

	newProduto := Produtos{
		ID: lastID ,  Name: name, Type: produtoType, Count: count, Price: price,
	}
	
	ps = append(ps, newProduto)

	return newProduto, nil
}

func (repository) LastID() (int, error) {
	if lastID <= 0  {
		return 0, fmt.Errorf("lastId não encontrado")
	}
	return lastID, nil
}

func (repository) Update(id int, name string, produtoType string, count int , price float64) (Produtos, error) {
	
	updated := false
	updateProduto := Produtos{}
	
	for k, value := range ps {
		if value.ID == id { 
			ps[k].Name = name 
			ps[k].Type = produtoType
			ps[k].Count = count 
			ps[k].Price = price
			updateProduto = ps[k]
			updated = true 
			break
		}
	}
	 
	if !updated {
		return Produtos{}, fmt.Errorf("produto não encontrado")
	}

	return updateProduto, nil
}

func (repository) UpdateName(id int, name string) (Produtos, error) {
	updated := false
	updatedNameProduto := Produtos{}

 	for index, value := range ps {
		if value.ID == id {
			ps[index].Name = name
			updatedNameProduto = ps[index]
			updated = true
			break
		}
	}

	if !updated {
		return Produtos{}, fmt.Errorf("error: falha ao atualizar produto %s", name)
	}

	return updatedNameProduto, nil
}

func (rep repository) Delete(id int) error {
	if len(ps) == 0 {
		return fmt.Errorf("produto não esta registrado")
	}

	produtosUpdated, deleted := []Produtos{}, false

	for key, value := range ps {
		if value.ID == id {
			if len(ps)-1 == key {
				produtosUpdated = append(produtosUpdated, ps[:key]... )
				deleted = true
			} else {
				produtosUpdated = append(ps[:key], ps[key+1:]... )
				deleted = true
			}
			break
		}
	}
	if deleted {
		ps = produtosUpdated
		return nil
	}
	 
	return fmt.Errorf("produto não esta registrado")
}

func NewRepository() Repository {
	return &repository{}
}