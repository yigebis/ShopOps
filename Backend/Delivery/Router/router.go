package Router

import (
	"ShopOps/Delivery/Controller"

	"github.com/gin-gonic/gin"
)

type Router struct{
	UserController *Controller.UserController
}

func NewRouter(uc *Controller.UserController) *Router{
	return &Router{
		UserController: uc,
	}
}

func (r *Router) Run(){
	router := gin.Default()


	router.POST("/register", r.UserController.Register)
	router.POST("/login", r.UserController.Login)
	router.GET("/verify", r.UserController.VerifyEmail)

	router.Run()
}