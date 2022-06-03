package products


type Produtos struct {
	ID int 			`json:"id" `
	Name string  	`json:"name"`
	Type string 	`json:"type"`
	Count int 		`json:"count"`
	Price float64 	`json:"price"`
}

