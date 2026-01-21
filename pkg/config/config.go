package config

import (
	"ares/pkg/logger"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// require for test unit
var Param *Configuration

func init() {
	if len(os.Args) > 1 && !strings.Contains(strings.ToLower(os.Args[1]), "test") {
		// do nothing
	} else {
		// _, filename, _, _ := runtime.Caller(0)
		// dir := path.Join(path.Dir(filename), "../../../")
		// err := os.Chdir(dir)
		// if err != nil {
		// 	panic(err)
		// }
	}

	New()
}

type Configuration struct {
	Apps      Apps
	Logger    logger.Options
	Library   Library
	Database  Database
	Cache     Cache
	Parameter Parameter
}

type Apps struct {
	Name     string
	GrpcPort int
	Mode     string
}

type Library struct {
	Path string
}

type Database struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	Schema             string `json:"schema"`
	Host               string `json:"host"`
	Port               string `json:"port"`
	MinIdleConnections int    `json:"minIdleConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
	SslMode            string `json:"sslMode"`
	Location           string `json:"location"`
}

type Cache struct {
	Address string
	Index   struct {
		Security int
		Content  int
		Default  int
	}
}

type Parameter struct {
}

func getEnvironment() string {
	if len(os.Args) > 1 && !strings.Contains(strings.ToLower(os.Args[1]), "test") {
		return os.Args[1]
	}
	return "local"
}

func New() {
	env := getEnvironment()
	pathConf := "./config/" + env + ".yaml"

	viper.SetConfigFile(pathConf)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	defaultConfig := Configuration{}
	err := viper.Unmarshal(&defaultConfig)
	if err != nil {
		panic(err)
	}

	// viper.WatchConfig()
	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	errConfig := Reload(env)
	// 	if errConfig != nil {
	// 		fmt.Println("error reload config: ", errConfig.Error())
	// 	} else {
	// 		fmt.Println("Param file changed:", time.Now().Format(time.RFC1123Z))
	// 	}
	// })

	Param = &defaultConfig
}

// func Reload(mode string) error {
// 	defaultConfig := Configuration{}
// 	pathConfig := "./resources/config." + mode + ".yaml"

// 	viper.SetConfigFile(pathConfig)
// 	viper.SetConfigType("yaml")

// 	if err := viper.ReadInConfig(); err != nil {
// 		return err
// 	}

// 	err := viper.Unmarshal(&defaultConfig)
// 	if err != nil {
// 		return err
// 	}
// 	Param = &defaultConfig
// 	return nil
// }
