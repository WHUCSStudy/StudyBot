package channel

import (
	"context"
	"fmt"
	"github.com/WHUCSStudy/StudyBot/logger"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"log"
	"path"
	"runtime"
	"time"
)

// 消息处理器，持有 openapi 对象
func init() {
}

func Run() {
	ctx := context.Background()
	overrideLogger, _ := logger.New("./", logger.FatalLevel)
	botgo.SetLogger(overrideLogger)
	// 加载 appid 和 token
	botToken := token.New(token.TypeBot)
	if err := botToken.LoadFromConfig(getConfigPath("../config.yaml")); err != nil {
		log.Fatalln(err)
	}

	// 初始化 openapi，正式环境
	api := botgo.NewOpenAPI(botToken).WithTimeout(200 * time.Second)
	// 把新的 logger 设置到 sdk 上，替换掉老的控制台 logger

	// 获取 websocket 信息
	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalln(err)
	}

	// websocket.RegisterResumeSignal(syscall.SIGUSR1)
	// 根据不同的回调，生成 intents
	intent := websocket.RegisterHandlers(
		// at 机器人事件，目前是在这个事件处理中有逻辑，会回消息，其他的回调处理都只把数据打印出来，不做任何处理
		ATMessageEventHandler(),
		// 如果想要捕获到连接成功的事件，可以实现这个回调
		ReadyHandler(),
		// 连接关闭回调
		ErrorNotifyHandler(),
		// 频道事件
		GuildEventHandler(),
		// 成员事件
		MemberEventHandler(),
		// 子频道事件
		ChannelEventHandler(),
		// 私信，目前只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		DirectMessageHandler(),
		// 频道消息，只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		CreateMessageHandler(),
		// 互动事件
		InteractionHandler(),
		// 发帖事件
		ThreadEventHandler(),
	)
	// 指定需要启动的分片数为 2 的话可以手动修改 wsInfo
	if err = botgo.NewSessionManager().Start(wsInfo, botToken, &intent); err != nil {
		log.Fatalln(err)
	}
}

func getConfigPath(name string) string {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s/%s", path.Dir(filename), name)
	}
	return ""
}
