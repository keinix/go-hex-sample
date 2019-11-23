package login_test

import (
	"go-hex-sample/pkg/domain/login"
	"go-hex-sample/pkg/mock"
	"testing"
)

func TestService_GetSessionToken(t *testing.T) {
	repo := mock.NewMockLoginRepository()
	cache := mock.NewMockTokenCache()
	service := login.NewService(repo, cache)

	// valid user
	err := service.AddNewUser("link", "zelda123")
	if err != nil {
		t.Errorf("error in test: %v", err)
	}
	_, err = service.GetSessionToken("link", "zelda123")
	if err != nil {
		t.Errorf("error in test: %v", err)
	}
	if cache.StoreTokenCallCount != 1 {
		t.Errorf("expected stored token count: 1; actual: %d", cache.StoreTokenCallCount)
		t.Fail()
	}
	// invalid user
	cache.StoreTokenCallCount = 0
	token, _ := service.GetSessionToken("bad_bad", "not_good")
	if token != "" {
		t.Error("token returned for invalid user")
		t.Fail()
	}
	if cache.StoreTokenCallCount > 0 {
		t.Error("token stored for invalid user")
		t.Fail()
	}
}
