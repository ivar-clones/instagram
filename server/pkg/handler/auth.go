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
			c.JSON(http.StatusBadRequest, gin.H{"error": "username already taken"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/")
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
			c.Status(http.StatusUnauthorized)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/")
}
