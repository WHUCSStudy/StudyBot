package group

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WHUCSStudy/StudyBot/logger"
	"github.com/WHUCSStudy/StudyBot/setup"
	"github.com/go-resty/resty/v2"
)

type MyBot struct {
	qq         string
	sessionKey string
}

func NewBot() MyBot {
	qq := setup.Config.GroupBot.BotQQ
	sessionKey := getSessionKey(setup.Config.GroupBot.VerifyKey)
	bindAndReleaseSession(qq, sessionKey, false)
	return MyBot{qq, sessionKey}
}

func (p *MyBot) Close() {
	bindAndReleaseSession(p.qq, p.sessionKey, true)
}

func (p *MyBot) SendGroupMsg(groupNum string, Msg MessageChain) {
	reqBody := map[string]interface{}{
		"target":       groupNum,
		"messageChain": Msg.messageChain,
	}
	marshalJsonStr, err := json.Marshal(reqBody)
	if err != nil {
		logger.Warning(err)
		return
	}

	type Result struct {
		Code      int    `json:"code"`
		Msg       string `json:"msg"`
		MessageId int    `json:"messageId"`
	}

	p.postAuthMsg(string(marshalJsonStr), "/sendGroupMessage", Result{})

}

// authentication
func (p *MyBot) postAuthMsg(body string, apiPath string, result interface{}) {
	postMsg(body, apiPath, result, p.sessionKey)
}

func bindAndReleaseSession(qq string, session string, isRelease bool) {
	apiPath := "/bind"
	if isRelease {
		apiPath = "/release"
		logger.Debug("释放会话, Q 群 Bot 登出...")
	}
	type Result struct {
		Code    int    `json:"code"`
		Session string `json:"session"`
	}
	postMsg(fmt.Sprintf(`{"sessionKey":%v,"qq":%v}`, session, qq), apiPath, Result{})
}

func getSessionKey(verifyKey string) string {
	type Result struct {
		Code    int    `json:"code"`
		Session string `json:"session"`
	}
	var result Result
	postMsg(fmt.Sprintf(`{"verifyKey":"%v"}`, verifyKey), "/verify", &result)
	return result.Session
}

func postMsg(body string, apiPath string, result interface{}, option ...string) string {
	sessionKey := ""
	if len(option) > 0 {
		sessionKey = option[0]
	}
	req := resty.New().R().SetContext(context.Background()).
		SetBody(body).
		SetHeader("sessionKey", sessionKey).
		SetHeader("Content-Type", "application/json").
		SetContentLength(true)
	resp, err := req.
		SetResult(result).
		Post(setup.Config.BaseUrl + apiPath)
	//logger.Debug(req.Header)
	if err != nil {
		logger.Error(err)
	}
	logger.Debug(resp)
	return ""
}
