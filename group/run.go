package group

import (
	"github.com/WHUCSStudy/StudyBot/setup"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	defer wg.Done()

	bot := NewBot()
	defer bot.Close()

	for text := range setup.MessageChannel {
		messageChain := NewMessageChain().BuildText(text)
		bot.SendGroupMsg("***REMOVED***", *messageChain)
	}

}
