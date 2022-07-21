package mocks

import (
	"context"

	model_products "web-service-gin/internal/products/domain"

	"github.com/stretchr/testify/mock"
)

type ProductRepository struct {
	mock.Mock
}

func (p *ProductRepository) GetAll(ctx context.Context) ([]model_products.Produtos, error) {
	args := p.Called(ctx)

	var productList []model_products.Produtos

	if rf, ok := args.Get(0).(func(ctx context.Context) []model_products.Produtos); ok {
		productList = rf(ctx)
	} else {
		if args.Get(0) != nil {
			productList = args.Get(0).([]model_products.Produtos)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func(ctx context.Context) error); ok {
		err = rf(ctx)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}

	return productList, err
}

func (p *ProductRepository) GetOne(id int) (model_products.Produtos, error) {
	args := p.Called(id)

	var product model_products.Produtos

	if rf, ok := args.Get(0).(func(int) model_products.Produtos); ok {
		product = rf(id)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(model_products.Produtos)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func(int) error); ok {
		err = rf(id)
	} else {
		err = args.Error(1)
	}

	return product, err
}

func (p *ProductRepository) Store(produto model_products.Produtos) (model_products.Produtos, error) {
	args := p.Called(produto)

	var product model_products.Produtos

	if rf, ok := args.Get(0).(func(model_products.Produtos) model_products.Produtos); ok {
		product = rf(produto)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(model_products.Produtos)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func(model_products.Produtos) error); ok {
		err = rf(produto)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return product, err
}

func (p *ProductRepository) Update(produto model_products.Produtos) (model_products.Produtos, error) {
	args := p.Called(produto)

	var product model_products.Produtos

	if rf, ok := args.Get(0).(func(model_products.Produtos) model_products.Produtos); ok {
		product = rf(produto)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(model_products.Produtos)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func(model_products.Produtos) error); ok {
		err = rf(produto)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return product, err
}

func (p *ProductRepository) UpdateName(id int, name string) (string, error) {
	args := p.Called(id, name)

	var product string

	if rf, ok := args.Get(0).(func(int, string) string); ok {
		product = rf(id, name)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(string)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func(int, string) error); ok {
		err = rf(id, name)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return product, err
}

func (p *ProductRepository) Delete(id int) error {
	args := p.Called(id)

	var err error

	if rf, ok := args.Get(0).(func(int) error); ok {
		err = rf(id)
	} else {
		if args.Get(0) != nil {
			err = args.Error(0)
		}
	}
	return err
}
