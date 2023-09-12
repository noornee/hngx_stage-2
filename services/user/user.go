package user

import (
	"net/http"
	"time"

	"github.com/noornee/hngx_stage-2/internal/models"
	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserService(db *mongodb.Database, req models.CreateUserRequest) (*models.User, int, error) {
	user := models.User{
		Name: req.Name,
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

	if req.Name != "" {
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
		return http.StatusBadRequest, err
	}

	if err := user.DeleteUserByID(db); err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
