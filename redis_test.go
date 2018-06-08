package redis_wrapper

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
//tests expects redis to be running
func TestClient_Create(t *testing.T) {
	c := Client{}
	cr, err := c.Create()
	assert.NoError(t, err)
	assert.NotEmpty(t, cr.RedisClient)
}
func TestClient_Get(t *testing.T) {
	c := Client{}
	cr, err := c.Create()
	assert.NoError(t, err)
	cr.Set("sample", "value", 0)
	re, err := cr.Get("sample")
	assert.NoError(t, err)
	assert.Equal(t, re, "value")
}
func TestClient_Set(t *testing.T) {
	c := Client{}
	cr, err := c.Create()
	assert.NoError(t, err)
	err = cr.Set("sample", "value", 0)
	assert.NoError(t, err)
}

func TestClient_Delete(t *testing.T) {
	c := Client{}
	cr, err := c.Create()
	assert.NoError(t, err)
	err = cr.Set("sample", "value", 0)
	assert.NoError(t, err)
	err = cr.Delete("sample")
	assert.NoError(t,err)
	val, _ := cr.Get("sample")
	assert.Empty(t, val)
}
