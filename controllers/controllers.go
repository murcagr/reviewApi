package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Status blablabla
func Status(c *gin.Context) {

	c.JSON(http.StatusOK, "Hello")

}

// UnrealizedMethod is dummy method
func UnrealizedMethod(c *gin.Context) {

	c.JSON(http.StatusOK, "This method wasn't realized.")

}
