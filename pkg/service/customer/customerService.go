package service

import (
	"context"
	"go-clean-architecture/pkg/dto"
	"go-clean-architecture/pkg/entity"
)

type CustomerService interface {
	GetAll() []entity.Customer
	Get(id string) (entity.Customer, error)
	GetMyInfo(context.Context) *dto.UserClaims
}
