package handler

import (
	"instagram/pkg/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	CreatePost(c *gin.Context)
	GetAllPosts(c *gin.Context)
	GetAllPostsOfUser(c *gin.Context)
}

func (h *handler) CreatePost(c *gin.Context) {
	var createPostRequest model.CreatePostRequest
	if err := c.ShouldBindJSON(&createPostRequest); err != nil {
		log.Println("error with create post request: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.GetString("username")
	if err := h.s.CreatePost(createPostRequest, username); err != nil {
		log.Println("error creating post: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *handler) GetAllPosts(c *gin.Context) {
	username := c.GetString("username")
	posts, err := h.s.GetAllPosts(username)
	if err != nil {
		log.Println("error getting all posts: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func (h *handler) GetAllPostsOfUser(c *gin.Context) {
	username := c.Param("username")
	posts, err := h.s.GetAllPosts(username)
	if err != nil {
		log.Println("error getting all posts: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}
