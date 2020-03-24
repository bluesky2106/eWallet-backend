package redis

import (
	"fmt"

	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/go-redis/redis"
)

// Config : Redis configurations
type Config struct {
	Addr     string `json: "addr"`
	DB       int    `json: "db"`
	Password string `json: "password"`
}

// Client : redis client struct
type Client struct {
	conf   *Config
	client *redis.Client
}

// Init : connect to redis servers
func Init(conf *commonConfig.Config) (*Client, error) {
	redisConf := &Config{
		Addr:     fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		DB:       conf.Redis.DB,
		Password: conf.Redis.Password,
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})

	pong, err := redisClient.Ping().Result()
	if pong == "PONG" {
		fmt.Println("Redis connect successfully!")
	} else {
		return nil, errs.New(errs.ECRedisConnection, err.Error(), errs.EMRedisConnection)
	}

	return &Client{
		conf:   redisConf,
		client: redisClient,
	}, nil
}

// Set : [key, value]
func (rc *Client) Set(key string, value interface{}) (err error) {
	err = rc.client.Set(key, value, 0).Err()
	return
}

// Get : key
func (rc *Client) Get(key string) (value string, err error) {
	value, err = rc.client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("Key does not exist")
	}
	return
}

// Del : []keys
func (rc *Client) Del(keys []string) (err error) {
	err = rc.client.Del(keys...).Err()
	return
}

// SetList : key, []interface{}
func (rc *Client) SetList(key string, values []interface{}) (err error) {
	err = rc.client.LPush(key, values...).Err()
	return
}
