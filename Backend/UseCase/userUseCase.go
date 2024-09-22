package UseCase

import (
	"ShopOps/Domain"
	"fmt"

	"strconv"
	"time"
)

type UserUseCase struct{
	UserRepo IUserRepository
	PasswordService IPasswordService
	TokenRepo ITokenRepository
	TokenService ITokenService
	MailService IMailService
	ErrorService IErrorService

	EmailExpiry string
	TokenExpiry string
	RefresherExpiry string
}

func NewUserUseCase(ur IUserRepository, ps IPasswordService, tr ITokenRepository, ts ITokenService, ms IMailService, es IErrorService, ex, tx, rx string) IUserUseCase{	
	return &UserUseCase{
		UserRepo: ur,
		PasswordService: ps,
		TokenRepo: tr,
		TokenService: ts,
		MailService: ms,
		ErrorService: es,
		EmailExpiry: ex,
		TokenExpiry: tx,
		RefresherExpiry: rx,
	}
}

func (uuc *UserUseCase) Register(user *Domain.User) (int, error){
	// check if the user with this email has already been registered
	existingUser, err := uuc.UserRepo.GetUserByEmail(user.Email)
	
	if err == nil{
		if !existingUser.Verified{
			return uuc.ErrorService.PendingVerification()
		}
		
		return uuc.ErrorService.UserExists()
	}

	// Set verified field false
	user.Verified = false

	//hash the password
	hashedPassword, err := uuc.PasswordService.HashPassword(user.Password)
	if err != nil{
		fmt.Println("hashing")
		return uuc.ErrorService.InternalServer()
	}

	user.Password = hashedPassword
	user.RegistrationDate = time.Now()
	user.Role = "owner"
	user.OwnerEmail = user.Email
	user.ShopCount = 1

	// store the user in the database
	err = uuc.UserRepo.CreateUser(user)
	if err != nil{
		fmt.Println("repo")
		return uuc.ErrorService.InternalServer()
	}

	//set the expiry time for accepting email verification
	seconds, _ := strconv.Atoi(uuc.EmailExpiry)
	expiryDuration := time.Now().Add(time.Second * time.Duration(seconds)).Unix()

	// send verification email
	token, err := uuc.TokenService.GenerateToken(user.Email, user.FirstName, expiryDuration)
	if err != nil {
		fmt.Println("token_mail")
		return uuc.ErrorService.InternalServer()
	}

	err = uuc.MailService.SendVerificationEmail(user.Email, token)

	if err != nil {
		fmt.Println("verification")
		return uuc.ErrorService.InternalServer()
	}

	return uuc.ErrorService.NoError()
}

func (uuc *UserUseCase) VerifyEmail(email, token string) (int, error){
	// verify the token and the email
	claims, err := uuc.TokenService.ValidateToken(token)
	if err != nil{
		return uuc.ErrorService.InvalidToken()
	}

	if claims["email"] != email{
		return uuc.ErrorService.InvalidToken()
	}

	// check if the user is already verified
	fmt.Println(email, token)
	user, err := uuc.UserRepo.GetUserByEmail(email)
	if err != nil{
		return uuc.ErrorService.InternalServer()
	}

	if user.Verified{
		return uuc.ErrorService.UserExists()
	}

	user.Verified = true

	err = uuc.UserRepo.VerifyUser(user)

	if err != nil{
		return uuc.ErrorService.InternalServer()
	}

	return uuc.ErrorService.NoError()
}

func (uuc *UserUseCase) LoginByEmail(email, password string) (string, string, int, error){
	// try to get the user with this email
	user, err := uuc.UserRepo.GetUserByEmail(email)

	if err != nil{
		code, err := uuc.ErrorService.InvalidEmailPassword()
		return "", "", code, err
	}

	//check whether the user is verified or not
	if !user.Verified{
		code, err := uuc.ErrorService.NotVerified()
		return "", "", code, err
	}
	// verify the password
	err = uuc.PasswordService.VerifyPassword(user.Password, password)
	if err != nil{
		code, err := uuc.ErrorService.InvalidEmailPassword()
		return "", "", code, err
	}

	//generate a new token and refresher for the user
	token_seconds, _ := strconv.Atoi(uuc.TokenExpiry)
	tokenExpiry := time.Now().Add(time.Second * time.Duration(token_seconds)).Unix()

	refresher_seconds, _ := strconv.Atoi(uuc.RefresherExpiry)
	refresherExpiry := time.Now().Add(time.Second * time.Duration(refresher_seconds)).Unix()

	token, err := uuc.TokenService.GenerateToken(email, user.FirstName, tokenExpiry)
	if err != nil{
		code, err := uuc.ErrorService.InternalServer()
		return "", "", code, err
	}

	refresher, err := uuc.TokenService.GenerateToken(email, user.FirstName, refresherExpiry)
	if err != nil{
		code, err := uuc.ErrorService.InternalServer()
		return "", "", code, err
	}

	//store the refresher token
	err = uuc.TokenRepo.InsertRefresher(email, refresher)
	if err != nil{
		code, err := uuc.ErrorService.InternalServer()
		return "", "", code, err
	}

	code, err := uuc.ErrorService.NoError()
	return token, refresher, code, err
}

func (uuc *UserUseCase) LoginByPhone(phone, password string) (string, string, int, error){
	// try to get the user with this phone
	user, err := uuc.UserRepo.GetUserByPhone(phone)

	if err != nil{
		code, err := uuc.ErrorService.InvalidPhonePassword()
		return "", "", code, err
	}

	//check whether the user is verified or not
	if !user.Verified{
		code, err := uuc.ErrorService.NotVerified()
		return "", "", code, err
	}
	// verify the password
	err = uuc.PasswordService.VerifyPassword(user.Password, password)
	if err != nil{
		code, err := uuc.ErrorService.InvalidPhonePassword()
		return "", "", code, err
	}

	//generate a new token and refresher for the user
	token_seconds, _ := strconv.Atoi(uuc.TokenExpiry)
	tokenExpiry := time.Now().Add(time.Second * time.Duration(token_seconds)).Unix()

	refresher_seconds, _ := strconv.Atoi(uuc.RefresherExpiry)
	refresherExpiry := time.Now().Add(time.Second * time.Duration(refresher_seconds)).Unix()

	token, err := uuc.TokenService.GenerateToken(user.Email, user.FirstName, tokenExpiry)
	if err != nil{
		code, err := uuc.ErrorService.InternalServer()
		return "", "", code, err
	}

	refresher, err := uuc.TokenService.GenerateToken(user.Email, user.FirstName, refresherExpiry)
	if err != nil{
		code, err := uuc.ErrorService.InternalServer()
		return "", "", code, err
	}

	//store the refresher token
	err = uuc.TokenRepo.InsertRefresher(phone, refresher)
	if err != nil{
		code, err := uuc.ErrorService.InternalServer()
		return "", "", code, err
	}

	code, err := uuc.ErrorService.NoError()
	return token, refresher, code, err
}

func (uuc *UserUseCase) GetSingleUser(email string) (*Domain.User, int, error){
	user, err := uuc.UserRepo.GetUserByEmail(email)
	if err != nil{
		code, err := uuc.ErrorService.UserNotFound()
		return nil, code, err
	}

	code, err := uuc.ErrorService.NoError()
	return user, code, err
}

func (uuc *UserUseCase) RefreshToken(email, refresher string) (string, int, error){
	//check the existence of refresher
	err := uuc.TokenRepo.CheckRefresher(email, refresher)
	if err != nil{
		code, err := uuc.ErrorService.InvalidEmailRefresher()
		return "", code, err
	}

	// generate a new token
	claims, err := uuc.TokenService.ValidateToken(refresher)
	if err != nil {
		code, err := uuc.ErrorService.InternalServer()
		return "", code, err
	}

	firstName := claims["firstName"]
	token_seconds, _ := strconv.Atoi(uuc.TokenExpiry)
	tokenExpiry := time.Now().Add(time.Second * time.Duration(token_seconds)).Unix()

	token, err := uuc.TokenService.GenerateToken(email, firstName.(string), tokenExpiry)
	if err != nil {
		code, err := uuc.ErrorService.InternalServer()
		return "", code, err
	}
	code, err := uuc.ErrorService.NoError()
	return token, code, err
}

func (uuc *UserUseCase) GenerateResetPasswordToken(email string) (int, error){
	return 0, nil
}

func (uuc *UserUseCase) ResetPassword(token string, newPassword string) (int, error){
	return 0, nil
}

func (uuc *UserUseCase) StoreToken(token string) (int, error){
	return 0, nil
}

// please consider this
func (uuc *UserUseCase) Logout(email, token, refresher string) (int, error){
	//Delete the refresher token
	err := uuc.TokenRepo.DeleteRefresher(email, refresher )
	if err != nil{
		return uuc.ErrorService.InvalidEmailRefresher()
	}

	//Add the normal token to the logged out tokens list
	err = uuc.TokenRepo.LogoutToken(token)
	if err != nil {
		return uuc.ErrorService.InternalServer()
	}

	return uuc.ErrorService.NoError()
}

//Only for admins
func (uuc *UserUseCase) DeleteUser(email string) (int, error){
	err := uuc.UserRepo.DeleteUser(email)
	if err != nil{
		return uuc.ErrorService.UserNotFound();
	}

	return uuc.ErrorService.NoError();
}