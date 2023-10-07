package main

import (
	"github.com/WHUCSStudy/StudyBot/channel"
	"github.com/WHUCSStudy/StudyBot/group"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// 启动群聊bot
	wg.Add(1)
	go group.Run(&wg)

	// 启动频道bot
	wg.Add(2)
	go channel.Run(&wg)

	wg.Wait()

}
