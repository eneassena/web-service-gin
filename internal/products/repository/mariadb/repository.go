package repository

import (
	"context"
	"database/sql"
	"errors"

	domain "web-service-gin/internal/products/domain"
)

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.Repository {
	return &mariaDBRepository{db: db}
}

func (rep mariaDBRepository) GetAll(ctx context.Context) ([]domain.Produtos, error) {
	var produtosList []domain.Produtos = []domain.Produtos{}
	rows, err := rep.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return produtosList, err
	}

	defer rows.Close()
	for rows.Next() {
		var product domain.Produtos
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			return produtosList, err
		}
		produtosList = append(produtosList, product)
	}

	return produtosList, nil
}

func (rep mariaDBRepository) GetOne(id int) (domain.Produtos, error) {
	var produto domain.Produtos
	stmt := rep.db.QueryRow(sqlGetOne, id)

	err := stmt.Scan(
		&produto.ID,
		&produto.Name,
		&produto.Type,
		&produto.Count,
		&produto.Price,
	)
	if err != nil {
		return domain.Produtos{}, errors.New("produto não está registrado")
	}
	return produto, nil
}

func (rep *mariaDBRepository) Store(
	id int,
	name string,
	produtoType string,
	count int, price float64,
) (domain.Produtos, error) {
	stmt, err := rep.db.Prepare(sqlStore)
	if err != nil {
		return domain.Produtos{}, err
	}

	defer stmt.Close()

	product := domain.Produtos{
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
		return domain.Produtos{}, err
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
) (domain.Produtos, error) {
	stmt, err := rep.db.Prepare(sqlUpdate)
	if err != nil {
		return domain.Produtos{}, err
	}

	defer stmt.Close()

	product := domain.Produtos{
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

func (rep *mariaDBRepository) UpdateName(id int, name string) (domain.Produtos, error) {
	p, err := rep.GetOne(id)
	if err != nil {
		return domain.Produtos{}, err
	}

	p.Name = name

	stmt, err := rep.db.Prepare(sqlUpdateName)
	if err != nil {
		return domain.Produtos{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&p.Name, &p.ID)
	if err != nil {
		return domain.Produtos{}, err
	}

	return p, nil
}

func (rep *mariaDBRepository) Delete(id int) error {
	result, err := rep.db.Exec(sqlDelete, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("produto não foi removido")
	}
	return nil
}
