package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBName        = "HNGX_STAGE_2"
	ConnectionURI = "mongodb://localhost:27017"
)

type Database struct {
	DB *mongo.Database
}

// connect to mongodb
func ConnectToDB() (*Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(ConnectionURI))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	db := Database{DB: client.Database(DBName)}
	return &db, nil
}

func (db *Database) GetCollection(collectionName string) *mongo.Collection {
	return db.DB.Collection(collectionName)
}
