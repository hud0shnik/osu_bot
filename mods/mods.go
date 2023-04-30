package mods

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Структуры для работы с OsuStatsApi

type userInfo struct {
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

type onlineInfo struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Status  bool   `json:"status"`
}

type mapInfo struct {
	Success            bool               `json:"success"`
	Error              string             `json:"error"`
	Artist             string             `json:"artist"`
	Covers             covers             `json:"covers"`
	Creator            string             `json:"creator"`
	FavoriteCount      string             `json:"favorite_count"`
	HypeCurrent        string             `json:"hype_current"`
	HypeRequired       string             `json:"hype_required"`
	Nsfw               string             `json:"nsfw"`
	PlayCount          string             `json:"play_count"`
	Spotlight          string             `json:"spotlight"`
	Status             string             `json:"status"`
	Title              string             `json:"title"`
	Video              string             `json:"video"`
	Bpm                string             `json:"bpm"`
	IsScoreable        string             `json:"is_scoreable"`
	NominationsSummary nominationsSummary `json:"nominations_summary"`
	Ranked             string             `json:"ranked"`
	Storyboard         string             `json:"storyboard"`
	GenreName          string             `json:"genre_name"`
	LanguageName       string             `json:"language_name"`
}

type covers struct {
	List2X string `json:"list@2x"`
}

type nominationsSummary struct {
	Current  string `json:"current"`
	Required string `json:"required"`
}

// Функция вывода информации о пользователе
func SendUserInfo(botUrl string, chatId int, username string) {

	// Проверка параметра
	if username == "" {
		SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/info <b>[id]</b>\n\nПример:\n/info <b>hud0shnik</b>")
		return
	}

	// Отправка запроса OsuStatsApi
	resp, err := http.Get("https://osustatsapi.vercel.app/api/v2/user?type=string&id=" + username)

	// Проверка на ошибку
	if err != nil {
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		log.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := ioutil.ReadAll(resp.Body)
	var user = new(userInfo)
	json.Unmarshal(body, &user)

	// Формирование текста респонса

	responseText := "Информация о <b>" + user.Username + "</b>\n"

	if len(user.Names) != 0 {
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
		responseText += "Цвет профиля " + user.ProfileColor + "\n"
	}

	// Отправка данных пользователю
	SendPict(botUrl, chatId, user.AvatarUrl, responseText)

}

// Функция вывода статуса пользователя в сети
func SendOnlineInfo(botUrl string, chatId int, username string) {

	// Проверка параметра
	if username == "" {
		SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/online <b>[id]</b>\n\nПример:\n/online <b>hud0shnik</b>")
		return
	}

	// Отправка запроса OsuStatsApi
	resp, err := http.Get("https://osustatsapi.vercel.app/api/v2/online?id=" + username)

	// Проверка на ошибку
	if err != nil {
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		log.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := ioutil.ReadAll(resp.Body)
	var response = new(onlineInfo)
	json.Unmarshal(body, &response)

	if response.Status {
		SendMsg(botUrl, chatId, "Пользователь сейчас онлайн")
	} else {
		SendMsg(botUrl, chatId, "Пользователь сейчас не в сети")
	}

}

// Функция отправки информации о карте
func SendMapInfo(botUrl string, chatId int, beatmapset, id string) {

	// Проверка параметров
	if beatmapset == "" || id == "" {
		SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/map <b>[beatmapset] [id]</b>\n\nПример:\n/map <b>26154 89799</b>")
		return
	}

	// Отправка запроса OsuStatsApi
	resp, err := http.Get("https://osustatsapi.vercel.app/api/v2/map?type=string&beatmapset=" + beatmapset + "&id=" + id)

	// Проверка на ошибку
	if err != nil {
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		log.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := ioutil.ReadAll(resp.Body)
	var response = new(mapInfo)
	json.Unmarshal(body, &response)

	// Формирование текста респонса

	responseText := "Информация о <b>" + response.Title + "</b> - <i>" + response.Artist + "</i>\n" +
		"Маппер <i>" + response.Creator + "</i>\n" +
		"Статус карты <b>" + response.Status + "</b>\n" +
		"Количество игр <b>" + response.PlayCount + "</b>\n" +
		"В избранных у <b>" + response.FavoriteCount + "</b>\n" +
		"Bpm <b>" + response.Bpm + "</b>\n" +
		"Жанр <b>" + response.GenreName + "</b>\n" +
		"Язык <b>" + response.LanguageName + "</b>\n"

	if response.HypeRequired != "" {
		responseText += "Хайп <b>" + response.HypeCurrent + "</b>/<b>" + response.HypeRequired + "</b>\n"
	}

	if response.NominationsSummary.Required != "" {
		responseText += "Номинации <b>" + response.NominationsSummary.Current + "</b>/<b>" + response.NominationsSummary.Required + "</b>\n"
	}

	if response.Spotlight == "true" {
		responseText += "Спотлайт карта\n"
	}

	if response.Nsfw == "true" {
		responseText += "NSFW карта\n"
	}

	if response.Video == "true" {
		responseText += "Есть видео\n"
	}

	if response.IsScoreable == "true" {
		responseText += "Есть таблица рекордов\n"
	}

	if response.Ranked == "1" {
		responseText += "Рейтинговая\n"
	}

	if response.Storyboard == "true" {
		responseText += "Есть сториборда"
	}

	SendPict(botUrl, chatId, response.Covers.List2X, responseText)

}

// Функция вывода списка всех команд
func Help(botUrl string, chatId int) {
	SendMsg(botUrl, chatId, "Привет👋🏻, вот список команд:"+"\n\n"+
		"/info <u>username</u> - информация о пользователе Osu\n"+
		"/map <u>beatmapset id</u> - информация о карте Osu\n"+
		"/online <u>username</u> - статус пользователя в сети")
}
