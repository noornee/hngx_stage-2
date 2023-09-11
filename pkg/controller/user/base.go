package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
)

type Controller struct {
	DB        *mongodb.Database
	Validator *validator.Validate
}
