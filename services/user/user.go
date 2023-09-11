package user

import (
	"fmt"
	"net/http"

	"github.com/noornee/hngx_stage-2/internal/models"
	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// check if user with the username already exist
func checkIfUserExist(db *mongodb.Database, username string) error {
	user := models.User{
		Username: username,
	}

	err := user.GetUserByUsername(db)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user does not exist")
		}
		return err
	}

	return nil
}

func CreateUserService(db *mongodb.Database, req models.CreateUserRequest) (int, error) {
	user := models.User{
		Username: req.Username,
		AboutMe:  fmt.Sprintf("My name is %s and I am Awesome!", req.Username),
	}

	// check if user with the username already exist
	if err := checkIfUserExist(db, req.Username); err == nil {
		return http.StatusBadRequest, fmt.Errorf("username already taken :P")
	}

	if err := user.CreateUser(db); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func GetUserService(db *mongodb.Database, username string) (models.User, int, error) {
	user := models.User{
		Username: username,
	}

	// check if user exist
	if err := checkIfUserExist(db, username); err != nil {
		return user, http.StatusBadRequest, err
	}

	if err := user.GetUserByUsername(db); err != nil {
		return user, http.StatusBadRequest, err
	}

	return user, http.StatusOK, nil
}

func UpdateUserService(db *mongodb.Database, req models.UpdateUserRequest, userID string) (int, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	user := models.User{
		ID: id,
	}

	if err := user.GetUserByID(db); err != nil {
		return http.StatusBadRequest, err
	}

	if req.Username != "" {
		// check if user with the same username already exist
		if err := checkIfUserExist(db, req.Username); err == nil {
			return http.StatusBadRequest, fmt.Errorf("username already taken :P")
		}

		// if not, go on and update
		user.Username = req.Username
	}

	if req.Age > 0 {
		user.Age = req.Age
	}

	if req.AboutMe != "" {
		user.AboutMe = req.AboutMe
	}

	if req.Gender != "" {
		user.Gender = req.Gender
	}

	if err := user.UpdateUserByID(db); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func DeleteUserService(db *mongodb.Database, userID string) (int, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	user := models.User{
		ID: id,
	}

	if err := user.GetUserByID(db); err != nil {
		return http.StatusBadRequest, err
	}

	if err := user.DeleteUserByID(db); err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
