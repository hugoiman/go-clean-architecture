package services

import (
	"go-clean-architecture/pkg/dto"
	"go-clean-architecture/pkg/entity"
)

type AuthService interface {
	Login(userClaims dto.UserCredential) (entity.Customer, error)
	GenerateToken(customer entity.Customer) string
}
