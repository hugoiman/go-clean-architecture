package repo

import (
	"database/sql"
	"go-clean-architecture/pkg/entity"
)

func NewProductRepo(sql *sql.DB) ProductRepo {
	return &ProductRepoImpl{
		Db: sql,
	}
}

type ProductRepoImpl struct {
	Db *sql.DB
}

func (repo *ProductRepoImpl) GetAll() (products []entity.Product) {
	con := repo.Db
	query := "SELECT id, customer_id, name, price, qty FROM product"
	rows, _ := con.Query(query)

	var id, customer_id, name string
	var price, qty int
	var product entity.Product

	for rows.Next() {
		rows.Scan(
			&id, &customer_id, &name, &price, &qty,
		)
		product.SetId(id)
		product.SetCustomerId(customer_id)
		product.SetName(name)
		product.SetPrice(price)
		product.SetQty(qty)

		products = append(products, product)
	}

	defer con.Close()

	return products
}

func (repo *ProductRepoImpl) Get(name string) (product entity.Product, err error) {
	con := repo.Db
	query := "SELECT id, customer_id, name, price, qty FROM product WHERE name = ?"

	var customerId, id string
	var price, qty int
	err = con.QueryRow(query, name).Scan(
		&id, &customerId, &name, &price, &qty,
	)

	product.SetId(id)
	product.SetCustomerId(customerId)
	product.SetName(name)
	product.SetPrice(price)
	product.SetQty(qty)

	defer con.Close()
	return product, err
}
