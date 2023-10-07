package setup

import (
	"github.com/spf13/viper"
	"log"
	"reflect"
)

var Config = struct {
	ChannelBot struct {
		Appid string `yaml:"appid"`
		Token string `yaml:"token"`
	} `yaml:"channelBot"`

	BaseUrl  string `yaml:"baseUrl,omitempty"`
	LogLevel string `yaml:"logLevel,omitempty"`
	IsUnix   string `yaml:"isUnix,omitempty"`

	GroupBot struct {
		VerifyKey string `yaml:"verifyKey"`
		BotQQ     string `yaml:"botQQ"`
	} `yaml:"groupBot"`
	Test struct {
		AA string `yaml:"aa"`
	} `yaml:"test"`
}{
	LogLevel: "debug",
	IsUnix:   "default",
	BaseUrl:  "http://127.0.0.1:8087",
}

func init() {

	viper.SetConfigFile("./config.yaml")
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("请您添加配置文件（将 config.yaml.demo 重命名为 config.yaml）")
		return // 自动退出
	}

	setConf(reflect.ValueOf(&Config))
}

// Elem()用于获取指针指向的值，如果不是接口或指针会panics
// Addr()用于获得值的指针
func setConf(value reflect.Value, lastFields ...string) {
	for i := 0; i < value.Elem().NumField(); i++ {
		field := value.Elem().Field(i)
		if field.Kind() == reflect.String {
			resKey := ""
			for _, lastField := range lastFields {
				resKey += lastField + "."
			}
			resKey += value.Type().Elem().Field(i).Name
			if tempParam := viper.GetString(resKey); tempParam != "" {
				field.Set(reflect.ValueOf(tempParam))
			}
		} else {
			// 回溯 (前进 => 处理 => 回退)
			lastFields = append(lastFields, value.Elem().Type().Field(i).Name)
			setConf(field.Addr(), lastFields...)
			lastFields = lastFields[:len(lastFields)-1]
		}
	}
}

//func getConfigPath(name string) string {
//	_, filename, _, ok := runtime.Caller(1)
//	if ok {
//		return fmt.Sprintf("%s/%s", path.Dir(filename), name)
//	}
//	return ""
//}
