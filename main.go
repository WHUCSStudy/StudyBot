package main

import (
	"github.com/WHUCSStudy/StudyBot/channelbot"
	"github.com/WHUCSStudy/StudyBot/groupbot"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// 启动群聊bot
	wg.Add(1)
	go groupbot.Run(&wg)

	// 启动频道bot
	wg.Add(2)
	go channelbot.Run(&wg)

	wg.Wait()

}
