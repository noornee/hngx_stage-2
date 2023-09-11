package mongodb

import (
	"context"
	"fmt"
)

func (db *Database) CreateOneRecord(collectionName string, model any) error {
	collection := db.GetCollection(collectionName)
	result, err := collection.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}

	fmt.Println("inserted id:", result.InsertedID)

	return nil
}
