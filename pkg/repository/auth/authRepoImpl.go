package repo

import (
	"go-clean-architecture/config"
	"go-clean-architecture/pkg/dto"
	"go-clean-architecture/pkg/entity"
)

func NewAuthRepository(db *config.DB) AuthRepository {
	return &authRepoImpl{
		db:       db,
		customer: entity.Customer{},
	}
}

type authRepoImpl struct {
	db       *config.DB
	customer entity.Customer
}

func (repo *authRepoImpl) Login(userCredentials dto.UserCredential) (entity.Customer, error) {
	conn := repo.db.ConnectSql()
	query := "SELECT id, username, name FROM customer WHERE username = ? AND password = ?"

	var id, username, name string
	err := conn.QueryRow(query, userCredentials.Username, userCredentials.Password).Scan(
		&id, &username, &name,
	)
	repo.customer.SetId(id)
	repo.customer.SetUsername(username)
	repo.customer.SetName(name)

	defer conn.Close()
	return repo.customer, err
}
