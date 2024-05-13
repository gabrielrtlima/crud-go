package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Password string             `bson:"password,omitempty""`
	Age      int8               `bson:"age,omitempty"`
}
