package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	keyWallet = "wallet:%s"
)

var (
	errCacheNotFound = errors.New("cache not foud")
)

type (
	module struct {
		client *redis.Client
	}

	Opts struct {
	}

	Cache interface {
		SetWallet(ctx context.Context, userID string, value interface{}) error
		GetWallet(ctx context.Context, userID string, v interface{}) error
	}
)

func New(client *redis.Client) Cache {
	return &module{
		client: client,
	}
}

func (c *module) SetWallet(ctx context.Context, userID string, value interface{}) error {
	return c.client.Set(ctx, fmt.Sprintf(keyWallet, userID), value, time.Second*time.Duration(5)).Err()
}

func (c *module) GetWallet(ctx context.Context, userID string, v interface{}) error {
	cmd := c.client.Get(ctx, fmt.Sprintf(keyWallet, userID))
	if cmd.Err() != nil && errors.Is(cmd.Err(), redis.Nil) {
		return errCacheNotFound
	}

	return json.Unmarshal([]byte(cmd.String()), &v)
}
