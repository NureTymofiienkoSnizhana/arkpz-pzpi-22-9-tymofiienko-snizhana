package data

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HealthDataDB interface {
	Get(pet_id primitive.ObjectID) (*HealthData, error)
	Insert(*HealthData) error
	Update(pet_id primitive.ObjectID, updateFields bson.M) error
}

type HealthData struct {
	ID       primitive.ObjectID `bson:"_id"`
	PetID    primitive.ObjectID `bson:"pet_id"`
	Activity string             `bson:"activity"`
	Sleep    string             `bson:"sleep"`
	Feeding  string             `bson:"feeding"`
	Time     string             `bson:"time"`
}
