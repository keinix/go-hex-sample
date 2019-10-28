package psql

import (
	"fmt"
	"go-hex-sample/pkg/domain/login"
)

type userRepository struct {
}

func NewUserRepository() login.Repository {
	return &userRepository{}
}

func (r *userRepository) GetUser(username string) (*login.User, error) {
	db, err := openDb()
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}
	var result login.User
	if db.Where("username = ?", username).First(&result).RecordNotFound() {
		return nil, fmt.Errorf("user %v does not exist", username)
	}
	if err := db.Close(); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) AddUser(user login.User) {

}
