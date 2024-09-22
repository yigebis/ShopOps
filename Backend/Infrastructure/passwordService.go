package Infrastructure

import (
	"ShopOps/UseCase"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func NewPasswordService() UseCase.IPasswordService {
	return &PasswordService{}
}

// HashPassword implements UseCase.IPasswordService.
func (p *PasswordService) HashPassword(password string) (string, error) {
	hashedPasswordSlice, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPassword := string(hashedPasswordSlice)
	return hashedPassword, nil
}

// VerifyPassword implements UseCase.IPasswordService.
func (p *PasswordService) VerifyPassword(hashedPassword string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
