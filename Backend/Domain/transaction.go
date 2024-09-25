package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SellTransaction struct{
	ID primitive.ObjectID
	ShopID primitive.ObjectID
	// BranchID
	EmployeeID primitive.ObjectID
	ProductID primitive.ObjectID
	Date time.Time               
	TotalQuantity int           
	PaymentMethod string         //Cash or Bank
}

type BuyTransaction struct{
	ID primitive.ObjectID
	ShopID primitive.ObjectID
	// BranchID
	ProductID primitive.ObjectID
	TotalQuantity int
	Date time.Time
	PaymentMethod string
}