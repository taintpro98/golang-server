package config

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/inhies/go-bytesize"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const KEY_CONFIG_URL = "--config_url"

type Config struct {
	AppInfo     ConfigAppInfo    `mapstructure:"app_info"`
	DBM         DatabaseConfig   `mapstructure:"dbm"`
	Database    DatabaseConfig   `mapstructure:"database"`
	Mongo       MongoConfig      `mapstructure:"mongo"`
	Redis       RedisConfig      `mapstructure:"redis"`
	RedisPubSub RedisConfig      `mapstructure:"redis_pub_sub"`
	RedisQueue  RedisQueueConfig `mapstructure:"redis_queue"`
	Kafka       KafkaConfig      `mapstructure:"kafka"`
	TelegramBot TelegramBot      `mapstructure:"telegram_bot"`
	Token       Token            `mapstructure:"token"`
	Elastic     ElasticConfig    `mapstructure:"elastic"`
}

type KafkaConfig struct {
	Uri         string     `mapstructure:"uri"`
	Consumer    string     `mapstructure:"consumer"`
	Partitioner string     `mapstructure:"partitioner"` // "The partitioning scheme to use. Can be `hash`, `manual`, or `random`")
	Topic       KafkaTopic `mapstructure:"topic"`
	// Sasl        KafkaSasl  `mapstructure:"sasl"`
}

type KafkaTopic struct {
	UserTopic      string `mapstructure:"user_topic"`
	MessageChannel string `mapstructure:"message_channel"`
}

type RedisQueueConfig struct {
	Prefix   string `mapstructure:"prefix"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type RedisConfig struct {
	Prefix   string `mapstructure:"prefix"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type ConfigAppInfo struct {
	Environment   string `mapstructure:"environment"`
	ApiPort       string `mapstructure:"api_port"`
	WebsocketPort string `mapstructure:"websocket_port"`
}

type DatabaseConfig struct {
	Schema       string `mapstructure:"schema"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"database_name"`
	SSLMode      string `mapstructure:"sslmode"`
}

type MongoConfig struct {
	DatabaseName string `mapstructure:"database_name"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
}

type TelegramBot struct {
	ChatID int64  `mapstructure:"chat_id"`
	Token  string `mapstructure:"token"`
}

type Token struct {
	PrivateKeyPath            string        `mapstructure:"priviate_key_path"`
	PublicKeyPath             string        `mapstructure:"public_key_path"`
	AccessTokenDuration       time.Duration `mapstructure:"access_token_duration"`
	RefreshTokenDuration      time.Duration `mapstructure:"refresh_token_duration"`
	AdminAccessTokenDuration  time.Duration `mapstructure:"admin_access_token_duration"`
	AdminRefreshTokenDuration time.Duration `mapstructure:"admin_refresh_token_duration"`
	EmddiSecretKey            string        `mapstructure:"emddi_secret_key"`
	CacheInMem                bool          `mapstructure:"cache_in_mem"`
	ForeverTokenRoles         []string      `mapstructure:"forever_token_roles"`
}

type ElasticConfig struct {
	Addresses            string `mapstructure:"addresses"`
	Username             string `mapstructure:"username"`
	Password             string `mapstructure:"password"`
	SupplierSettingIndex string `mapstructure:"supplier_setting_index"`
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

func loadConfigByLocalPath(v *viper.Viper, envi string) {
	configName := "config.local"
	if envi != "" {
		configName = fmt.Sprintf("config.%s", envi)
	}
	v.SetConfigName(configName)
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

func Init(envi string) (config Config) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.AutomaticEnv()

	configUrl := getConfigUrl()
	if configUrl == "" {
		loadConfigByLocalPath(v, envi)
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
