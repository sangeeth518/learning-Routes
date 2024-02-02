package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/learning-routes/database"
	"github.com/sangeeth518/learning-routes/models"
)

var Body struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HelloUser(c *gin.Context) {

	if c.BindJSON(&Body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to read body"})
		return

	}
	// database.ConnectTODb()
	newuser := models.User{Name: Body.Name, Email: Body.Email}
	result := database.DB.Create(&newuser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create body"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})

}
