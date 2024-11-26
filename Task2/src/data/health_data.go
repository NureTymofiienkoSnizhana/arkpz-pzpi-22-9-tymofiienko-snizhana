package data

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HealthDataDB interface {
	Get(id string) (*HealthData, error)
	Insert(*HealthData) error
	Update(id string, updateFields bson.M) error
}

type HealthData struct {
	ID       primitive.ObjectID `bson:"_id"`
	PetID    primitive.ObjectID `bson:"pet_id"`
	Activity string             `bson:"activity"`
	Sleep    string             `bson:"sleep"`
	Feeding  string             `bson:"feeding"`
	Time     time.Time          `bson:"time"`
}
