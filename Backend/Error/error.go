package Error

import (
	"ShopOps/UseCase"

	"errors"
	"net/http"
)

var ErrInternalServer = errors.New("internal server error")

var ErrUserExists = errors.New("user already exists")
var ErrPendingVerification = errors.New("registration is waiting email verification")
var ErrInvalidToken = errors.New("invalid token")
var ErrUserNotFound = errors.New("not found")
var ErrInvalidEmailPassword = errors.New("invalid email or password")
var ErrInvalidPhonePassword = errors.New("invalid phone number or password")
var ErrInvalidEmailRefresher = errors.New("invalid email or refresher")
var ErrNotVerified = errors.New("unverified user")
var ErrNotActivated = errors.New("account not activated")
var ErrVerifiedOrNotEmployee = errors.New("user verified or not an employee")
var ErrSamePassword = errors.New("old and new password should be different")
var ErrNotAuthorized = errors.New("unauthorized")

type Error struct{}

func NewErrorService() UseCase.IErrorService{
	return &Error{}
} 

func (e *Error) NoError() (int, error){
	return http.StatusOK, nil
}

func (e *Error) UserExists() (int, error){
	return http.StatusConflict, ErrUserExists
}

func (e *Error) PendingVerification() (int, error){
	return http.StatusConflict, ErrPendingVerification
}

func (e *Error) InternalServer() (int, error){
	return http.StatusInternalServerError, ErrInternalServer
}

func (e *Error) InvalidToken() (int, error){
	return http.StatusBadRequest, ErrInvalidToken
}

func (e *Error) UserNotFound() (int, error){
	return http.StatusNotFound, ErrUserNotFound
}

func (e *Error) InvalidEmailPassword() (int, error){
	return http.StatusBadRequest, ErrInvalidEmailPassword
}

func (e *Error) InvalidPhonePassword() (int, error){
	return http.StatusBadRequest, ErrInvalidPhonePassword 
}

func (e *Error) InvalidEmailRefresher() (int, error){
	return http.StatusBadRequest, ErrInvalidEmailRefresher
}

func (e *Error) NotVerified() (int, error){
	return http.StatusBadRequest, ErrNotVerified
}

func (e *Error) NotActivated() (int, error){
	return http.StatusBadRequest, ErrNotActivated
}

func (e *Error) VerifiedOrNotEmploye() (int, error){
	return http.StatusBadRequest, ErrVerifiedOrNotEmployee
}

func (e *Error) SamePassword() (int, error){
	return http.StatusBadRequest, ErrSamePassword
}

func (e *Error) NotAuthorized() (int, error){
	return http.StatusUnauthorized, ErrNotAuthorized
}