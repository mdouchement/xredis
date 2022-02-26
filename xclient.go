package xredis

import (
	"context"
	"strings"
)

type (
	// A XClient wraps default Redis client and adds more features.
	XClient interface {
		Client
		Run(ctx context.Context, script Script, keys []string, args ...interface{}) (Value, error)
		RunOnce(ctx context.Context, script Script, keys []string, args ...interface{}) (Value, error)
	}

	client struct {
		Client
	}
)

// New returns a new Client.
func New(r Client) XClient {
	return &client{
		Client: r,
	}
}

// Runs excutes the given static script.
// It loads the script in the Redis script cache.
func (c *client) Run(ctx context.Context, script Script, keys []string, args ...interface{}) (Value, error) {
	r := c.Client.EvalSha(ctx, script.Hash(), keys, args...)
	if err := r.Err(); err != nil && strings.HasPrefix(err.Error(), "NOSCRIPT ") {
		c.Client.ScriptLoad(ctx, script.Source()) // ignoring error, relying on uncached EVAL

		r = c.Client.Eval(ctx, script.Source(), keys, args...) // Fallback
	}

	v, err := r.Result()
	return NewValue(v), err
}

// Runs excutes the given dynamic script.
// It does not load the script in the Redis script cache.
func (c *client) RunOnce(ctx context.Context, script Script, keys []string, args ...interface{}) (Value, error) {
	v, err := c.Client.Eval(ctx, script.Source(), keys, args...).Result()
	return NewValue(v), err
}
