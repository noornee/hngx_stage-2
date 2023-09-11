package models

import (
	"fmt"

	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CollectionName string = "users"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
}

func (u *User) CreateUser(db *mongodb.Database) error {
	err := db.CreateOneRecord(CollectionName, &u)
	if err != nil {
		return fmt.Errorf("there was an error creating user %s", err.Error())
	}
	return nil
}

func (u *User) GetUserByUsername(db *mongodb.Database) error {
	return nil
}

func (u *User) UpdateUserByUsername(db *mongodb.Database) error {
	return nil
}

func (u *User) DeleteUserByUsername(db *mongodb.Database) error {
	return nil
}
