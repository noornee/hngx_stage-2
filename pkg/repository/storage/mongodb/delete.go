package mongodb

import (
	"context"
	"fmt"
)

func (db *Database) DeleteOneFromDb(collectionName string, filter map[string]any) error {
	collection := db.GetCollection(collectionName)
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	fmt.Println(result.DeletedCount)
	return nil
}
