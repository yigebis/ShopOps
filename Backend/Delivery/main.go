package main

import (
	"ShopOps/Delivery/Controller"
	"ShopOps/Delivery/Router"
	"ShopOps/Error"
	"ShopOps/Infrastructure"
	"ShopOps/Repository"
	"ShopOps/UseCase"

	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
	if err := godotenv.Load(); err != nil{
		log.Fatalf("error loading .env file")
	}

	// setting up the usecase

	//user usecase
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	uri := "mongodb+srv://" + username + ":" + password + "@cluster0.isgee.mongodb.net/"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to db!")
	user_collection := client.Database("ShopOps").Collection("users")
	user_context := context.TODO()
	ur := Repository.NewUserRepository(user_context, user_collection)

	ps := Infrastructure.NewPasswordService()

	token_collection := client.Database("ShopOps").Collection("refreshers")
	token_context := context.TODO()
	tr := Repository.NewTokenRepository(token_context, token_collection)

	jwtSecret := os.Getenv("JWT_SECRET")
	ts := Infrastructure.NewTokenService(jwtSecret)
	ms := Infrastructure.NewMailService(os.Getenv("SENDER_EMAIL"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("FROM"))
	es := Error.NewErrorService()
	
	ex := os.Getenv("EMAIL_EXPIRY")
	tx := os.Getenv("TOKEN_EXPIRY")
	rx := os.Getenv("REFRESHER_EXPIRY")

	uuc := UseCase.NewUserUseCase(ur, ps, tr, ts, ms, es, ex, tx, rx)


	// setting up the controllers
	user_controller := Controller.NewUserController(uuc, ts)

	// setting up the router
	router := Router.NewRouter(user_controller)
	router.Run()
}