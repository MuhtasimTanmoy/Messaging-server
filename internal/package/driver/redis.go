package driver

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

// Redis driver
type Redis struct {
	Client   *redis.Client
	Addr     string
	Password string
	DB       int
}

// NewRedisDriver create a new instance
func NewRedisDriver() *Redis {
	return &Redis{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
}

// Connect establish a redis connection
func (r *Redis) Connect() (bool, error) {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})

	_, err := r.Ping()

	if err != nil {
		return false, err
	}

	return true, nil
}

// Ping checks the redis connection
func (r *Redis) Ping() (bool, error) {
	pong, err := r.Client.Ping().Result()

	if err != nil {
		return false, err
	}
	return pong == "PONG", nil
}

// Set sets a record
func (r *Redis) Set(key, value string, expiration time.Duration) (bool, error) {
	result := r.Client.Set(key, value, expiration)

	if result.Err() != nil {
		return false, result.Err()
	}

	return result.Val() == "OK", nil
}

// Get gets a record value
func (r *Redis) Get(key string) (string, error) {
	result := r.Client.Get(key)

	if result.Err() != nil {
		return "", result.Err()
	}

	return result.Val(), nil
}

// Exists deletes a record
func (r *Redis) Exists(key string) (bool, error) {
	result := r.Client.Exists(key)

	if result.Err() != nil {
		return false, result.Err()
	}

	return result.Val() > 0, nil
}

// Del deletes a record
func (r *Redis) Del(key string) (int64, error) {
	result := r.Client.Del(key)

	if result.Err() != nil {
		return 0, result.Err()
	}

	return result.Val(), nil
}