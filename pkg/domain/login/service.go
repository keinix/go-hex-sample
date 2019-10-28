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
	AddNewUser(username string, password string)
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
	hash := hashPassword(password)
	if hash != user.PasswordHash {
		return "", errors.New("password is incorrect")
	}
	token = makeSessionToken()
	s.cache.StoreToken(token, user.Id)
	return token, nil
}

func (s *service) GetSessionTokenFromRefreshToken(refreshToken string) (token string, err error) {
	return "", nil
}

func (s *service) IsTokenValid(token string) (bool, error) {
	return s.cache.IsTokenValid()
}

func (s *service) AddNewUser(username string, password string) {
	user := User{
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	s.repo.AddUser(user)
}
