package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName          string             `json:"first_name" bson:"first_name" validate:"required,min=1,max=50"`
	LastName           string             `json:"last_name" bson:"last_name" validate:"required,min=1,max=50"`
	Sex				   string  			  `json:"sex" bson:"sex" validate:"required,min=1,max=1"`
	PhoneNumber		   string             `json:"phone_number" bson:"phone_number" validate:"required"`
	Email              string             `json:"email" bson:"email" validate:"required,email"`
	Password           string             `json:"password" bson:"password" validate:"required"`
	Role               string             `json:"role" bson:"role"`
	OwnerEmail         string             `json:"owner_email" bson:"owner_email"`
	ProfilePhoto       string             `json:"profile_photo" bson:"profile_photo"`
	RegistrationDate   time.Time          `json:"registration_date" bson:"registration_date"`
	ShopCount          int                `json:"shop_count" bson:"shop_count"`
	Verified           bool               `json:"verified" bson:"verified"`
}

type Credential struct{
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	// PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type ActivationCredential struct{
	Email string `json:"email" bson:"email"`
	OldPassword string `json:"old_password" bson:"old_password"`
	NewPassword string `json:"new_password" bson:"new_password"`
}