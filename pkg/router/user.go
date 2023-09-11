package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/noornee/hngx_stage-2/pkg/controller/user"
	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
)

func User(r *gin.Engine, db *mongodb.Database, validate *validator.Validate) *gin.Engine {
	user := user.Controller{DB: db, Validator: validate}

	v1 := r.Group("api/v1")
	{
		v1.POST("/create", user.CreateUser)
		v1.GET("/get/:username", user.GetUser)
		v1.PUT("/update/:id", user.UpdateUser)
		v1.DELETE("/delete/:id", user.DeleteUser)
	}

	return r
}
