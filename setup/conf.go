package setup

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
)

var Config = struct {
	ChannelBot struct {
		Enable bool   `yaml:"enable"`
		Appid  string `yaml:"appid"`
		Token  string `yaml:"token"`
	} `yaml:"channelBot"`

	LogLevel string `yaml:"logLevel,omitempty"`
	GroupBot struct {
		Enable    bool   `yaml:"enable"`
		VerifyKey string `yaml:"verifyKey"`
		BotQQ     string `yaml:"botQQ"`
		BotGroup  string `yaml:"botGroup"`
		BaseUrl   string `yaml:"baseUrl"`
	} `yaml:"groupBot"`
	Test struct {
		AA string `yaml:"aa"`
	} `yaml:"test"`
}{
	LogLevel: "debug",
}

func init() {
	// 默认只启用频道机器人
	Config.ChannelBot.Enable = true
	Config.GroupBot.Enable = false

	viper.SetConfigFile(GetAbsolutePath("config.yaml"))
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		//log.Println(GetAbsolutePath("config.yaml"))
		log.Println("请您添加配置文件（将 config.yaml.demo 重命名为 config.yaml）")
		os.Exit(0)
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

// GetAbsolutePath 获取从程序根路径开始的绝对路径
func GetAbsolutePath(filePath string) (absolutePath string) {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		// 从源码获取，尝试查找源码根路径的 confName 文件
		absolutePath = fmt.Sprintf("%s/../%s", path.Dir(filename), filePath)
	}
	_, err := os.Stat(absolutePath)
	if os.IsNotExist(err) {
		// 未找到则直接返回当前目录下的文件
		absolutePath = "./" + filePath
	}
	_, err = os.Stat(absolutePath)
	if os.IsNotExist(err) {
		// 仍未找到则直接返回当前目录下的文件
		log.Fatal(fmt.Sprintf("%s/../%s", path.Dir(filename), filePath))
	}

	return
}
