package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *Database) CreateOneRecord(collectionName string, model any) (any, error) {
	collection := db.GetCollection(collectionName)
	result, err := collection.InsertOne(context.Background(), model)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": result.InsertedID.(primitive.ObjectID)}
	update := bson.M{"$set": bson.M{"created_at": time.Now(), "updated_at": time.Now()}}

	// add the newly inserted timestamps
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}
