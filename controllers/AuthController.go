package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Lets(c *gin.Context) {
	c.JSON(http.StatusOK, "lets go")
}
