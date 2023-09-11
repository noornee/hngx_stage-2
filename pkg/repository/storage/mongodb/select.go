package mongodb

import (
	"context"
)

func (db *Database) SelectOneFromDb(collectionName string, result any, filter map[string]any) error {
	collection := db.GetCollection(collectionName)
	return collection.FindOne(context.Background(), filter).Decode(result)
}
