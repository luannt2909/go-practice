package configfx

import (
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func Initialize(configFile, configPath string) fx.Option {
	return fx.Invoke(func() {
		replacer := strings.NewReplacer(".", "_")
		viper.SetEnvKeyReplacer(replacer)
		viper.AutomaticEnv()

		// viper.SetConfigName(configFile)
		viper.SetConfigFile(configFile)
		viper.AddConfigPath(configPath)
		viper.AddConfigPath(".")

		_ = viper.ReadInConfig()
	})
}
