package Domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{
	ID primitive.ObjectID
	Name string
	Description string
	UnitPrice float64
	Quantity int
	Category string
}