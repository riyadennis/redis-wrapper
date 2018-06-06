package redis_wrapper

import (
	"github.com/go-redis/redis"
	"time"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key, value string, lifeTime time.Duration) error
	Delete(key string) error
}
type Client struct {
	RedisClient *redis.Client
}

func (c *Client) Create() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	c.RedisClient = client
	return c, nil
}
func (c Client) Ping() (error) {
	_, err := c.RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func (c Client) Get(key string) (string, error) {
	value, err := c.RedisClient.Get(key).Result()
	if err != nil {
		return "", err
	}
	if err == redis.Nil {
		return "", err
	}
	return value, nil
}
func (c Client) Set(key, value string, lifeTime time.Duration) error {
	_, err := c.RedisClient.Set(key, value, lifeTime).Result()
	if err != nil {
		return err
	}
	if err == redis.Nil {
		return err
	}
	return nil
}
func (c Client) Delete(key string) error {
	_, err := c.RedisClient.Del(key).Result()
	if err != nil {
		return err
	}
	if err == redis.Nil {
		return err
	}
	return nil
}
