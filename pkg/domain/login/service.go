package login

import "errors"

type Repository interface {
	GetUser(username string) (*User, error)
	AddUser(user User)
}

type TokenCache interface {
	IsTokenValid() (bool, error)
	StoreToken(token string, userId int64)
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
	ok, err := checkPlainTextMatchesHash(password, user.PasswordHash)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("password is incorrect")
	}
	token = newSessionToken()
	s.cache.StoreToken(token, user.Id)
	return token, nil
}

func (s *service) GetSessionTokenFromRefreshToken(refreshToken string) (token string, err error) {
	return "", nil
}

func (s *service) IsTokenValid(token string) (bool, error) {
	return s.cache.IsTokenValid()
}

func (s *service) AddNewUser(username string, password string) error {
	hash, err := newPasswordHash(password)
	if err != nil {
		return err
	}
	user := User{
		Username:     username,
		PasswordHash: hash,
	}
	s.repo.AddUser(user)
	return nil
}
