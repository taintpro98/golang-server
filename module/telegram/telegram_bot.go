package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang-server/config"
	"golang-server/pkg/logger"
)

type ITelegramBot interface {
	SendMessage(ctx context.Context, chatID int64, text string) error
}

type telegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(cnf config.TelegramBot) (ITelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(cnf.Token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	return telegramBot{
		bot: bot,
	}, nil
}

func (b telegramBot) SendMessage(ctx context.Context, chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.bot.Send(msg)
	if err != nil {
		logger.Error(ctx, err, "send message error")
		return err
	}
	return nil
}
