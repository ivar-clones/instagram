package service

import (
	"instagram/pkg/model"
	"log"
)

type PostService interface {
	CreatePost(postToCreate model.CreatePostRequest, username string) error
	GetAllPosts(username string) ([]model.PostResponse, error)
	DeletePost(username string, id int) error
	GetPost(id int) (model.PostResponse, error)
}

func (s *service) CreatePost(postToCreate model.CreatePostRequest, username string) error {
	dbPostToCreate := model.MapFromCreatePostRequestToPostDB(postToCreate, username)

	if err := s.repo.CreatePost(dbPostToCreate); err != nil {
		log.Println("error creating post: ", err.Error())
		return err
	}

	return nil
}

func (s *service) GetAllPosts(username string) ([]model.PostResponse, error) {
	dbPosts, err := s.repo.GetAllPosts(username)
	if err != nil {
		log.Println("error fetching all posts: ", err.Error())
		return nil, err
	}

	var posts []model.PostResponse

	for _, dbPost := range dbPosts {
		posts = append(posts, model.MapFromPostDBToPostResponse(dbPost))
	}

	return posts, nil
}

func (s *service) DeletePost(username string, id int) error {
	if err := s.repo.DeletePost(username, id); err != nil {
		log.Println("error deleting post: ", err.Error())
		return err
	}

	return nil
}

func (s *service) GetPost(id int) (model.PostResponse, error) {
	dbPost, err := s.repo.GetPost(id)
	if err != nil {
		log.Println("error fetching all posts: ", err.Error())
		return model.PostResponse{}, err
	}

	post := model.MapFromPostDBToPostResponse(dbPost)

	return post, nil
}
