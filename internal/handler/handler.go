package handler

import (
	"strings"

	"github.com/hud0shnik/osu_bot/internal/api"
	"github.com/hud0shnik/osu_bot/internal/commands"
	"github.com/hud0shnik/osu_bot/internal/send"
	"github.com/hud0shnik/osu_bot/internal/telegram"
)

// Функция генерации и отправки ответа
func Respond(botUrl string, update telegram.Update) {

	// Проверка на сообщение
	if update.Message.Text == "" {
		send.SendMsg(botUrl, update.Message.Chat.ChatId, "Пока я воспринимаю только текст")
		return
	}

	// Разделение текста пользователя на слайс
	request := append(strings.Split(update.Message.Text, " "), "", "")

	// Обработчик команд
	switch request[0] {
	case "/info":
		api.SendUserInfo(botUrl, update.Message.Chat.ChatId, request[1])
	case "/recent":
		api.SendRecentBeatmap(botUrl, update.Message.Chat.ChatId, request[1])
	case "/online":
		api.SendOnlineInfo(botUrl, update.Message.Chat.ChatId, request[1])
	case "/map":
		api.SendMapInfo(botUrl, update.Message.Chat.ChatId, request[1])
	case "/start", "/help":
		commands.Help(botUrl, update.Message.Chat.ChatId)
	}

}
