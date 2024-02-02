package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloUser(c *gin.Context) {
	fmt.Println("hello")
	c.JSON(http.StatusOK, "hello")
}
