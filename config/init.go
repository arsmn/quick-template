package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Init initializes application configuration from config files and/or env variables
func Init(cfgFile string) error {
	if cfgFile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath("../")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/quicktmp/")
		viper.SetConfigName("quicktmp")
	} else {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigType("toml")
	viper.SetEnvPrefix("quicktmp")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	setDefault()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return err
		}
	}

	return nil
}

func setDefault() {
	viper.SetDefault(keyAPIPort, "8080")
}
