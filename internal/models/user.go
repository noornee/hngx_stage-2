package models

import (
	"fmt"
	"time"

	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var CollectionName string = "users"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

func (u *User) CreateUser(db *mongodb.Database) (*User, error) {
	insertedID, err := db.CreateOneRecord(CollectionName, &u)
	if err != nil {
		return u, fmt.Errorf("there was an error creating user %s", err.Error())
	}

	id, _ := insertedID.(primitive.ObjectID)
	u.ID = id

	// retrieve user from db
	u.GetUserByID(db)

	return u, nil
}

func (u *User) GetUserByID(db *mongodb.Database) error {
	err := db.SelectOneFromDb(CollectionName, &u, bson.M{"_id": u.ID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user with specified id does not exist")
		}
		return err
	}
	return nil
}

func (u *User) UpdateUserByID(db *mongodb.Database) error {
	err := db.UpdateAllFields(CollectionName, &u, bson.M{"_id": u.ID})
	if err != nil {
		return fmt.Errorf("there was an error updating user %s", err.Error())
	}
	return nil
}

func (u *User) DeleteUserByID(db *mongodb.Database) error {
	err := db.DeleteOneFromDb(CollectionName, bson.M{"_id": u.ID})
	if err != nil {
		return fmt.Errorf("there was an error deleting user %s", err.Error())
	}
	return nil
}
