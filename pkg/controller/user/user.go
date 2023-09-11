package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noornee/hngx_stage-2/internal/models"
	services "github.com/noornee/hngx_stage-2/services/user"
	"github.com/noornee/hngx_stage-2/utility"
)

func (base *Controller) CreateUser(c *gin.Context) {
	req := models.CreateUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request", "error": err.Error()})
		return
	}

	if err := base.Validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed", "error": utility.ValidationResponse(err, base.Validator)})
		return
	}

	code, err := services.CreateUserService(base.DB, req)
	if err != nil {
		c.JSON(code, gin.H{"message": "failed to parse request", "error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"message": "user creation successful"})
}
