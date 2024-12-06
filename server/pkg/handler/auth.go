package handler

import (
	"errors"
	"instagram/pkg/constants"
	"instagram/pkg/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Signup(c *gin.Context)
}

func (h *handler) Signup(c *gin.Context) {
	var authRequest model.AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		log.Println("error with signup request: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.s.CreateUser(authRequest.Username, authRequest.Password)
	if err != nil {
		log.Println("error creating user: ", err.Error())
		if errors.Is(err, constants.ErrUsernameAlreadyTaken) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "username": authRequest.Username})
}

func (h *handler) Login(c *gin.Context) {
	var authRequest model.AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		log.Println("error with signup request: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.s.GetUserToken(authRequest.Username, authRequest.Password)
	if err != nil {
		log.Println("error logging in user: ", err.Error())
		if errors.Is(err, constants.ErrIncorrectPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "username": authRequest.Username})
}
