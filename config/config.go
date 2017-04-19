package config

import (
	"os"
	"strings"

	"statseventrelay/log"

	"github.com/spf13/viper"
)

// Configuration defaults
func init() {
	viper.SetTypeByDefaultValue(true)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/sfm/statseventrelay")
	viper.AddConfigPath("$HOME/.config/sfm/statseventrelay")
	viper.SetEnvPrefix("SFM")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Warn("error reading configuration, using default configuration")
	}
}

// Load custom configuration file
func Read(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	return viper.ReadConfig(f)
}
