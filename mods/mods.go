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

// Структуры для работы с OsuStatsApi

type UserInfo struct {
	Success        bool     `json:"success"`
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
	Medals         string   `json:"medals"`
}

type OnlineInfo struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Status  string `json:"status"`
}

type MapInfo struct {
	Success            bool               `json:"success"`
	Error              string             `json:"error"`
	Artist             string             `json:"artist"`
	Covers             Covers             `json:"covers"`
	Creator            string             `json:"creator"`
	FavoriteCount      string             `json:"favorite_count"`
	HypeCurrent        string             `json:"hype_current"`
	HypeRequired       string             `json:"hype_required"`
	Id                 string             `json:"id"`
	Nsfw               string             `json:"nsfw"`
	PlayCount          string             `json:"play_count"`
	PreviewUrl         string             `json:"preview_url"`
	Source             string             `json:"source"`
	Spotlight          string             `json:"spotlight"`
	Status             string             `json:"status"`
	Title              string             `json:"title"`
	UserId             string             `json:"user_id"`
	Video              string             `json:"video"`
	DownloadDisabled   string             `json:"download_disabled"`
	Bpm                string             `json:"bpm"`
	IsScoreable        string             `json:"is_scoreable"`
	LastUpdated        string             `json:"last_updated"`
	NominationsSummary NominationsSummary `json:"nominations_summary"`
	Ranked             string             `json:"ranked"`
	RankedDate         string             `json:"ranked_date"`
	Storyboard         string             `json:"storyboard"`
	Tags               []string           `json:"tags"`
	GenreName          string             `json:"genre_name"`
	LanguageName       string             `json:"language_name"`
}

type Covers struct {
	List   string `json:"list"`
	List2X string `json:"list@2x"`
}

// Оценка номинаций
type NominationsSummary struct {
	Current  string `json:"current"`
	Required string `json:"required"`
}

// Функция вывода информации о пользователе
func SendOsuInfo(botUrl string, update Update, username string) {

	// Значение по дефолту
	if username == "" {
		username = "hud0shnik"
	}

	// Отправка запроса OsuStatsApi
	resp, err := http.Get("https://osustatsapi.vercel.app/api/user?type=string&id=" + username)

	// Проверка на ошибку
	if err != nil {
		log.Printf("http.Get error: %s", err)
		return
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var user = new(UserInfo)
	json.Unmarshal(body, &user)

	// Проверка респонса
	if !user.Success {
		SendMsg(botUrl, update, user.Error)
		return
	}

	// Формирование текста респонса

	responseText := "Информация о <b>" + user.Username + "</b>\n"

	if user.Names != nil {
		responseText += "Aka " + user.Names[0] + "\n"
	}

	responseText += "Код страны " + user.CountryCode + "\n" +
		"Рейтинг в мире <b>" + user.GlobalRank + "</b>\n" +
		"Рейтинг в стране <b>" + user.CountryRank + "</b>\n" +
		"Точность попаданий <b>" + user.Accuracy + "%</b>\n" +
		"PP <b>" + user.PP + "</b>\n" +
		"-------карты---------\n" +
		"SSH: <b>" + user.SSH + "</b>\n" +
		"SH: <b>" + user.SH + "</b>\n" +
		"SS: <b>" + user.SS + "</b>\n" +
		"S: <b>" + user.S + "</b>\n" +
		"A: <b>" + user.A + "</b>\n" +
		"---------------------------\n" +
		"Рейтинговые очки <b>" + user.RankedScore + "</b>\n" +
		"Количество игр <b>" + user.PlayCount + "</b>\n" +
		"Всего очков <b>" + user.TotalScore + "</b>\n" +
		"Всего попаданий <b>" + user.TotalHits + "</b>\n" +
		"Максимальное комбо <b>" + user.MaximumCombo + "</b>\n" +
		"Реплеев просмотрено другими <b>" + user.Replays + "</b>\n" +
		"Уровень <b>" + user.Level + "</b>\n" +
		"---------------------------\n" +
		"Время в игре <i>" + user.PlayTime + "</i>\n" +
		"Достижений <i>" + user.Medals + "</i>\n"

	if user.SupportLvl != "0" {
		responseText += "Уровень подписки " + user.SupportLvl + "\n"
	}

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

// Функция вывода статуса пользователя в сети
func SendOnlineInfo(botUrl string, update Update, username string) {

	// Значение по дефолту
	if username == "" {
		username = "hud0shnik"
	}

	// Отправка запроса OsuStatsApi
	resp, err := http.Get("https://osustatsapi.vercel.app/api/online?id=" + username)

	// Проверка на ошибку
	if err != nil {
		log.Printf("http.Get error: %s", err)
		return
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response = new(OnlineInfo)
	json.Unmarshal(body, &response)

	// Проверка респонса
	if !response.Success {
		SendMsg(botUrl, update, response.Error)
		return
	}

	if response.Status == "true" {
		SendMsg(botUrl, update, "Пользователь сейчас онлайн")
	} else {
		SendMsg(botUrl, update, "Пользователь сейчас не в сети")
	}

}

// Функция отправки информации о карте
func SendMapInfo(botUrl string, update Update, beatmapset, id string) {

	if beatmapset == "" || id == "" {
		SendMsg(botUrl, update, "Синтаксис команды:\n\n/map [beatmapset] [id]\nПример:\n/map 26154 89799")
		return
	}

	// Отправка запроса OsuStatsApi
	resp, err := http.Get("https://osustatsapi.vercel.app/api/map?beatmapset=" + beatmapset + "&id=" + id)

	// Проверка на ошибку
	if err != nil {
		log.Printf("http.Get error: %s", err)
		return
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response = new(MapInfo)
	json.Unmarshal(body, &response)

	// Проверка респонса
	if !response.Success {
		SendMsg(botUrl, update, response.Error)
		return
	}

	// Формирование текста респонса

	responseText := "Информация о <b>" + response.Title + "</b>\n" +
		"Автор <i>" + response.Artist + "</i>"

	SendPict(botUrl, update, SendPhoto{
		PhotoUrl: response.Covers.List2X,
		Caption:  responseText,
	})

}

// Функция вывода списка всех команд
func Help(botUrl string, update Update) {
	SendMsg(botUrl, update, "Привет👋🏻, вот список команд:"+"\n\n"+
		"/osu <u>username</u> - информация о пользователе Osu\n"+
		"/online <u>username</u> - статус пользователя в сети")
}

// Функция инициализации конфига (всех токенов)
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
