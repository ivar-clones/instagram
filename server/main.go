package main

import (
	"context"
	"instagram/pkg/database"
	"instagram/pkg/handler"
	"instagram/pkg/middleware"
	"instagram/pkg/service"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	r := gin.Default()
	allowedOriginsFromEnv := os.Getenv("ALLOWED_ORIGINS")
	allowedOrigins := strings.Split(allowedOriginsFromEnv, ",")
	r.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		panic("error connecting to database: " + err.Error())
	}
	defer conn.Close()

	repository := database.New(conn)
	service := service.New(repository)
	handler := handler.New(service)

	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)

	rg := r.Group("/api/v1")
	rg.Use(middleware.Auth())
	rg.POST("/posts", handler.CreatePost)
	rg.GET("/posts", handler.GetAllPosts)
	rg.GET("/posts/username/:username", handler.GetAllPostsOfUser)
	rg.DELETE("/posts/:id", handler.DeletePost)
	rg.GET("/posts/:id", handler.GetPost)

	if err := r.Run(":8080"); err != nil {
		panic("error creating server: " + err.Error())
	}
}
