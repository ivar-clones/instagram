package service

import (
	"instagram/pkg/constants"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(username, password string) (string, error)
	GetUserToken(username, password string) (string, error)
}

func (s *service) CreateUser(username, password string) (string, error) {
	exists, err := s.repo.IsUserPresent(username)
	if err != nil {
		log.Println("error checking username: ", err.Error())
		return "", err
	}

	if exists {
		return "", constants.ErrUsernameAlreadyTaken
	}

	token, err := createToken(username)
	if err != nil {
		log.Println("error creating token: ", err.Error())
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	if err := s.repo.CreateUser(username, string(hashedPassword)); err != nil {
		log.Println("error creating user: ", err.Error())
		return "", err
	}

	return token, nil
}

func (s *service) GetUserToken(username, password string) (string, error) {
	hashedPassword, err := s.repo.GetUserPassword(username)
	if err != nil {
		log.Println("error getting user: ", err.Error())
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		log.Println("error comparing password: ", err.Error())
		return "", constants.ErrIncorrectPassword
	}

	token, err := createToken(username)
	if err != nil {
		log.Println("error creating token: ", err.Error())
		return "", err
	}

	return token, nil
}

func createToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "instagram-app",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	if err != nil {
		log.Println("error signing claims: ", err.Error())
		return "", err
	}

	return tokenString, nil
}
