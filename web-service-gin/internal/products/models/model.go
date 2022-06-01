package models


type Product struct {
	ID int `json:"id"`
	Nome string  `json:"nome"`
	Preco float64 `json:"preco"`
}

