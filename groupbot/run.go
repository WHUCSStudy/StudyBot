package groupbot

import (
	"github.com/WHUCSStudy/StudyBot/setup"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	defer wg.Done()

	// 未启用 bot，直接结束
	if !setup.Config.GroupBot.Enable {
		return
	}

	bot := NewBot()
	defer bot.Close()

	for text := range setup.MessageChannel {
		messageChain := NewMessageChain().BuildText(text)
		bot.SendGroupMsg(setup.Config.GroupBot.BotGroup, *messageChain)
	}

}
