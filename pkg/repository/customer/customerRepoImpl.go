package repo

import (
	"database/sql"
	"go-clean-architecture/pkg/entity"
)

func NewCustomerRepo(sql *sql.DB) CustomerRepo {
	return &CustomerRepoImpl{
		Customer: entity.Customer{},
		Db:       sql,
	}
}

type CustomerRepoImpl struct {
	Db       *sql.DB
	Customer entity.Customer
}

func (repo *CustomerRepoImpl) GetAll() (customers []entity.Customer) {
	con := repo.Db
	query := "SELECT id, username, name FROM customer"
	rows, _ := con.Query(query)

	var id, username, name string
	for rows.Next() {
		rows.Scan(
			&id, &username, &name,
		)
		repo.Customer.SetId(id)
		repo.Customer.SetUsername(username)
		repo.Customer.SetName(name)

		customers = append(customers, repo.Customer)
	}

	defer con.Close()

	return customers
}

func (repo *CustomerRepoImpl) Get(id string) (customer entity.Customer, err error) {
	con := repo.Db
	query := "SELECT id, username, name FROM customer WHERE id = ? OR username = ?"

	var username, name string
	err = con.QueryRow(query, id, id).Scan(
		&id, &username, &name,
	)

	defer con.Close()
	return customer, err
}
