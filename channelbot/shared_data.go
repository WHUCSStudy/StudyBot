package channelbot

import "github.com/tencent-connect/botgo/dto"

// UserMap Guild -> (userId -> userName)
var UserMap = make(map[string]map[string]dto.User)

// ChannelMap Guild -> (channelId -> channelName)
var ChannelMap = make(map[string]map[string]dto.Channel)

var FreshmanCourseMap = map[string]string{
	"高等数学":   "https://pd.qq.com/s/hmjpd7fzl",
	"大学物理":   "https://pd.qq.com/s/2ffmv9mh9",
	"数字电路":   "https://pd.qq.com/s/chxcd0ndc",
	"高级程序设计": "https://pd.qq.com/s/h8u8jw80c",
	"离散数学":   "https://pd.qq.com/s/1vg90cgh7",
	"线性代数":   "https://pd.qq.com/s/2yhw9kuzz",
	"大学英语":   "https://pd.qq.com/s/ggwhoywvh",
}
var SophomoreCourseMap = map[string]string{
	"数据结构":    "https://pd.qq.com/s/1udvk4wwv",
	"计算机组成原理": "https://pd.qq.com/s/6v33dnhc",
	"人工智能引论":  "https://pd.qq.com/s/w2f05x8x",
	"最优化方法":   "https://pd.qq.com/s/7pamq1zlt",
	"概率论":     "https://pd.qq.com/s/3a3dxj609",
	"大学英语":    "https://pd.qq.com/s/1ooiiut87",
}

func init() {

}
