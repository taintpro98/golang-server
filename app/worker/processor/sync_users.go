package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-server/app/worker/task"
	"golang-server/module/core/model"
	"golang-server/module/core/storage"
	"golang-server/pkg/constants"
	"golang-server/pkg/logger"
	"strconv"

	"github.com/hibiken/asynq"
)

type syncUsersProcessor struct {
	constantStorage storage.IConstantStorage
	mUserStorage    storage.IDbmStorage
	userStorage     storage.IUserStorage
}

func NewSyncUsersProcessor(
	constantStorage storage.IConstantStorage,
	mUserStorage storage.IDbmStorage,
	userStorage storage.IUserStorage,
) *syncUsersProcessor {
	return &syncUsersProcessor{
		constantStorage: constantStorage,
		mUserStorage:    mUserStorage,
		userStorage:     userStorage,
	}
}

func (processor *syncUsersProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p task.SyncUsersData
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		logger.Error(ctx, err, fmt.Sprintf("json.Unmarshal failed: %v: %v", err, asynq.SkipRetry))
		return asynq.SkipRetry
	}
	logger.Info(ctx, "syncUsersProcessor ProcessTask")
	// duration := time.Duration(5) * time.Second
	// <-time.After(duration)
	for {
		userNum, err := processor.constantStorage.FindOne(ctx, constants.UsersNum)
		if err != nil {
			return err
		}
		logger.Info(ctx, fmt.Sprintf("handle batch with offset %s", userNum))

		num, err := strconv.ParseInt(userNum.Value, 10, 64)
		if err != nil {
			return err
		}
		mUsers, err := processor.mUserStorage.ListUsers(ctx, num, 50)
		if err != nil {
			continue
		}
		var dataInsert []model.UserModel
		for _, item := range mUsers {
			dataInsert = append(dataInsert, model.UserModel{
				ID:            item.UserID,
				LoyaltyID:     item.LoyaltyID,
				Phone:         item.Phone,
				Email:         &item.Email,
				CurOriginalID: item.CurOriginalID,
			})
		}
		err = processor.userStorage.InsertBatch(ctx, dataInsert)
		if err != nil {
			continue
		}
		err = processor.constantStorage.UpdateMany(ctx, constants.UsersNum, fmt.Sprintf("%d", num+50))
		if err != nil {
			break
		}
	}
	return nil
}
