package worker

import (
	"context"
	"golang-server/app/worker/processor"
	"golang-server/app/worker/task"
	"golang-server/config"
	"golang-server/module/core/storage"
	"golang-server/pkg/cache"
	"golang-server/pkg/telegram"

	"github.com/elastic/go-elasticsearch/v7"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

func NewWorkerDispatcher(
	ctx context.Context,
	cnf config.Config,
	redisClient cache.IRedisClient,
	redisPubsub cache.IRedisClient,
	es *elasticsearch.Client,
	db *gorm.DB,
	mDb *gorm.DB,
	mux *asynq.ServeMux,
	redisQueue *asynq.Client,
	telegramBot telegram.ITelegramBot,
) {
	constantStorage := storage.NewConstantStorage(cnf.Database, db)
	userStorage := storage.NewUserStorage(cnf.Database, db)
	mUserStorage := storage.NewDbmStorage(cnf.DBM, mDb)
	notificationStorage := storage.NewNotificationStorage(telegramBot)
	elasticStorage := storage.NewElasticStorage(es)
	asynqStorage := storage.NewAsynqStorage(cnf.RedisQueue, redisQueue)

	mux.Handle(task.CreatePostQueueName(cnf.RedisQueue.Prefix), processor.NewCreatePostProcessor(
		redisPubsub,
		userStorage,
	))
	mux.Handle(task.RegisterUserQueueName(cnf.RedisQueue.Prefix), processor.NewRegisterUserProcessor(
		notificationStorage,
		elasticStorage,
	))
	mux.Handle(task.SyncUsersQueueName(cnf.RedisQueue.Prefix), processor.NewSyncUsersProcessor(
		constantStorage,
		mUserStorage,
		userStorage,
		asynqStorage,
	))
}
