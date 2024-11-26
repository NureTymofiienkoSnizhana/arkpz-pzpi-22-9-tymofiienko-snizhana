package data

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PetsDB interface {
	Get(id primitive.ObjectID) (*Pet, error)
	Insert(*Pet) error
	Update(id primitive.ObjectID, updateFields bson.M) error
	Delete(id primitive.ObjectID) error
	GetAll() ([]*Pet, error)
}

type Pet struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Speices string             `bson:"speices"`
	Breed   string             `bson:"breed"`
	Age     int                `bson:"age"`
	OwnerID primitive.ObjectID `bson:"owner_id"`
}
