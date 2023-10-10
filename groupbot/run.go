package groupbot

import (
	"github.com/WHUCSStudy/StudyBot/setup"
	"strconv"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	defer wg.Done()

	// 未启用 bot，直接结束
	enable, _ := strconv.ParseBool(setup.Config.GroupBot.Enable)
	if !enable {
		return
	}

	bot := NewBot()
	defer bot.Close()

	for text := range setup.MessageChannel {
		messageChain := NewMessageChain().BuildText(text)
		bot.SendGroupMsg(setup.Config.GroupBot.BotGroup, *messageChain)
	}

}
