package util

import (
	"time"

	"github.com/spf13/viper"
)

// * Note [codermuss]: Config stores all configuraiton of the application
// * Note [codermuss]: The values are read by viper from a config file or environment variables.

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrationUrl         string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymetricKey     string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// * Note [codermuss]: LoadConfig reads configuration from file or environment variables.

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // * Note [codermuss]: Any other format like json, xml
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	// * Note [codermuss]: return config or err
	err = viper.Unmarshal(&config)
	return
}
