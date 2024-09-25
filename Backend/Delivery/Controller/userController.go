package Controller

import (
	"ShopOps/Domain"
	"ShopOps/UseCase"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct{
	UserUseCase UseCase.IUserUseCase
	V *validator.Validate
	TokenService UseCase.ITokenService
}

func NewUserController(uuc UseCase.IUserUseCase, ts UseCase.ITokenService) *UserController{
	return &UserController{
		UserUseCase: uuc,
		V: validator.New(),
		TokenService: ts,
	}
}

func (c *UserController) Register(ctx *gin.Context){
	user := Domain.User{}

	err := ctx.ShouldBindJSON(&user)

	if err != nil{
		ctx.JSON(400, gin.H{"error" : "invalid request payload"})
		return
	}

	err = c.V.Struct(user)
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

	var token, refresher string
	var code int

	if credential.Email != "" &&credential.Password != ""{
		token, refresher, code, err = c.UserUseCase.LoginByEmail(credential.Email, credential.Password)
	}else{
		ctx.JSON(code, gin.H{"error" : "email and password are required"})
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

// func (c *UserController) EditUser(ctx *gin.Context){
// 	var employee = Domain.User{}
// 	if err := ctx.ShouldBindJSON(&employee); err != nil{
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request payload"})
// 		return
// 	}

// 	if err := c.V.Struct(&employee); err != nil{
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request payload"})
// 		return
// 	}

// 	code, err := c.UserUseCase.EditUser(&employee)
// 	if err != nil{
// 		ctx.JSON(code, gin.H{"error" : err.Error()})
// 		return
// 	}

// 	ctx.JSON(code, gin.H{"message" : "user updated successfully"})
// }

func (c *UserController) GetHeader(ctx *gin.Context) (map[string]interface{}, error){
	authHeader := ctx.GetHeader("Authorization")
	parts := strings.Split(authHeader, " ")
	claims, err := c.TokenService.ValidateToken(parts[1])
	return claims, err
}
func (c *UserController) GetAllEmployees(ctx *gin.Context){
	claims, err := c.GetHeader(ctx)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error" : "unauthorized"})
		return
	}

	ownerEmail := claims["email"].(string)

	employees, code, err := c.UserUseCase.GetAllEmployees(ownerEmail)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, employees)
}

func (c *UserController) GetEmployee(ctx *gin.Context){
	claims, err := c.GetHeader(ctx)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error" : "unauthorized"})
		return
	}

	ownerEmail := claims["email"].(string)

	email := ctx.Param("email")
	employee, code, err := c.UserUseCase.GetEmployee(email, ownerEmail)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, employee)
}

func (c *UserController) AddEmployee(ctx *gin.Context){
	claims, err := c.GetHeader(ctx)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error" : "unauthorized"})
		return
	}

	ownerEmail := claims["email"].(string)

	var employee = Domain.User{}
	if err := ctx.ShouldBindJSON(&employee); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request payload"})
		return
	}

	if err := c.V.Struct(&employee); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request payload"})
		return
	}

	code, err := c.UserUseCase.AddEmployee(&employee, ownerEmail)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "employee added successfully"})
}

func (c *UserController) EditEmployee(ctx *gin.Context){
	claims, err := c.GetHeader(ctx)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error" : "unauthorized"})
		return
	}

	ownerEmail := claims["email"].(string)
	var employee = Domain.User{}

	if err := ctx.ShouldBindJSON(&employee); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"errro" : "invalid request payload"})
		return
	}

	if err := c.V.Struct(employee); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request payload"})
		return
	}

	code, err := c.UserUseCase.EditEmployee(&employee, ownerEmail)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "employee updated successfully"})
}

func (c *UserController) DeleteEmployee(ctx *gin.Context){
	email := ctx.Param("email")
	claims, err := c.GetHeader(ctx)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error" : "unauthorized"})
		return
	}

	ownerEmail := claims["email"].(string)

	code, err := c.UserUseCase.DeleteEmployee(email, ownerEmail)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "employee deleted successfully"})
}

func (c *UserController) ActivateAccount(ctx *gin.Context){
	var activationCredential = Domain.ActivationCredential{}
	err := ctx.ShouldBindJSON(&activationCredential)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request payload"})
		return
	}

	code, err := c.UserUseCase.ActivateAccount(activationCredential.Email, activationCredential.OldPassword, activationCredential.NewPassword)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" :"account activation successful"})
}





