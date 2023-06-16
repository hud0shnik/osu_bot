package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/hud0shnik/osu_bot/mods"

	"github.com/spf13/viper"
)

// Структуры для работы с Telegram API

type telegramResponse struct {
	Result []update `json:"result"`
}

type update struct {
	UpdateId int     `json:"update_id"`
	Message  message `json:"message"`
}

type message struct {
	Chat    chat    `json:"chat"`
	Text    string  `json:"text"`
	Sticker sticker `json:"sticker"`
}

type chat struct {
	ChatId int `json:"id"`
}

type sticker struct {
	File_id string `json:"file_id"`
}

// Функция получения апдейтов
func getUpdates(botUrl string, offset int) ([]update, error) {

	// Rest запрос для получения апдейтов
	resp, err := http.Get(botUrl + "/getUpdates?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Запись и обработка полученных данных
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse telegramResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Result, nil
}

// Функция генерации и отправки ответа
func respond(botUrl string, update update) {

	// Проверка на сообщение
	if update.Message.Text == "" {
		mods.SendMsg(botUrl, update.Message.Chat.ChatId, "Пока я воспринимаю только текст")
		return
	}

	// Разделение текста пользователя на слайс
	request := append(strings.Split(update.Message.Text, " "), "", "")

	// Обработчик команд
	switch request[0] {
	case "/info":
		mods.SendUserInfo(botUrl, update.Message.Chat.ChatId, request[1])
	case "/recent":
		mods.SendRecentBeatmap(botUrl, update.Message.Chat.ChatId, request[1])
	case "/online":
		mods.SendOnlineInfo(botUrl, update.Message.Chat.ChatId, request[1])
	case "/map":
		mods.SendMapInfo(botUrl, update.Message.Chat.ChatId, request[1], request[2])
	case "/start", "/help":
		mods.Help(botUrl, update.Message.Chat.ChatId)
	}

}

func main() {

	// Инициализация конфига (токенов)
	err := initConfig()
	if err != nil {
		log.Fatalf("initConfig error: %s", err)
		return
	}

	// Url бота для отправки и приёма сообщений
	botUrl := "https://api.telegram.org/bot" + viper.GetString("token")
	offSet := 0

	// Цикл работы бота
	for {

		// Получение апдейтов
		updates, err := getUpdates(botUrl, offSet)
		if err != nil {
			log.Fatalf("getUpdates error: %s", err)
		}

		// Обработка апдейтов
		for _, update := range updates {
			respond(botUrl, update)
			offSet = update.UpdateId + 1
		}

		// Вывод в консоль для тестов
		// fmt.Println(updates)
	}
}

// Функция инициализации конфига (всех токенов)
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
