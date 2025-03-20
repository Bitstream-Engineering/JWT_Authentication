package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(c *gin.Context) {
	// GET the email & password off the request body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

	}
}
