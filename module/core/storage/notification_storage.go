package storage

import (
	"context"
	"fmt"
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
	var content string
	for _, item := range param.Users {
		content += fmt.Sprintf("%s\n", item.ID)
	}
	return n.telegramBot.SendMessage(ctx, content)
}
