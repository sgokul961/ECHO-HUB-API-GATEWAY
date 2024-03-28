package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"PORT"`
	AuthHubUrl         string `mapstructure:"auth_hub_url"`
	PostHubUrl         string `mapstructure:"post_hub_url"`
	NotificationHubUrl string `mapstructure:"notification_hub_url"`
	ChatHubUrl         string `mapstructure:"chat_hub_url"`
	JWTSecretKey       string `mapstructure:"JWT_SECRET_KEY"`

	//add image to aws s3 bucket
	AWS_REGION            string `mapstructure:"AWS_REGION"`
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	fmt.Println("print c", c)
	return

}
