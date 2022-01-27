package services

import (
	"time"

	"go-clean-architecture/pkg/dto"
	"go-clean-architecture/pkg/entity"
	authRepo "go-clean-architecture/pkg/repository/auth"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func NewAuthService(authRepository *authRepo.AuthRepository) AuthService {
	return &authServiceImpl{
		authRepository: *authRepository,
	}
}

type authServiceImpl struct {
	authRepository authRepo.AuthRepository
}

func (service *authServiceImpl) Login(userCredentials dto.UserCredential) (entity.Customer, error) {
	customer, err := service.authRepository.Login(userCredentials)
	return customer, err
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
