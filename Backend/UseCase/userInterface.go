package UseCase

import (
	"ShopOps/Domain"
)


type IUserUseCase interface {
	LoginByEmail(email, password string) (string, string, int, error)
	LoginByPhone(phone, password string) (string, string, int, error)
	GetSingleUser(email string) (*Domain.User, int, error)
	Register(user *Domain.User) (int, error)
	VerifyEmail(email, token string) (int, error)
	
	RefreshToken(email, refresher string) (string, int, error)
	GenerateResetPasswordToken(email string) (int, error)
	ResetPassword(token string, newPassword string) (int, error)
	StoreToken(token string) (int, error)
	Logout(email, token, refresher string) (int, error)
	DeleteUser(email string) (int, error) //Only for admins
}

type IUserRepository interface {
	CreateUser(user *Domain.User) error
	GetUserByEmail(email string) (*Domain.User, error)
	GetUserByPhone(phone string) (*Domain.User, error)
	VerifyUser(user *Domain.User) error
	GetUserByVerificationToken(token string) (*Domain.User, error)
	GetUserCount() (int64, error)
	UpdatePasswordByEmail(email string, newPassword string) error
	StoreResetToken(email string, resetToken string) error
	InvalidateResetToken(email string) error
	GetResetTokenByEmail(email string) (string, error)

	DeleteUser(email string) error
}

type IPasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, plainPassword string) error
}
type ITokenService interface {
	GenerateToken(email, firstName string, expiryDuration int64) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}

type ITokenRepository interface {
	InsertRefresher(email, refresher string) error
	CheckRefresher(email, refresher string) error
	InvalidateResetToken(email string) error
	StoreResetToken(email string, resetToken string) error
	UpdateRefresher(email, refresher string) error
	DeleteRefresher(email, refresher string) error
	DeleteAllRefreshers(email string) error
	LogoutToken(token string) error
}

type IMailService interface {
	SendVerificationEmail(to, token string) error
	SendPasswordResetEmail(to ,resetToken string) error
}

