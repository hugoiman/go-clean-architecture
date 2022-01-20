package repo

import (
	"go-clean-architecture/config"
	"go-clean-architecture/pkg/dto"
)

func NewAuthRepository(db *config.DB) AuthRepository {
	return &authRepoImpl{
		db: db,
	}
}

type authRepoImpl struct {
	db *config.DB
}

func (repo *authRepoImpl) Login(userCredentials dto.UserCredential) error {
	return nil
}
