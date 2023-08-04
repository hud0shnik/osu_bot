package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hud0shnik/osu_bot/internal/telegram"
	"github.com/sirupsen/logrus"
)

// Структуры для работы с OsuStatsApi

type userInfo struct {
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
	Status bool `json:"status"`
}

type mapInfo struct {
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

type historicalResponse struct {
	Recent struct {
		Items []recent `json:"items"`
	} `json:"recent"`
}

type recent struct {
	BeatmapID string `json:"beatmap_id"`
	Passed    string `json:"passed"`
	Rank      string `json:"rank"`
	PP        string `json:"pp"`
	Beatmap   struct {
		URL              string `json:"url"`
		DifficultyRating string `json:"difficulty_rating"`
		Version          string `json:"version"`
	} `json:"beatmap"`
	Beatmapset struct {
		Title string `json:"title"`
	} `json:"beatmapset"`
}

// Функция вывода информации о пользователе
func SendUserInfo(botUrl string, chatId int, username string) {

	// Проверка параметра
	if username == "" {
		telegram.SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/info <b>[id]</b>\n\nПример:\n/info <b>hud0shnik</b>")
		return
	}

	// Отправка запроса OsuStatsApi
	resp, err := http.Get("https://osustatsapi.vercel.app/api/user?type=string&id=" + username)

	// Проверка на ошибку
	if err != nil {
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		logrus.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		telegram.SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		telegram.SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)
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
	telegram.SendPict(botUrl, chatId, user.AvatarUrl, responseText)

}

// Функция вывода статуса пользователя в сети
func SendOnlineInfo(botUrl string, chatId int, username string) {

	// Проверка параметра
	if username == "" {
		telegram.SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/online <b>[id]</b>\n\nПример:\n/online <b>hud0shnik</b>")
		return
	}

	// Отправка запроса OsuStatsApi
	resp, err := http.Get("https://osustatsapi.vercel.app/api/online?id=" + username)

	// Проверка на ошибку
	if err != nil {
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		logrus.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		telegram.SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		telegram.SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)
	var response = new(onlineInfo)
	json.Unmarshal(body, &response)

	if response.Status {
		telegram.SendMsg(botUrl, chatId, "Пользователь сейчас онлайн")
	} else {
		telegram.SendMsg(botUrl, chatId, "Пользователь сейчас не в сети")
	}

}

// Функция отправки информации о карте
func SendMapInfo(botUrl string, chatId int, id string) {

	// Проверка на пустой id
	if id == "" {
		telegram.SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/map <b>[id]</b>\n\nПример:\n/map <b>89799</b>")
		return
	}

	// Отправка запроса OsuStatsApi
	resp, err := http.Get("https://osustatsapi.vercel.app/api/map?type=string&id=" + id)

	// Проверка на ошибку
	if err != nil {
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		logrus.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		telegram.SendMsg(botUrl, chatId, "Карта не найдена")
		return
	case 400:
		telegram.SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)
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

	telegram.SendPict(botUrl, chatId, response.Covers.List2X, responseText)

}

// Функция отправки последней сыгранной карты
func SendRecentBeatmap(botUrl string, chatId int, username string) {

	// Проверка параметра
	if username == "" {
		telegram.SendMsg(botUrl, chatId, "Синтаксис команды:\n\n/recent <b>[id]</b>\n\nПример:\n/recent <b>hud0shnik</b>")
		return
	}

	// Отправка запроса OsuStatsApi для поиска пользователя
	resp, err := http.Get("https://osustatsapi.vercel.app/api/user?type=string&id=" + username)
	if err != nil {
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		logrus.Printf("http.Get error: %s", err)
		return
	}
	defer resp.Body.Close()

	// Проверка респонса
	switch resp.StatusCode {
	case 200:
		// При хорошем статусе респонса продолжение выполнения кода
	case 404:
		telegram.SendMsg(botUrl, chatId, "Пользователь не найден")
		return
	case 400:
		telegram.SendMsg(botUrl, chatId, "Плохой реквест")
		return
	default:
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		return
	}

	// Запись респонса
	body, _ := io.ReadAll(resp.Body)
	var user = new(struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	})
	json.Unmarshal(body, &user)

	// Отправка запроса OsuStatsApi для получения последней активности
	resp2, err := http.Get("https://osustatsapi.vercel.app/api/historical?type=string&id=" + user.ID)
	if err != nil {
		telegram.SendMsg(botUrl, chatId, "Внутренняя ошибка")
		logrus.Printf("http.Get error: %s", err)
		return
	}
	defer resp2.Body.Close()

	// Запись респонса
	body, _ = io.ReadAll(resp2.Body)
	var historical historicalResponse
	json.Unmarshal(body, &historical)

	// Проверка на наличие активности
	if len(historical.Recent.Items) == 0 {
		telegram.SendMsg(botUrl, chatId, "Пользователь <i>"+user.Username+"</i> не играл карты за последние 24 часа")
		return
	}

	// Вывод информации о последней сыгранной карте
	recentScore := historical.Recent.Items[0]
	telegram.SendMsg(botUrl, chatId, "Последняя сыгранная карта <i>"+user.Username+
		"</i> - <b>"+recentScore.Beatmapset.Title+"</b>\nНа сложности <b>"+
		recentScore.Beatmap.Version+"</b> <i>("+recentScore.Beatmap.DifficultyRating+")</i>\n"+recentScore.Beatmap.URL)
	SendMapInfo(botUrl, chatId, recentScore.BeatmapID)

	// Проверка на результат игры
	if recentScore.Passed == "true" {
		telegram.SendMsg(botUrl, chatId, "<i>"+user.Username+"</i> прошёл её на <b>"+recentScore.Rank+"</b> получив <b>"+recentScore.PP+"</b> pp")
	} else {
		telegram.SendMsg(botUrl, chatId, "<i>"+user.Username+"</i> не прошёл её :^(")
	}

}
