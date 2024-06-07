package business

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang-server/pkg/constants"
	"golang-server/pkg/logger"
)

func (b biz) HandleEventStreamConnection(ctx context.Context, userID string) (*redis.PubSub, error) {
	logger.Info(ctx, "biz HandleEventStreamConnection", logger.LogField{
		Key:   "userID",
		Value: userID,
	})
	key := fmt.Sprintf("%s:%s", constants.PostsChannel, userID) // each user has a posts channel consuming all post
	return b.redisPubsub.Subscribe(ctx, key)
}
