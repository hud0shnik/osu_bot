package mods

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Структуры для отправки сообщений, стикеров и картинок
type SendMessage struct {
	ChatId    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type SendSticker struct {
	ChatId     int    `json:"chat_id"`
	StickerUrl string `json:"sticker"`
}

type SendPhoto struct {
	ChatId    int    `json:"chat_id"`
	PhotoUrl  string `json:"photo"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}

// Функция отправки сообщения
func SendMsg(botUrl string, update Update, msg string) error {

	// Формирование сообщения
	botMessage := SendMessage{
		ChatId:    update.Message.Chat.ChatId,
		Text:      msg,
		ParseMode: "HTML",
	}
	buf, err := json.Marshal(botMessage)
	if err != nil {
		log.Printf("json.Marshal error: %s", err)
		return err
	}

	// Отправка сообщения
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Printf("sendMessage error: %s", err)
		return err
	}
	return nil
}

// Функция отправки стикера
func SendStck(botUrl string, update Update, url string) error {

	// Формирование стикера
	botStickerMessage := SendSticker{
		ChatId:     update.Message.Chat.ChatId,
		StickerUrl: url,
	}
	buf, err := json.Marshal(botStickerMessage)
	if err != nil {
		log.Printf("json.Marshal error: %s", err)
		return err
	}
	// Отправка стикера
	_, err = http.Post(botUrl+"/sendSticker", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Printf("sendSticker error: %s", err)
		return err
	}
	return nil
}

// Функция отправки картинки
func SendPict(botUrl string, update Update, pic SendPhoto) error {

	// Указание парсмода текста под картинкой и айди чата
	pic.ParseMode = "HTML"
	pic.ChatId = update.Message.Chat.ChatId

	// Формирование картинки
	buf, err := json.Marshal(pic)
	if err != nil {
		log.Printf("json.Marshal error: %s", err)
		return err
	}
	// Отправка картинки
	_, err = http.Post(botUrl+"/sendPhoto", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Printf("sendPhoto error: %s", err)
		return err
	}
	return nil
}
