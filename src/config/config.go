package config

import (
	"logger"
	"strings"

	"github.com/spf13/viper"
)

var log = logger.GetLogger()

var cfg *viper.Viper

func Init() {
	cfg = viper.New()
	cfg.AutomaticEnv()
	cfg.SetEnvPrefix("fdd")
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	log.Debug("config initialized")
}

func GetConfig() *viper.Viper {
	return cfg
}
