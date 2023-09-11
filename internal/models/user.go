package models

import (
	"fmt"

	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CollectionName string = "users"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Gender   string             `bson:"gender" json:"gender"`
	Age      int                `bson:"age" json:"age"`
	AboutMe  string             `bson:"about_me" json:"about_me"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	AboutMe  string `json:"about_me"`
}

func (u *User) CreateUser(db *mongodb.Database) error {
	err := db.CreateOneRecord(CollectionName, &u)
	if err != nil {
		return fmt.Errorf("there was an error creating user %s", err.Error())
	}
	return nil
}

func (u *User) GetUserByID(db *mongodb.Database) error {
	err := db.SelectOneFromDb(CollectionName, &u, bson.M{"_id": u.ID})
	if err != nil {
		return err
	}
	return err
}

func (u *User) GetUserByUsername(db *mongodb.Database) error {
	err := db.SelectOneFromDb(CollectionName, &u, bson.M{"username": u.Username})
	if err != nil {
		return err
	}
	return err
}

func (u *User) UpdateUserByUsername(db *mongodb.Database) error {
	err := db.UpdateAllFields(CollectionName, &u, bson.M{"username": u.Username})
	if err != nil {
		return fmt.Errorf("there was an error updating user %s", err.Error())
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
