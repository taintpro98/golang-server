package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-server/config"
	"golang-server/pkg/logger"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedisClient interface {
	Get(ctx context.Context, key string, outputType interface{}) error
	Set(ctx context.Context, key string, value interface{}, ttl int64) error // interface is not a pointer

	Publish(ctx context.Context, channel string, value interface{}) error // value is not a pointer
	Subscribe(ctx context.Context, channel string) *redis.PubSub
}

type redisClient struct {
	client *redis.Client
	cfg    config.RedisConfig
}

func NewRedisClient(ctx context.Context, cfg config.RedisConfig) (IRedisClient, error) {
	redisDb, err := strconv.Atoi(cfg.DB)
	if err != nil {
		logger.Error(ctx, err, "NewRedisClient err: redisDb not number %s")
		return nil, err
	}

	redisOpt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		Username: cfg.Username,
		DB:       redisDb,
	}

	rdb := redis.NewClient(redisOpt)
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return &redisClient{
		client: rdb,
		cfg:    cfg,
	}, nil
}

func (c *redisClient) generateKey(key string) string {
	return fmt.Sprintf("%s:%s", c.cfg.Prefix, key)
}

func (c *redisClient) Get(ctx context.Context, key string, outputType interface{}) error {
	val, err := c.client.Get(ctx, c.generateKey(key)).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), &outputType)
}

func (c *redisClient) Set(ctx context.Context, key string, value interface{}, seconds int64) error {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	ttl := 0 * time.Second
	if seconds != 0 {
		ttl = time.Duration(seconds) * time.Second
	}
	return c.client.Set(ctx, c.generateKey(key), jsonBytes, ttl).Err()
}

func (c *redisClient) Publish(ctx context.Context, channel string, value interface{}) error {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Publish(ctx, channel, jsonBytes).Err()
}

func (c *redisClient) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return c.client.Subscribe(ctx, channel)
}
