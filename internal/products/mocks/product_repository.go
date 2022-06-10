package mocks

import (
	model_products "web-service-gin/internal/products/model"

	"github.com/stretchr/testify/mock"
)


type ProductRepository struct {
	mock.Mock
}

// GetAll() ([]Produtos, error)
// GetOne(id int) (Produtos, error)
// Store(id int, name string, produtoType string, count int, price float64) (Produtos, error)
// LastID() (int, error)
// Update(id int, name string, produtoType string, count int, price float64) (Produtos, error)
// UpdateName(id int, name string) (Produtos, error)
// Delete(id int) error

func (p *ProductRepository) GetAll() ([]model_products.Produtos, error) {
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

func (p *ProductRepository) GetOne(id int) (model_products.Produtos, error) {
	args := p.Called()

	var product model_products.Produtos

	if rf, ok := args.Get(0).(func(id int) model_products.Produtos); ok {
		product = rf(id)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(model_products.Produtos)
		}
	}

	var err error 

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return product, err
}

func (p *ProductRepository) Store(
		id int, 
		name string, 
		produtoType string, 
		count int, 
		price float64,
	) (model_products.Produtos, error){

		args := p.Called()

		var product model_products.Produtos

		if rf, ok := args.Get(0).(func(
				id int, 
				name string, 
				produtoType string, 
				count int, 
				price float64,
		) model_products.Produtos); ok {
			product = rf(id, name, produtoType, count, price)
		}  else {
			if args.Get(0) != nil {
				product = args.Get(0).(model_products.Produtos)
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

		return product, err
}

func (p *ProductRepository) LastID()(int, error) {

	args := p.Called()

	var lastID int

	if rf, ok := args.Get(0).(func() int); ok {
		lastID = rf()
	} else {
		if args.Get(0) != nil {
			lastID = args.Get(0).(int)
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

	return lastID, err
}

func (p *ProductRepository) Update(
		id int, 
		name string, 
		produtoType string, 
		count int, 
		price float64,
	) (model_products.Produtos, error) {

		args := p.Called()

		var product model_products.Produtos

		if rf, ok := args.Get(0).(func(id int, 
			name string, produtoType string, count int, price float64) model_products.Produtos); ok {
				product = rf(id, name, produtoType, count, price)
		} else {
			if args.Get(0) != nil {
				product = args.Get(0).(model_products.Produtos)
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
		return product, err
}


func (p *ProductRepository) UpdateName(id int, name string) (model_products.Produtos, error) {
	args := p.Called()
	
	var product model_products.Produtos

	if rf, ok := args.Get(0).(func(id int, name string) model_products.Produtos); ok {
		product = rf(id, name)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(model_products.Produtos)
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

	return product, err

}

func (p *ProductRepository) Delete(id int ) error {
	args := p.Called()

	var err error 

	if rf, ok := args.Get(0).(func(id int) error); ok {
		err = rf(id)
	} else {
		if args.Get(0) != nil {
			err = args.Error(0)
		}
	}

	return err
}