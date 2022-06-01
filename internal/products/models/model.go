package models


type Product struct {
	ID int 			`json:"id"`
	Name string  	`json:"nome"`
	Type string 	`json:"type`
	Count int 		`json:"count"`
	Prico float64 	`json:"preco"`
}

