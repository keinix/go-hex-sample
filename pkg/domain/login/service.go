package login

import (
	"errors"
	"fmt"
	"go-hex-sample/pkg/domain/crypto"
)

type Repository interface {
	GetUser(username string) (*User, error)
	AddUser(user User) error
}

type TokenCache interface {
	IsTokenValid(token string) (bool, error)
	StoreToken(token string, userId int64) error
}

type Service interface {
	GetSessionToken(username string, password string) (token string, err error)
	GetSessionTokenFromRefreshToken(refreshToken string) (token string, err error)
	IsTokenValid(token string) (bool, error)
	AddNewUser(username string, password string) error
}

type service struct {
	repo  Repository
	cache TokenCache
}

func NewService(repo Repository, cache TokenCache) Service {
	return &service{
		repo:  repo,
		cache: cache,
	}
}

func (s *service) GetSessionToken(username string, password string) (token string, err error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", err
	}
	ok, err := crypto.CheckPlainTextMatchesHash(password, user.PasswordHash)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("password is incorrect")
	}
	token, err = crypto.NewSessionToken()
	if err != nil {
		return "", fmt.Errorf("error creating new token: %w", err)
	}
	err = s.cache.StoreToken(token, user.Id)
	if err != nil {
		return "", fmt.Errorf("error storing token %w", err)
	}
	return token, nil
}

func (s *service) GetSessionTokenFromRefreshToken(refreshToken string) (token string, err error) {
	return "", nil
}

func (s *service) IsTokenValid(token string) (bool, error) {
	return s.cache.IsTokenValid(token)
}

func (s *service) AddNewUser(username string, password string) error {
	hash, err := crypto.NewPasswordHash(password)
	if err != nil {
		return err
	}
	user := User{
		Username:     username,
		PasswordHash: hash,
	}
	err = s.repo.AddUser(user)
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}
	return nil
}
