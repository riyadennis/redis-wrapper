package redis_wrapper

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockStorage struct {
	mock.Mock
}

func (m MockStorage) Get(key string) (string, error) {
	return "value", nil
}
func (m MockStorage) Set(key, value string, lifeTime time.Duration)  error {
	return  nil
}
func (m MockStorage) Delete(key string)  error {
	return  nil
}

// this tests expects redis to be running
func TestClient_Create(t *testing.T) {
	c := Client{}
	cr, err := c.Create()
	assert.NoError(t, err)
	assert.NotEmpty(t, cr.RedisClient)
}
func TestClient_Get(t *testing.T) {
	mockClient := MockStorage{}
	re, err := mockClient.Get("string")
	assert.NoError(t, err)
	assert.Equal(t, re, "value")
}
func TestClient_Set(t *testing.T) {
	mockClient := MockStorage{}
	err := mockClient.Set("sample", "value", 0)
	assert.NoError(t, err)
}

func TestClient_Delete(t *testing.T) {
	mockClient := MockStorage{}
	err := mockClient.Delete("sample")
	assert.NoError(t, err)
}
