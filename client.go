package xredis

import (
	"github.com/redis/go-redis/v9"
)

// A Client defines all client Redis actions (Client or ClusterClient).
type Client = redis.Cmdable

var _ Client = (*redis.Client)(nil)
