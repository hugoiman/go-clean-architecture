package services

import (
	"time"

	"go-clean-architecture/pkg/dto"
	"go-clean-architecture/pkg/entity"
	authRepo "go-clean-architecture/pkg/repository/auth"
	customerServ "go-clean-architecture/pkg/service/customer"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func NewAuthService(authRepository *authRepo.AuthRepository, customerService *customerServ.CustomerService) AuthService {
	return &authServiceImpl{
		authRepository:  *authRepository,
		customerService: *customerService,
	}
}

type authServiceImpl struct {
	authRepository  authRepo.AuthRepository
	customerService customerServ.CustomerService
}

func (service *authServiceImpl) Login(userCredentials dto.UserCredential) (entity.Customer, error) {
	err := service.authRepository.Login(userCredentials)
	if err == nil {
		customer, _ := service.customerService.Get(userCredentials.Username)
		return customer, err
	}
	return entity.Customer{}, err
}

func (service *authServiceImpl) GenerateToken(customer entity.Customer) string {
	userClaims := dto.UserClaims{
		ID:       customer.GetId(),
		Username: customer.GetUsername(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, _ := token.SignedString([]byte(viper.GetString("secret-key")))

	return tokenString
}
