package repo

import (
	"go-clean-architecture/config"
	"go-clean-architecture/pkg/entity"
)

func NewCustomerRepository(db *config.DB) CustomerRepository {
	return &CustomerRepoImpl{
		db:       db,
		customer: entity.Customer{},
	}
}

type CustomerRepoImpl struct {
	db       *config.DB
	customer entity.Customer
}

func (repo *CustomerRepoImpl) GetAll() (customers []entity.Customer) {
	conn := repo.db.ConnectSql()
	query := "SELECT id, username, name FROM customer"
	rows, _ := conn.Query(query)

	var id, username, name string
	for rows.Next() {
		rows.Scan(
			&id, &username, &name,
		)
		repo.customer.SetId(id)
		repo.customer.SetUsername(username)
		repo.customer.SetName(name)

		customers = append(customers, repo.customer)
	}

	defer rows.Close()
	defer conn.Close()

	return customers
}

func (repo *CustomerRepoImpl) Get(id string) (entity.Customer, error) {
	conn := repo.db.ConnectSql()
	query := "SELECT id, username, name FROM customer WHERE id = ? OR username = ?"

	var username, name string
	err := conn.QueryRow(query, id, id).Scan(
		&id, &username, &name,
	)
	repo.customer.SetId(id)
	repo.customer.SetUsername(username)
	repo.customer.SetName(name)

	defer conn.Close()
	return repo.customer, err
}
