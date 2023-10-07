package channel

import "github.com/tencent-connect/botgo/dto"

// UserMap Guild -> (userId -> userName)
var UserMap = make(map[string]map[string]dto.User)
