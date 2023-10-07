package channel

import (
	"context"
	"fmt"
	"github.com/WHUCSStudy/StudyBot/logger"
	"github.com/WHUCSStudy/StudyBot/setup"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/event"
	"regexp"
	"strings"
)

// ReadyHandler 自定义 ReadyHandler 感知连接成功事件
func ReadyHandler() event.ReadyHandler {
	return func(event *dto.WSPayload, data *dto.WSReadyData) {
		logger.InfoF("准备接收来自频道机器人「%v」的事件... ", data.User.Username)
	}
}

func ErrorNotifyHandler() event.ErrorNotifyHandler {
	return func(err error) {
		//log.Println("error notify receive: ", err)
	}
}

// ATMessageEventHandler 实现处理 at 消息的回调
func ATMessageEventHandler() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		msg := strings.ToLower(message.ETLInput(data.Content)) // 去掉@符号和首尾空格，同时全改小写
		logger.Debug("收到艾特消息：", msg)
		return nil // 处理数据
	}
}

// GuildEventHandler 处理频道事件
func GuildEventHandler() event.GuildEventHandler {
	return func(event *dto.WSPayload, data *dto.WSGuildData) error {
		//fmt.Println(data)
		return nil
	}
}

// ChannelEventHandler 处理子频道事件
func ChannelEventHandler() event.ChannelEventHandler {
	return func(event *dto.WSPayload, data *dto.WSChannelData) error {
		//fmt.Println(data)
		return nil
	}
}

// MemberEventHandler 处理成员变更事件
func MemberEventHandler() event.GuildMemberEventHandler {
	return func(event *dto.WSPayload, data *dto.WSGuildMemberData) error {
		//fmt.Println(data)
		return nil
	}
}

// DirectMessageHandler 处理私信事件
func DirectMessageHandler() event.DirectMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSDirectMessageData) error {
		//fmt.Println(data.Content)

		return nil
	}
}

// CreateMessageHandler 处理消息事件
func CreateMessageHandler() event.MessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSMessageData) error {
		//fmt.Println(data)
		return nil
	}
}

// InteractionHandler 处理内联交互事件
func InteractionHandler() event.InteractionEventHandler {
	return func(event *dto.WSPayload, data *dto.WSInteractionData) error {
		//fmt.Println(data)
		return nil
	}
}

var isCall = true

// ThreadEventHandler 论坛主贴事件
func ThreadEventHandler() event.ThreadEventHandler {
	return func(event *dto.WSPayload, data *dto.WSThreadData) error {
		// sdk 不完善，存在多次调用的情况
		// 这里交替调用解决问题
		if !isCall {
			isCall = true
			return nil
		}

		logger.Debug("接收到论坛事件：", event.Type)
		logger.Debug(data.ThreadInfo.Title)
		title := func(toBeDecoded string) string {
			matches := regexp.
				MustCompile("\\{\"text\":\\{\"text\":\"([\\S\\s]+)\"},\"type\":1}").
				FindAllStringSubmatch(toBeDecoded, -1)

			for _, elem := range matches {
				return strings.ReplaceAll(elem[1], "\\", "")
			}
			return ""
		}(data.ThreadInfo.Title)
		logger.Debug(title)
		subject, _ := Api.Channel(context.Background(), data.ChannelID)

		author := GetUserById(data.GuildID, data.AuthorID)
		text := fmt.Sprintf("[%v] [%v] %v", subject.Name, author.Username, title)
		logger.InfoF(text)
		setup.MessageChannel <- text
		isCall = false
		return nil
	}
}
