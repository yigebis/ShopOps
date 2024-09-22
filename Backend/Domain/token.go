package Domain

import (
	"time"
)

type RefresherTable struct{
	Email string `json:"email" bson:"email"`
	Refresher string `json:"refresher" bson:"refresher"`
}

type LogoutTable struct{
	Token string `json:"token" bson:"token"`
	ExpiryTime time.Time `json:"expiry_time" bson:"expiry_time"`
}