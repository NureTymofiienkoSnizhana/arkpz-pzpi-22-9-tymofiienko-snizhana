package mongodb

import (
	"context"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DevicesCollectionName = "Devices"

type devicesDB struct {
	collection *mongo.Collection
}

func newDevicesDB(db *mongo.Database) *devicesDB {
	return &devicesDB{
		collection: db.Collection(DevicesCollectionName),
	}
}

func NewDevicesDB(db *mongo.Database) data.DevicesDB {
	return newDevicesDB(db)
}

func (d *devicesDB) Get(id primitive.ObjectID) (*data.Device, error) {
	var result data.Device
	err := d.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (d *devicesDB) Insert(device *data.Device) error {
	_, err := d.collection.InsertOne(context.TODO(), device)
	return err
}

func (d *devicesDB) Update(id primitive.ObjectID, updateFields bson.M) error {
	_, err := d.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": updateFields},
	)
	return err
}
