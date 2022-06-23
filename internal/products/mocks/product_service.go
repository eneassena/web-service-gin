package mocks


import (
	model_products "web-service-gin/internal/products/model"

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

func (p *ProductService) Store(
		id int, 
		name string, 
		produtoType string, 
		count int, 
		price float64,
	) (model_products.Produtos, error){

		args := p.Called(id, name, produtoType, count, price)

		var product model_products.Produtos

		if rf, ok := args.Get(0).(func(int,string,string,int,float64) model_products.Produtos); ok {
			product = rf(id, name, produtoType, count, price)
		}  else {
			if args.Get(0) != nil {
				product = args.Get(0).(model_products.Produtos)
			}
		}

		var err error 

		if rf, ok := args.Get(1).(func(int, string, string, int, float64) error); ok {
			err = rf(id, name, produtoType, count, price)
		} else {
			if args.Get(1) != nil {
				err = args.Error(1)
			}
		}

		return product, err
} 

func (p *ProductService) Update(
		id int, 
		name string, 
		produtoType string, 
		count int, 
		price float64,
	) (model_products.Produtos, error) {

		args := p.Called(id, name, produtoType, count, price)

		var product model_products.Produtos

		if rf, ok := args.Get(0).(func(int,string,string,int,float64) model_products.Produtos); ok {
			product = rf(id, name, produtoType, count, price)
		} else {
			if args.Get(0) != nil {
				product = args.Get(0).(model_products.Produtos)
			}
		}

		var err error 

		if rf, ok := args.Get(1).(func(int,string,string,int,float64) error); ok {
			err = rf(id, name, produtoType, count, price)
		} else {
			if args.Get(1) != nil {
				err = args.Error(1)
			}
		}
		return product, err
}


func (p *ProductService) UpdateName(id int, name string) (model_products.Produtos, error) {
	args := p.Called(id, name)
	
	var product model_products.Produtos

	if rf, ok := args.Get(0).(func(int,string) model_products.Produtos); ok {
		product = rf(id, name)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(model_products.Produtos)
		}
	}

	var err error 

	if rf, ok := args.Get(1).(func(int,string) error); ok {
		err = rf(id, name)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return product, err
}

func (p *ProductService) Delete(id int ) error {
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