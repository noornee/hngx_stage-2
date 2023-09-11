package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
)

func Setup(db *mongodb.Database, validator *validator.Validate) *gin.Engine {
	r := gin.New()

	// middleware
	r.Use(gin.Logger())

	// User route
	User(r, db, validator)

	return r
}
