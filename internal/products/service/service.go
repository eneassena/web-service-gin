package service_products

import (
	"fmt"
	"web-service-gin/internal/products/model"
)

// endpoint <-> controller <-> service <-> repository <-> db
// arquivo chama as regras de negocio

type Service interface {
	GetAll() ([]model_products.Produtos, error)
	Store(id int, name string, produtoType string, count int, price float64) (model_products.Produtos, error)
	Update(id int, name string, produtoType string, count int, price float64) (model_products.Produtos, error)
	UpdateName(id int, name string) (model_products.Produtos, error)
	Delete(id int) error
	GetOne(id int ) (model_products.Produtos, error)
}

type service struct {
	repository model_products.Repository 
}

func NewService(repository model_products.Repository) Service {
	service := service{ 
		repository: repository,
	}
	return &service
}

func (s service) GetAll() ([]model_products.Produtos, error) { 
	produtos, err := s.repository.GetAll()
	if err != nil {
		return []model_products.Produtos{}, err
	}
	return produtos, nil
}

func (s service) GetOne(id int) (model_products.Produtos, error) {
	produto, erro := s.repository.GetOne(id)
	if erro != nil {
		return model_products.Produtos{}, erro
	}
	return produto, nil
}

func (s service) Store(id int, 
		name string, 
		produtoType string, 
		count int, 
		price float64,
	) (model_products.Produtos, error) {
	newProduto, err  := s.repository.Store(id, name, produtoType, count, price)
	if err != nil {
		return model_products.Produtos{}, fmt.Errorf("error: falha ao registra um novo produto, %w", err)
	}
	return newProduto, nil
}

func (s service) UpdateName(id int, name string) (model_products.Produtos, error) {
	produto, err := s.repository.UpdateName(id, name)
	if err != nil {
		return produto, err
	}
	return produto, nil
}

func (s service) Update(
		id int, 
		name string, 
		produtoType string, 
		count int, 
		price float64,
	) (model_products.Produtos, error) {
	produto,err := s.repository.Update(id, name, produtoType, count, price)
	if err != nil {
		return produto, err
	}

	fmt.Println(produto, err)
	return produto, nil
}

func (s service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}
