package channelbot

import "C"
import (
	"bytes"
	"context"
	"fmt"
	"github.com/WHUCSStudy/StudyBot/logger"
	"github.com/WHUCSStudy/StudyBot/setup"
	"github.com/go-resty/resty/v2"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

func MentionUser(channelId string, userId string) {
	_, err := Api.PostMessage(context.Background(), channelId,
		&dto.MessageToCreate{
			Content: message.MentionUser(userId) + " 请及时选出优秀回答~",
		})
	if err != nil {
		logger.Warning(err)
		return
	}
}
func SendPicToChannelMsg(
	channelID string,
	qrContent []byte,
	data map[string]string) ([]byte, error) {
	resp, err := resty.New().R().SetContext(context.Background()).SetAuthScheme("Bot").
		SetAuthToken(setup.Config.ChannelBot.Appid+"."+setup.Config.ChannelBot.Token).
		SetFormData(data).
		SetFileReader("file_image", "qrcode.png", bytes.NewReader(qrContent)).
		SetContentLength(true).
		SetResult(dto.Message{}).
		SetPathParam("channel_id", channelID).
		Post(fmt.Sprintf("%s://%s%s", "https", "api.sgroup.qq.com", "/channels/{channel_id}/messages"))
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func GetUserById(guildId string, userId string) dto.User {

	if users := UserMap[guildId]; len(users) > 0 {
		return users[userId]
	}

	// UserMap 未记录该主频道, 获取频道用户
	members, err := GetChannelMembers(guildId)
	if err != nil {
		logger.Warning("获取频道成员失败：", err)
		return dto.User{}
	}
	// 完善 map
	UserMap[guildId] = make(map[string]dto.User)
	for _, member := range members {
		UserMap[guildId][member.ID] = member
	}

	return UserMap[guildId][userId]

}

// GetChannelMembers 查询频道内所有成员（该 API 仅私域机器人可调用，难绷）
func GetChannelMembers(guildId string) ([]dto.User, error) {
	type tempUser struct {
		User dto.User `json:"user"`
	}
	var tempUsers = make([]tempUser, 0)
	postAuthMsg(map[string]string{"limit": "400"},
		&tempUsers, "/guilds/{guild_id}/members",
		map[string]string{"guild_id": guildId})
	var users = make([]dto.User, 0)
	for _, elem := range tempUsers {
		users = append(users, elem.User)
	}
	return users, nil
}

func GetChannelById(guildId string, channelId string) dto.Channel {

	if channels := ChannelMap[guildId]; len(channels) > 0 {
		return channels[channelId]
	}

	// ChannelMap 未配置，配置
	channels, err := Api.Channels(context.Background(), guildId)
	if err != nil {
		logger.Warning("获取频道失败：", err)
		return dto.Channel{}
	}

	// 完善 map
	ChannelMap[guildId] = make(map[string]dto.Channel)
	for _, channel := range channels {
		ChannelMap[guildId][channel.ID] = *channel
	}
	return ChannelMap[guildId][channelId]
}

func postAuthMsg(
	query map[string]string,
	result interface{},
	apiPath string,
	pathParam map[string]string,
) {
	_, err := resty.New().R().SetContext(context.Background()).SetAuthScheme("Bot").
		SetAuthToken(setup.Config.ChannelBot.Appid + "." + setup.Config.ChannelBot.Token).
		SetContentLength(true).
		SetQueryParams(query).
		SetResult(result).
		SetPathParams(pathParam).
		Get(fmt.Sprintf("%s://%s%s", "https", "api.sgroup.qq.com", apiPath))
	if err != nil {
		logger.Warning(err)
		return
	}
}
