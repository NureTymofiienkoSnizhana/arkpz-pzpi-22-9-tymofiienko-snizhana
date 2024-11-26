package mongodb

import (
	"context"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollectionName = "Users"

type usersDB struct {
	collection *mongo.Collection
}

func newUsersDB(db *mongo.Database) *usersDB {
	return &usersDB{
		collection: db.Collection(UsersCollectionName),
	}
}

func NewUsersDB(db *mongo.Database) data.UsersDB {
	return newUsersDB(db)
}

func (u *usersDB) Get(id primitive.ObjectID) (*data.User, error) {
	var result data.User
	err := u.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (u *usersDB) Insert(user *data.User) error {
	_, err := u.collection.InsertOne(context.TODO(), user)
	return err
}

func (u *usersDB) Update(id primitive.ObjectID, updateFields bson.M) error {
	_, err := u.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": updateFields},
	)
	return err
}

func (u *usersDB) FindByEmail(email string) (*data.User, error) {
	var user data.User
	err := u.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
