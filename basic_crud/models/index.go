package models

type CustomerDetails struct {
	Id       int32   `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Email    string  `json:"email" bson:"email"`
	Password string  `json:"password" bson:"password"`
	Salary   float32 `json:"salary" bson:"salary"`
	Address  string  `json:"address" bson:"address"`
}
