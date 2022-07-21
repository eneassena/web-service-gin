package mocks

import (
	model_products "web-service-gin/internal/products/domain"

	"github.com/stretchr/testify/mock"
)

type ProductService struct {
	mock.Mock
}

func (p *ProductService) GetAll() ([]model_products.Produtos, error) {
	args := p.Called()

	var productList []model_products.Produtos

	if rf, ok := args.Get(0).(func() []model_products.Produtos); ok {
		productList = rf()
	} else {
		if args.Get(0) != nil {
			productList = args.Get(0).([]model_products.Produtos)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}

	return productList, err
}

func (p *ProductService) GetOne(id int) (model_products.Produtos, error) {
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

func (p *ProductService) Store(produto model_products.ProdutoRequest) (model_products.Produtos, error) {
	args := p.Called(produto)

	var product model_products.Produtos

	if rf, ok := args.Get(0).(func(produto model_products.ProdutoRequest) model_products.Produtos); ok {
		product = rf(produto)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(model_products.Produtos)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func(produto model_products.ProdutoRequest) error); ok {
		err = rf(produto)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}

	return product, err
}

func (p *ProductService) Update(id int, produto model_products.ProdutoRequest) (model_products.Produtos, error) {
	args := p.Called(id, produto)

	var product model_products.Produtos

	if rf, ok := args.Get(0).(func(id int, produto model_products.ProdutoRequest) model_products.Produtos); ok {
		product = rf(id, produto)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(model_products.Produtos)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func(id int, produto model_products.ProdutoRequest) error); ok {
		err = rf(id, produto)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return product, err
}

func (p *ProductService) UpdateName(id int, name string) (string, error) {
	args := p.Called(id, name)

	var nameObj string

	if rf, ok := args.Get(0).(func(int, string) string); ok {
		nameObj = rf(id, name)
	} else {
		if args.Get(0) != nil {
			nameObj = args.Get(0).(string)
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
	return nameObj, err
}

func (p *ProductService) Delete(id int) error {
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
