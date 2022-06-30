package model_products

 
type Produtos struct {
	ID int 			`json:"id" `
	Name string  	`json:"name"`
	Type string 	`json:"type"`
	Count int 		`json:"count"`
	Price float64 	`json:"price"`
}

type ProdutoRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Count int `json:"count" binding:"required"`
	Price float64 `json:"price" binding:"required"` 
}
type Repository interface {
	GetAll() ([]Produtos, error)
	GetOne(id int) (Produtos, error)
	Store(id int, name string, produtoType string, count int, price float64) (Produtos, error)
	LastID() (int, error)
	Update(id int, name string, produtoType string, count int, price float64) (Produtos, error)
	UpdateName(id int, name string) (Produtos, error)
	Delete(id int) error
} 
