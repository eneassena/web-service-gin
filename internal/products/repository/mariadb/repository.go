package repository

import (
	"database/sql"
	"errors"

	model_products "web-service-gin/internal/products/model"
)

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) model_products.Repository {
	return &mariaDBRepository{db: db}
}

func (rep mariaDBRepository) GetAll() ([]model_products.Produtos, error) {
	var produtosList []model_products.Produtos = []model_products.Produtos{}
	rows, err := rep.db.Query(sqlGetAll)
	if err != nil {
		return produtosList, err
	}

	defer rows.Close()
	for rows.Next() {
		var product model_products.Produtos
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			return produtosList, err
		}
		produtosList = append(produtosList, product)
	}

	return produtosList, nil
}

func (rep mariaDBRepository) GetOne(id int) (model_products.Produtos, error) {
	var produto model_products.Produtos
	stmt := rep.db.QueryRow(sqlGetOne, id)
	defer rep.db.Close()	
	err := stmt.Scan(&produto.ID,
		&produto.Name,
		&produto.Type,
		&produto.Count,
		&produto.Price,
	)
	if err != nil {
		return model_products.Produtos{}, errors.New("produto não está registrado")
	}
	return produto, nil
}

func (rep *mariaDBRepository) Store(
	id int,
	name string,
	produtoType string,
	count int, price float64,
) (model_products.Produtos, error) {
	stmt, err := rep.db.Prepare(sqlStore)
	if err != nil {
		return model_products.Produtos{}, err
	}

	defer stmt.Close()

	product := model_products.Produtos{
		Name:  name,
		Type:  produtoType,
		Count: count,
		Price: price,
	}

	res, err := stmt.Exec(
		&product.Name, 
		&product.Type, 
		&product.Count, 
		&product.Price,
	)

	if err != nil {
		return model_products.Produtos{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return product, err
	}
	product.ID = int(lastID)

	return product, nil
}
 

func (rep *mariaDBRepository) Update(
	id int,
	name string,
	produtoType string,
	count int,
	price float64,
) (model_products.Produtos, error) {
	stmt, err := rep.db.Prepare(sqlUpdate)
	if err != nil {
		return model_products.Produtos{}, err
	}

	defer stmt.Close()

	product := model_products.Produtos{
		ID:    id,
		Name:  name,
		Type:  produtoType,
		Count: count,
		Price: price,
	}
	_, err = stmt.Exec(
		&product.Name,
		&product.Type,
		&product.Count,
		&product.Price,
		product.ID,
	)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (rep *mariaDBRepository) LastID() (int, error) {
	return 0, nil
}

func (rep *mariaDBRepository) UpdateName(id int, name string) (model_products.Produtos, error) {
	p, err := rep.GetOne(id)
	if err != nil {
		return model_products.Produtos{}, err
	}

	p.Name = name

	stmt, err := rep.db.Prepare(sqlUpdateName)
	if err != nil {
		return model_products.Produtos{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&p.Name, &p.ID)
	if err != nil {
		return model_products.Produtos{}, err
	}

	return model_products.Produtos{}, nil
}

func (rep *mariaDBRepository) Delete(id int) error {
	stmt, err := rep.db.Prepare(sqlDelete)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&id)
	if err != nil {
		return err
	}

	return nil
}
