package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct{
	ID primitive.ObjectID
	ShopID primitive.ObjectID
	EmployeeID primitive.ObjectID
	ProductID primitive.ObjectID
	CustomerID primitive.ObjectID
	Date time.Time               
	TotalQuantity int           
	PaymentMethod string         //Cash or Bank
	Status string                //S(Sold) or B(Bought)
}