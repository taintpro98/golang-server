package storage

import (
	"context"
	"golang-server/module/core/dto"
	"golang-server/pkg/telegram"
)

type INotificationStorage interface {
	SendTelegramNotification(ctx context.Context, param dto.UserCreatedNotification) error
}

type notificationStorage struct {
	telegramBot telegram.ITelegramBot
}

func NewNotificationStorage(
	telegramBot telegram.ITelegramBot,
) INotificationStorage {
	return notificationStorage{
		telegramBot: telegramBot,
	}
}

// SendTelegramNotification implements INotificationStorage.
func (n notificationStorage) SendTelegramNotification(ctx context.Context, param dto.UserCreatedNotification) error {
	content := param.UserID
	return n.telegramBot.SendMessage(ctx, content)
}
