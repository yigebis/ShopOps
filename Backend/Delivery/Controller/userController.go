package Controller

import (
	"ShopOps/Domain"
	"ShopOps/UseCase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct{
	UserUseCase UseCase.IUserUseCase
	v *validator.Validate
}

func NewUserController(uuc UseCase.IUserUseCase) *UserController{
	return &UserController{
		UserUseCase: uuc,
		v: validator.New(),
	}
}

func (c *UserController) Register(ctx *gin.Context){
	user := Domain.User{}

	err := ctx.ShouldBindJSON(&user)

	if err != nil{
		ctx.JSON(400, gin.H{"error" : "invalid request payload"})
		return
	}

	err = c.v.Struct(user)
	if err != nil{
		ctx.JSON(400, gin.H{"error" : "invalid request payload"})
		return
	}

	code, err := c.UserUseCase.Register(&user)

	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "registration successful. Verification has been sent to the Email"})
}

func (c *UserController) VerifyEmail(ctx *gin.Context){
	email := ctx.Query("email")
	token := ctx.Query("token")

	code, err := c.UserUseCase.VerifyEmail(email, token)

	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}


	ctx.JSON(code, gin.H{"message" : "email verified successfully"})
}

func (c *UserController) Login(ctx *gin.Context){
	credential := Domain.Credential{}
	err := ctx.ShouldBindJSON(&credential)

	if err != nil{
		ctx.JSON(400, gin.H{"error" : "invalid request payload"})
		return
	}

	if credential.Email != "" && credential.PhoneNumber != ""{
		ctx.JSON(400, gin.H{"error" : "either email or phone_number are required"})
		return
	}

	var token, refresher string
	var code int

	if credential.Email != ""{
		token, refresher, code, err = c.UserUseCase.LoginByEmail(credential.Email, credential.Password)
	}else if credential.PhoneNumber != ""{
		token, refresher, code, err = c.UserUseCase.LoginByPhone(credential.PhoneNumber, credential.Password)
	}else{
		ctx.JSON(code, gin.H{"error" : "either email or phone_number are required"})
		return
	}
	

	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{
		"token" : token,
		"refresher" : refresher,
	})
}




