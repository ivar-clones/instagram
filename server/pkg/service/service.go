package service

import "instagram/pkg/database"

type Service interface {
	UserService
	PostService
}

type service struct {
	repo database.Database
}

func New(repo database.Database) Service {
	return &service{
		repo: repo,
	}
}
