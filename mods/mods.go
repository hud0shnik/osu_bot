package mods

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// Структуры для работы с Telegram API

type TelegramResponse struct {
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat    Chat    `json:"chat"`
	Text    string  `json:"text"`
	Sticker Sticker `json:"sticker"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type Sticker struct {
	File_id string `json:"file_id"`
}

// Структуры для работы с другими API

type OsuUserInfo struct {
	Error          string   `json:"error"`
	Username       string   `json:"username"`
	Names          []string `json:"previous_usernames"`
	AvatarUrl      string   `json:"avatar_url"`
	CountryCode    string   `json:"country_code"`
	GlobalRank     string   `json:"global_rank"`
	CountryRank    string   `json:"country_rank"`
	PP             string   `json:"pp"`
	PlayTime       string   `json:"play_time"`
	SSH            string   `json:"ssh"`
	SS             string   `json:"ss"`
	SH             string   `json:"sh"`
	S              string   `json:"s"`
	A              string   `json:"a"`
	RankedScore    string   `json:"ranked_score"`
	Accuracy       string   `json:"accuracy"`
	PlayCount      string   `json:"play_count"`
	TotalScore     string   `json:"total_score"`
	TotalHits      string   `json:"total_hits"`
	MaximumCombo   string   `json:"maximum_combo"`
	Replays        string   `json:"replays"`
	Level          string   `json:"level"`
	SupportLvl     string   `json:"support_level"`
	DefaultGroup   string   `json:"default_group"`
	IsOnline       string   `json:"is_online"`
	IsActive       string   `json:"is_active"`
	IsDeleted      string   `json:"is_deleted"`
	IsNat          string   `json:"is_nat"`
	IsModerator    string   `json:"is_moderator"`
	IsBot          string   `json:"is_bot"`
	IsSilenced     string   `json:"is_silenced"`
	IsRestricted   string   `json:"is_restricted"`
	IsLimitedBn    string   `json:"is_limited_bn"`
	IsSupporter    string   `json:"is_supporter"`
	ProfileColor   string   `json:"profile_color"`
	PmFriendsOnly  string   `json:"pm_friends_only"`
	PostCount      string   `json:"post_count"`
	FollowersCount string   `json:"follower_count"`
}

// Функция вывода информации о пользователе Osu
func SendOsuInfo(botUrl string, update Update, username string) {

	// Значение по дефолту
	if username == "" {
		username = "hud0shnik"
	}

	// Отправка запроса моему API
	resp, err := http.Get("https://osustatsapi.vercel.app/api/user?type=string&id=" + username)

	// Проверка на ошибку
	if err != nil {
		log.Printf("http.Get error: %s", err)
		return
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var user = new(OsuUserInfo)
	json.Unmarshal(body, &user)

	// Проверка респонса
	if user.Username == "" {
		SendMsg(botUrl, update, user.Error)
		return
	}

	// Формирование текста респонса

	responseText := "Информация о <b>" + user.Username + "</b>\n"

	if user.Names[0] != "" {
		responseText += "Aka " + user.Names[0] + "\n"
	}

	responseText += "Код страны " + user.CountryCode + "\n" +
		"Рейтинг в мире <b>" + user.GlobalRank + "</b>\n" +
		"Рейтинг в стране <b>" + user.CountryRank + "</b>\n" +
		"Точность попаданий <b>" + user.Accuracy + "%</b>\n" +
		"PP <b>" + user.PP + "</b>\n" +
		"-------карты---------\n" +
		"SSH: <b>" + user.SSH + "</b>\n" +
		"SH:   <b>" + user.SH + "</b>\n" +
		"SS:   <b>" + user.SS + "</b>\n" +
		"S:     <b>" + user.S + "</b>\n" +
		"A:     <b>" + user.A + "</b>\n" +
		"---------------------------\n" +
		"Рейтинговые очки <b>" + user.RankedScore + "</b>\n" +
		"Количество игр <b>" + user.PlayCount + "</b>\n" +
		"Всего очков <b>" + user.TotalScore + "</b>\n" +
		"Всего попаданий <b>" + user.TotalHits + "</b>\n" +
		"Максимальное комбо <b>" + user.MaximumCombo + "</b>\n" +
		"Реплеев просмотрено другими <b>" + user.Replays + "</b>\n" +
		"Уровень <b>" + user.Level + "</b>\n" +
		"---------------------------\n" +
		"Время в игре " + user.PlayTime + "\n" +
		"Уровень подписки " + user.SupportLvl + "\n"

	if user.PostCount != "0" {
		responseText += "Постов на форуме " + user.PostCount + "\n"
	}

	if user.FollowersCount != "0" {
		responseText += "Подписчиков " + user.FollowersCount + "\n"
	}

	if user.IsOnline == "true" {
		responseText += "Сейчас онлайн\n"
	} else {
		responseText += "Сейчас не в сети\n"
	}

	if user.IsActive == "true" {
		responseText += "Аккаунт активен\n"
	} else {
		responseText += "Аккаунт не активен\n"
	}

	if user.IsDeleted == "true" {
		responseText += "Аккаунт удалён\n"
	}

	if user.IsBot == "true" {
		responseText += "Это аккаунт бота\n"
	}

	if user.IsNat == "true" {
		responseText += "Это аккаунт члена команды оценки номинаций\n"
	}

	if user.IsModerator == "true" {
		responseText += "Это аккаунт модератора\n"
	}

	if user.ProfileColor != "" {
		responseText += "Цвет профиля" + user.ProfileColor + "\n"
	}

	// Отправка данных пользователю
	SendPict(botUrl, update, SendPhoto{
		PhotoUrl: user.AvatarUrl,
		Caption:  responseText,
	})
}

// Функция вывода списка всех команд
func Help(botUrl string, update Update) {
	SendMsg(botUrl, update, "Привет👋🏻, вот список команд:"+"\n\n"+
		"/osu <u>username</u> - информация о пользователе Osu")
}

// Функция инициализации конфига (всех токенов)
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
