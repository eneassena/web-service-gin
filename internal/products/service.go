package products

import "fmt"

// endpoint <-> controller <-> service <-> repository <-> db
// arquivo chama as regras de negocio

type Service interface {
	GetAll() ([]Produtos, error)
	Store(id int, name string, produtoType string, count int, price float64) (Produtos, error)
	Update(id int, name string, produtoType string, count int, price float64) (Produtos, error)
	UpdateName(id int, name string) (Produtos, error)
	Delete(id int) error
}

type service struct {
	repository Repository 
}

func NewService(repository Repository) Service {
	service := service{ 
		repository: repository,
	}
	return &service
}

func (s service) GetAll() ([]Produtos, error) { 
	produtos, err := s.repository.GetAll()
	if err != nil {
		return []Produtos{}, err
	}
	return produtos, nil
}

func (s service) Store(id int, name string, produtoType string, count int, price float64) (Produtos, error) {
	newProduto, err  := s.repository.Store(id, name, produtoType, count, price)
	if err != nil {
		return Produtos{}, fmt.Errorf("error: falha ao registra um novo produto, %w", err)
	}
	return newProduto, nil
}

func (s service) UpdateName(id int, name string) (Produtos, error) {
	produto, err := s.repository.UpdateName(id, name)
	if err != nil {
		return produto, err
	}
	return produto, nil
}

func (s service) Update(id int, name string, produtoType string, count int, price float64) (Produtos, error) {
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
