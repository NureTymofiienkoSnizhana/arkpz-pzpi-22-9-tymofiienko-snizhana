package mongodb

import (
	"context"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const HealthDataCollectionName = "HealthDatas"

type healthDataDB struct {
	collection *mongo.Collection
}

func newHealthDataDB(db *mongo.Database) *healthDataDB {
	return &healthDataDB{
		collection: db.Collection(HealthDataCollectionName),
	}
}

func NewHealthDataDB(db *mongo.Database) data.HealthDataDB {
	return newHealthDataDB(db)
}

func (hd *healthDataDB) Get(id string) (*data.HealthData, error) {
	var result data.HealthData
	err := hd.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (hd *healthDataDB) Insert(healthData *data.HealthData) error {
	_, err := hd.collection.InsertOne(context.TODO(), healthData)
	return err
}

func (hd *healthDataDB) Update(id string, updateFields bson.M) error {
	_, err := hd.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": updateFields},
	)
	return err
}
