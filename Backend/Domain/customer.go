package Domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct{
	ID primitive.ObjectID
	FirstName string
	LastName string
	Sex string
	
}