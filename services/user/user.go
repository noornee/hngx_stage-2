package user

import (
	"net/http"

	"github.com/noornee/hngx_stage-2/internal/models"
	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
)

func CreateUserService(db *mongodb.Database, req models.CreateUserRequest) (int, error) {
	user := models.User{
		Username: req.Username,
	}

	err := user.CreateUser(db)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
