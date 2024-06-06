package business

import (
	"context"
	"golang-server/pkg/logger"
	"io"
)

func (b biz) HandleEventStreamConnection(ctx context.Context, w io.Writer) (bool, error) {
	logger.Info(ctx, "biz HandleEventStreamConnection")
	return false, nil
}
