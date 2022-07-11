package service_products

import (
	"context"
	"fmt"

	"web-service-gin/internal/products/domain"
)

// endpoint <-> controller <-> service <-> repository <-> db
// arquivo chama as regras de negocio

type service struct {
	repository domain.Repository
}

func NewService(repository domain.Repository) domain.Service {
	service := service{
		repository: repository,
	}
	return &service
}

func (s service) GetAll(ctx context.Context) ([]domain.Produtos, error) {
	produtos, err := s.repository.GetAll(ctx)
	if err != nil {
		return []domain.Produtos{}, err
	}
	return produtos, nil
}

func (s *service) GetOne(id int) (domain.Produtos, error) {
	produto, erro := s.repository.GetOne(id)
	if erro != nil {
		return domain.Produtos{}, erro
	}
	return produto, nil
}

func (s *service) Store(
	name string,
	produtoType string,
	count int,
	price float64,
) (domain.Produtos, error) {
	newProduto, err := s.repository.Store(0, name, produtoType, count, price)
	if err != nil {
		return domain.Produtos{}, fmt.Errorf("error: falha ao registra um novo produto, %w", err)
	}
	return newProduto, nil
}

func (s *service) UpdateName(id int, name string) (domain.Produtos, error) {
	if ok, err := productExists(s, id); !ok {
		return domain.Produtos{}, err
	}
	produto, err := s.repository.UpdateName(id, name)
	if err != nil {
		return domain.Produtos{}, err
	}
	return produto, nil
}

func (s *service) Update(
	id int,
	name string,
	produtoType string,
	count int,
	price float64,
) (domain.Produtos, error) {
	if ok, err := productExists(s, id); !ok {
		return domain.Produtos{}, err
	}

	produto, err := s.repository.Update(id, name, produtoType, count, price)
	if err != nil {
		return produto, err
	}

	fmt.Println(produto, err)
	return produto, nil
}

func (s *service) Delete(id int) error {
	if ok, err := productExists(s, id); !ok {
		return err
	}

	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func productExists(s *service, id int) (bool, error) {
	_, err := s.GetOne(id)
	if err != nil {
		return false, err
	}
	return true, nil
}
