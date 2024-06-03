package handlers

import (
	"context"
	"golang-server/module/core/storage"
	"golang-server/pkg/logger"
)

type SyncUsersHandler struct {
	asynqStorage storage.IAsynqStorage
}

func NewSyncUsersHandler(
	asynqStorage storage.IAsynqStorage,
) SyncUsersHandler {
	return SyncUsersHandler{
		asynqStorage: asynqStorage,
	}
}

func (a SyncUsersHandler) Handle(ctx context.Context) error {
	for idx := 0; idx < 10; idx++ {
		err := a.asynqStorage.AddToSyncUsersTask(ctx)
		if err != nil {
			logger.Error(ctx, err, "SyncUsersHandler AddToSyncUsersTask error")
		}
	}
	return nil
}
