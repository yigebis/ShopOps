package Router

import (
	"ShopOps/Delivery/Controller"
	"ShopOps/Infrastructure"

	"github.com/gin-gonic/gin"
)

type Router struct{
	UserController *Controller.UserController
	JWTSigner string
}

func NewRouter(uc *Controller.UserController, jwtSigner string) *Router{
	return &Router{
		UserController: uc,
		JWTSigner: jwtSigner,
	}
}

func (r *Router) Run(){
	router := gin.Default()


	router.POST("/register", r.UserController.Register)
	router.POST("/login", r.UserController.Login)
	router.POST("/verify", r.UserController.VerifyEmail)

	router.GET("/employees", Infrastructure.UserMiddleware(r.JWTSigner), Infrastructure.OwnerMiddleWare(r.JWTSigner), r.UserController.GetAllEmployees)
	router.GET("/employee/:email", Infrastructure.UserMiddleware(r.JWTSigner), Infrastructure.OwnerMiddleWare(r.JWTSigner), r.UserController.GetEmployee)
	router.POST("/employee/add", Infrastructure.UserMiddleware(r.JWTSigner), Infrastructure.OwnerMiddleWare(r.JWTSigner), r.UserController.AddEmployee)
	router.PUT("/employee/edit", Infrastructure.UserMiddleware(r.JWTSigner), Infrastructure.OwnerMiddleWare(r.JWTSigner), r.UserController.EditEmployee)
	router.POST("/employee/delete/:email", Infrastructure.UserMiddleware(r.JWTSigner), Infrastructure.OwnerMiddleWare(r.JWTSigner), r.UserController.DeleteEmployee)

	router.POST("/activate", r.UserController.ActivateAccount)

	router.Run()
}