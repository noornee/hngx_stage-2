package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/noornee/hngx_stage-2/internal/models"
	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkIfUserExist(db *mongodb.Database, name string) error {
	user := models.User{
		Name: name,
	}
	err := user.GetUserByName(db)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user does not exist")
		}
		return err
	}
	return nil
}

func CreateUserService(db *mongodb.Database, req models.CreateUserRequest) (*models.User, int, error) {
	user := models.User{
		Name: req.Name,
	}

	// check if user with the username already exist
	if err := checkIfUserExist(db, req.Name); err == nil {
		return &models.User{}, http.StatusBadRequest, fmt.Errorf("username already taken :P")
	}

	userData, err := user.CreateUser(db)
	if err != nil {
		return &models.User{}, http.StatusInternalServerError, err
	}

	return userData, http.StatusOK, nil
}

func GetUserService(db *mongodb.Database, userID string) (models.User, int, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.User{}, http.StatusBadRequest, err
	}

	user := models.User{
		ID: id,
	}
	if err := user.GetUserByID(db); err != nil {
		return user, http.StatusNotFound, err
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
		return http.StatusNotFound, err
	}

	if req.Name != "" {
		// check if user with the same username already exist
		// if err is nil, then user exist
		if err := checkIfUserExist(db, req.Name); err == nil {
			return http.StatusBadRequest, fmt.Errorf("username already taken :P")
		}

		// if not, go on and update
		user.Name = req.Name
		user.UpdatedAt = time.Now()
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
		return http.StatusNotFound, err
	}

	if err := user.DeleteUserByID(db); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
