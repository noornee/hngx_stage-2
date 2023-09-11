package mongodb

import (
	"context"
	"log"

	"github.com/noornee/hngx_stage-2/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var env = config.Setup()

var (
	DBName        = env.DB_NAME
	ConnectionURI = env.CONNECTION_URI
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
