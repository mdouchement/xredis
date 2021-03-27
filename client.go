package xredis

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
)

// A Client wraps default Redis client and adds more features.
type Client struct {
	*redis.Client
}

// New returns a new Client.
func New(r *redis.Client) *Client {
	return &Client{
		Client: r,
	}
}

// Runs excutes the given script.
func (c *Client) Run(ctx context.Context, script Script, keys []string, args ...interface{}) (Value, error) {
	r := c.Client.EvalSha(ctx, script.Hash(), keys, args...)
	if err := r.Err(); err != nil && strings.HasPrefix(err.Error(), "NOSCRIPT ") {
		c.Client.ScriptLoad(ctx, script.Source()) // ignoring error, relying on uncached EVAL

		r = c.Client.Eval(ctx, script.Source(), keys, args...) // Fallback
	}

	v, err := r.Result()
	return &value{value: v}, err
}
