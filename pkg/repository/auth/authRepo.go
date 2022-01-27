package repo

import (
	"go-clean-architecture/pkg/dto"
	"go-clean-architecture/pkg/entity"
)

type AuthRepository interface {
	Login(userCredentials dto.UserCredential) (entity.Customer, error)
}
