package bot

import (
	"fmt"
	"log/slog"
	"os"
	"unicode/utf8"

	"girls/internal/model"
	"girls/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	writingAuthor   = 0
	writtingMessage = 1
	NoAction        = 2
)

type Bot struct {
	token   string
	service *service.Service
	Users   map[int64]*UserSettings `json:"current_user"`
}
type UserSettings struct {
	Username string
	Nickname string
	Text     string
	Stage    int
}

func New(token string, service *service.Service) *Bot {
	return &Bot{service: service, token: token, Users: make(map[int64]*UserSettings)}
}

func (b *Bot) Start() {
	bot, err := tgbotapi.NewBotAPI(b.token)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatid := update.Message.Chat.ID
		if b.Users[chatid] == nil {
			b.Users[chatid] = &UserSettings{Username: update.Message.From.UserName}
		}
		msg := tgbotapi.NewMessage(chatid, "")
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "write":
				msg.Text = "Напиши своё имя/кличку/прозвище(чтоб понятно было от кого пожелание)"
				b.Users[chatid].Stage = writingAuthor
				bot.Send(msg)
				continue
			case "send":
				if b.Users[chatid].Nickname != "" && b.Users[chatid].Text != "" {
					_, err := b.service.WriteCongratulations(model.Congratulations{
						Text:         b.Users[chatid].Text,
						NickName:     b.Users[chatid].Nickname,
						TelegramName: b.Users[chatid].Username,
					})
					if err != nil {
						slog.Error(err.Error())
						msg.Text = "Ошибка при отправке пожелания"
						bot.Send(msg)
						continue
					}
					msg.Text = "Пожелание отправлено. Если хочешь поменять текст - просто создай новое пожелание /write"
					bot.Send(msg)
					continue
				}

			default:
				continue
			}
		}
		if !update.Message.IsCommand() {
			switch b.Users[chatid].Stage {
			case writingAuthor:
				text := update.Message.Text
				b.Users[chatid].Nickname = text
				msg.Text = fmt.Sprintf("Ваше имя: %s\nВведите ваше пожелание(до 2000 символов)", text)
				bot.Send(msg)
				b.Users[chatid].Stage = writtingMessage
				continue
			case writtingMessage:
				text := update.Message.Text
				b.Users[chatid].Text = text
				if utf8.RuneCount([]byte(text)) > 2000 {
					msg.Text = "Слишком длинное сообщение. Укороти его"
					bot.Send(msg)
					continue
				}
				msg.Text = fmt.Sprintf("Ваше пожелание:\n%s\nВведите /send чтобы отправить ваше пожелание", text)

				bot.Send(msg)
				b.Users[chatid].Stage = NoAction
				continue

			}
		}

	}

}
