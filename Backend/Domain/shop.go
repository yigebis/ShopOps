package Domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct{
	ID primitive.ObjectID
	EmployeeCount int
	OwnerEmail string
	Adress string

}