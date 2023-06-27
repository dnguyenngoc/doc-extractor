// doc-extractor/config/config.go

package config

import (
	"os"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	Config = viper.New()
	Config.Set("APIeUrl", os.Getenv("API_URL"))
	Config.Set("ApiTokenSecret", os.Getenv("API_TOKEN_SECRET"))
	Config.Set("ApiTokenExpire", os.Getenv("API_TOKEN_EXPIRE"))
	Config.Set("ApiTokenType", os.Getenv("API_TOKEN_TYPE"))
	Config.Set("ApiSyncTimeOut", os.Getenv("API_TOKEN_SECRET"))
	Config.Set("DbUri", os.Getenv("DB_URI"))
}
