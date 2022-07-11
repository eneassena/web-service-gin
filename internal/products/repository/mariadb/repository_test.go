package repository

import (
	"context"
	"database/sql"
	"log"
	"testing"

	domain "web-service-gin/internal/products/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

var product = domain.Produtos{
	ID:    1,
	Name:  "iphone",
	Type:  "Eletronicos",
	Count: 1,
	Price: 10000,
}

func TestGetAll(t *testing.T) {
	t.Run("teste insert products, (sucesso)", func(t *testing.T) {
		db, mock := NewMock() // cria um mock do bando de dados

		defer db.Close() // fecha conexão

		// cria um array de products mockado
		mockProducts := []domain.Produtos{
			{
				ID:    1,
				Name:  "Playstation 5",
				Type:  "Eletrônicos",
				Count: 1,
				Price: 4500,
			},
			{
				ID:    2,
				Name:  "XBOX Series X",
				Type:  "Eletrônicos",
				Count: 1,
				Price: 4500,
			},
		}
		// cria os retorno que esperamos do banco de dados
		rows := sqlmock.NewRows([]string{
			"id", "name", "type", "count", "price",
		}).AddRow(
			mockProducts[0].ID,
			mockProducts[0].Name,
			mockProducts[0].Type,
			mockProducts[0].Count,
			mockProducts[0].Price,
		).AddRow(
			mockProducts[1].ID,
			mockProducts[1].Name,
			mockProducts[1].Type,
			mockProducts[1].Count,
			mockProducts[1].Price,
		)

		// cria a query que será executada
		query := "SELECT \\* FROM products"
		// executa a query criada
		mock.ExpectQuery(query).WillReturnRows(rows)
		// cria um repository com db mockado
		productsRepo := NewMariaDBRepository(db)
		// obtem o resultado do metodo testado
		ctx := context.Background()
		result, err := productsRepo.GetAll(ctx)
		// verifica o error de retorno do metodo
		assert.NoError(t, err)
		// verifica a lista de produtos retornada pelo metodo testado
		assert.Equal(t, result[0].Name, "Playstation 5")
		assert.Equal(t, result[1].Name, "XBOX Series X")
	})
}

func TestGetAllFailSelect(t *testing.T) {
	t.Run("test listar products (error)", func(t *testing.T) {
		db, mock := NewMock()

		defer db.Close()

		query := "SELECT \\* FROM products"

		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

		productsRepo := NewMariaDBRepository(db)
		ctx := context.Background()
		_, err := productsRepo.GetAll(ctx)

		assert.Error(t, err)
	})
}

func TestStore(t *testing.T) {
}
