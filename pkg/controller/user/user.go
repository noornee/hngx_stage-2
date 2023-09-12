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
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request", "error": err.Error(), "status": http.StatusBadRequest})
		return
	}

	if err := base.Validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation failed", "error": utility.ValidationResponse(err, base.Validator), "status": http.StatusBadRequest})
		return
	}

	code, err := services.CreateUserService(base.DB, req)
	if err != nil {
		c.JSON(code, gin.H{"message": "failed to parse request", "error": err.Error(), "status": code})
		return
	}

	c.JSON(code, gin.H{"message": "user creation successful", "status": code})
}

func (base *Controller) GetUser(c *gin.Context) {
	username := c.Params.ByName("username")

	result, code, err := services.GetUserService(base.DB, username)
	if err != nil {
		c.JSON(code, gin.H{"message": "failed to parse request", "error": err.Error(), "status": code})
		return
	}

	c.JSON(code, gin.H{"message": "success", "data": result, "status": code})
}

func (base *Controller) UpdateUser(c *gin.Context) {
	req := models.UpdateUserRequest{}

	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request", "error": err.Error(), "status": http.StatusBadRequest})
		return
	}

	code, err := services.UpdateUserService(base.DB, req, id)
	if err != nil {
		c.JSON(code, gin.H{"message": "failed to parse request", "error": err.Error(), "status": code})
		return
	}

	c.JSON(code, gin.H{"message": "user update successful", "status": code})
}

func (base *Controller) DeleteUser(c *gin.Context) {
	username := c.Params.ByName("id")

	code, err := services.DeleteUserService(base.DB, username)
	if err != nil {
		c.JSON(code, gin.H{"message": "failed to parse request", "error": err.Error(), "status": code})
		return
	}

	c.JSON(code, gin.H{"message": "user delete successful", "status": code})
}
