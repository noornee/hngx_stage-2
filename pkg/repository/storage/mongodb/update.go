package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *Database) UpdateAllFields(collectionName string, model any, filter map[string]any) error {
	collection := db.GetCollection(collectionName)

	update := bson.M{"$set": model}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("match count:", result.ModifiedCount)

	return nil
}
