package mock

type mockTokenCache struct {
	IsTokenValidCallCount int
	StoreTokenCallCount   int
}

func NewMockTokenCache() *mockTokenCache {
	return &mockTokenCache{}
}

func (t *mockTokenCache) IsTokenValid() (bool, error) {
	t.IsTokenValidCallCount++
	return true, nil
}

func (t *mockTokenCache) StoreToken(token string, userId int64) {
	t.StoreTokenCallCount++
}
