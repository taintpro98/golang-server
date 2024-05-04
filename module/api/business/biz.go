package business

import (
	"context"
	"golang-server/module/telegram"
)

type IBiz interface {
	GetSports(ctx context.Context) error
}

type biz struct {
	telegramBot telegram.ITelegramBot
}

func NewBiz(
	telegramBot telegram.ITelegramBot,
) IBiz {
	return biz{
		telegramBot: telegramBot,
	}
}
