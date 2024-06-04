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
	"time"

	"github.com/hibiken/asynq"
)

type SyncUsersProcessor struct {
	constantStorage storage.IConstantStorage
	mUserStorage    storage.IDbmStorage
	userStorage     storage.IUserStorage
	asynqStorage    storage.IAsynqStorage
}

func NewSyncUsersProcessor(
	constantStorage storage.IConstantStorage,
	mUserStorage storage.IDbmStorage,
	userStorage storage.IUserStorage,
	asynqStorage storage.IAsynqStorage,
) *SyncUsersProcessor {
	return &SyncUsersProcessor{
		constantStorage: constantStorage,
		mUserStorage:    mUserStorage,
		userStorage:     userStorage,
		asynqStorage:    asynqStorage,
	}
}

func (processor *SyncUsersProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p task.SyncUsersData
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		logger.Error(ctx, err, fmt.Sprintf("json.Unmarshal failed: %v: %v", err, asynq.SkipRetry))
		return asynq.SkipRetry
	}
	logger.Info(ctx, "syncUsersProcessor ProcessTask")
	duration := time.Duration(2) * time.Second
	<-time.After(duration)
	for T := 0; T < 5; T++ {
		userNum, err := processor.constantStorage.FindOne(ctx, constants.UsersNum)
		if err != nil {
			return err
		}
		logger.Info(ctx, fmt.Sprintf("handle batch with offset %s", userNum))

		num, err := strconv.Atoi(userNum.Value)
		if err != nil {
			return err
		}
		mUsers, err := processor.mUserStorage.ListUsers(ctx, num, constants.MBatchSize)
		if err != nil {
			continue
		}
		var dataInsert []*model.UserModel
		for _, item := range mUsers {
			dataDB := model.UserModel{
				ID:            item.UserID,
				LoyaltyID:     item.LoyaltyID,
				Phone:         item.Phone,
				CurOriginalID: item.CurOriginalID,
			}
			if item.Email != "" {
				dataDB.Email = &item.Email
			}
			dataInsert = append(dataInsert, &dataDB)
		}
		err = processor.userStorage.TxInsertMUsers(ctx, num, dataInsert)
		if err != nil {
			logger.Error(ctx, err, fmt.Sprintf("Stopped by updating %d constant error", num+constants.MBatchSize))
			break
		}
		processor.asynqStorage.AddToRegisterUserTask(ctx, dataInsert)
	}
	return nil
}
