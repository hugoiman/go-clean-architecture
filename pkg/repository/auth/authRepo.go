package repo

import "go-clean-architecture/pkg/dto"

type AuthRepository interface {
	Login(userCredentials dto.UserCredential) error
}
