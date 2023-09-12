package router

import (
	"net/http"

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

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found.",
			"status":  http.StatusNotFound,
		})
	})

	return r
}
