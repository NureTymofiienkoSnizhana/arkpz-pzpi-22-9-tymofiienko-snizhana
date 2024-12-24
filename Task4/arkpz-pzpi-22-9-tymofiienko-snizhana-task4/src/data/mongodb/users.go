package mongodb

import (
	"context"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

func (u *usersDB) Update(userID primitive.ObjectID, updateFields bson.M) error {
	result, err := u.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID},
		updateFields,
	)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	if result.MatchedCount == 0 {
		log.Printf("No user found with ID: %v", userID.Hex())
		return mongo.ErrNoDocuments
	}
	return nil
}

func (u *usersDB) UpdatePets(userID primitive.ObjectID, petID primitive.ObjectID) error {
	_, err := u.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID, "$or": []bson.M{
			{"pets_id": bson.M{"$exists": false}},
			{"pets_id": nil},
		}},
		bson.M{"$set": bson.M{"pets_id": []primitive.ObjectID{}}},
	)
	if err != nil {
		log.Printf("Error initializing pets_id: %v", err)
		return err
	}

	update := bson.M{
		"$addToSet": bson.M{"pets_id": petID},
	}

	_, err = u.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID},
		update,
	)
	if err != nil {
		log.Printf("Error updating user pets: %v", err)
		return err
	}

	return nil
}

func (u *usersDB) FindByEmail(email string) (*data.User, error) {
	var user data.User
	err := u.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersDB) Delete(userID primitive.ObjectID) error {
	result, err := u.collection.DeleteOne(
		context.TODO(),
		bson.M{"_id": userID},
	)

	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	if result.DeletedCount == 0 {
		log.Printf("No user found with ID: %v", userID.Hex())
		return mongo.ErrNoDocuments
	}
	return nil
}