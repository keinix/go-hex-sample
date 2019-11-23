package mock

import (
	"errors"
	"go-hex-sample/pkg/domain/login"
)

type mockLoginRepository struct {
	GetUseCallCount  int
	AddUserCallCount int
	users            []login.User
}

func NewMockLoginRepository() *mockLoginRepository {
	return &mockLoginRepository{}
}

func (m *mockLoginRepository) GetUser(username string) (*login.User, error) {
	m.GetUseCallCount++
	var user login.User
	for _, u := range m.users {
		if u.Username == username {
			user = u
		}
	}
	if user == (login.User{}) {
		return nil, errors.New("user not in mock LoginRepository")
	}
	return &user, nil
}

func (m *mockLoginRepository) AddUser(user login.User) {
	m.AddUserCallCount++
	m.users = append(m.users, user)
}
