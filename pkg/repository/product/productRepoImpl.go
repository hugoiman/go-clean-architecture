package repo

import (
	"go-clean-architecture/config"
	"go-clean-architecture/pkg/entity"
)

func NewProductRepository(db *config.DB) ProductRepository {
	return &productRepositoryImpl{
		db:      db,
		product: entity.Product{},
	}
}

type productRepositoryImpl struct {
	db      *config.DB
	product entity.Product
}

func (repo *productRepositoryImpl) GetAll() (products []entity.Product) {
	conn := repo.db.ConnectSql()
	query := "SELECT id, customer_id, name, price, qty FROM product"
	rows, _ := conn.Query(query)

	var id, customerId, name string
	var price, qty int

	for rows.Next() {
		rows.Scan(
			&id, &customerId, &name, &price, &qty,
		)
		repo.product.SetId(id)
		repo.product.SetCustomerId(customerId)
		repo.product.SetName(name)
		repo.product.SetPrice(price)
		repo.product.SetQty(qty)

		products = append(products, repo.product)
	}

	defer rows.Close()
	defer conn.Close()

	return products
}

func (repo *productRepositoryImpl) Get(name string) (entity.Product, error) {
	conn := repo.db.ConnectSql()
	query := "SELECT id, customer_id, name, price, qty FROM product WHERE name = ?"

	var customerId, id string
	var price, qty int
	err := conn.QueryRow(query, name).Scan(
		&id, &customerId, &name, &price, &qty,
	)

	repo.product.SetId(id)
	repo.product.SetCustomerId(customerId)
	repo.product.SetName(name)
	repo.product.SetPrice(price)
	repo.product.SetQty(qty)

	defer conn.Close()
	return repo.product, err
}
