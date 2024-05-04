package config

import (
	"fmt"
	"github.com/inhies/go-bytesize"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"reflect"
	"strings"
)

const KEY_CONFIG_URL = "--config_url"

type Config struct {
	AppInfo     ConfigAppInfo `mapstructure:"app_info"`
	TelegramBot TelegramBot   `mapstructure:"telegram_bot"`
}

type ConfigAppInfo struct {
	Environment string `mapstructure:"environment"`
	ApiPort     string `mapstructure:"api_port"`
}

type TelegramBot struct {
	Token string `mapstructure:"token"`
}

// hex string to bytesize.ByteSize.
func StringToByteSizeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(bytesize.B) {
			return data, nil
		}

		sDec, err := bytesize.Parse(data.(string))
		if err != nil {
			return nil, err
		}
		return sDec, nil
	}
}

func getConfigUrl() string {
	args := os.Args[1:]
	configUrl := ""

	for ix, x := range args {
		if x[:2] == "--" && x == KEY_CONFIG_URL {
			configUrl = args[ix+1]
		}
	}

	return configUrl
}

func loadConfigByLocalPath(v *viper.Viper) {
	v.SetConfigName("config.local")
	v.AddConfigPath(".")          // Look for config in current directory
	v.AddConfigPath("config/")    // Optionally look for config in the working directory.
	v.AddConfigPath("../config/") // Look for config needed for tests.
	v.AddConfigPath("../")        // Look for config needed for tests.

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {
		v.SetConfigName("config") // Sử dụng config.yaml nếu không tìm thấy config.local.yaml
		err = v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
}

func loadConfigByUrl(v *viper.Viper, configURL string) {
	// down file from url
	response, err := http.Get(configURL)
	if err != nil {
		fmt.Printf("Error downloading config file: %s\n", err)
		return
	}
	defer response.Body.Close()

	err = v.ReadConfig(response.Body) // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func Init() (config Config) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.AutomaticEnv()

	configUrl := getConfigUrl()
	if configUrl == "" {
		loadConfigByLocalPath(v)
	} else {
		loadConfigByUrl(v, configUrl)
	}
	err := v.Unmarshal(
		&config, viper.DecodeHook(
			mapstructure.ComposeDecodeHookFunc(
				mapstructure.StringToTimeDurationHookFunc(),
				mapstructure.RecursiveStructToMapHookFunc(),
				StringToByteSizeHookFunc(),
			),
		),
	)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	return
}
