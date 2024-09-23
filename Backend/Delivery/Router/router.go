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
	router.POST("/verify", r.UserController.VerifyEmail)

	router.GET("/employees", r.UserController.GetAllEmployees)
	router.GET("/employee/:email", r.UserController.GetEmployee)
	router.POST("/employee/add", r.UserController.AddEmployee)
	router.POST("/employee/edit", r.UserController.EditEmployee)
	router.POST("/employee/delete/:email", r.UserController.DeleteEmployee)

	router.Run()
}