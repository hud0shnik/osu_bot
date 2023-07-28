package commands

import "github.com/hud0shnik/osu_bot/internal/send"

// Функция вывода списка всех команд
func Help(botUrl string, chatId int) {
	send.SendMsg(botUrl, chatId, "Привет👋🏻, вот список команд:"+"\n\n"+
		"/info <u>username</u> - информация о пользователе Osu\n"+
		"/recent <u>username</u> - последняя сыгранная карта пользователя\n"+
		"/map <u>id</u> - информация о карте Osu\n"+
		"/online <u>username</u> - статус пользователя в сети")
}
