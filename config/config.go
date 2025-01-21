package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config 应用配置结构体
type Config struct {
	Lark    LarkConfig    `mapstructure:"lark"`
	Monitor MonitorConfig `mapstructure:"monitor"`
}

type LarkConfig struct {
	WebhookUrl string `mapstructure:"webhook_url"`
}

type MonitorConfig struct {
	NodeUrl       string `mapstructure:"node_url"`
	LocalUrl      string `mapstructure:"local_url"`
	CheckInterval int    `mapstructure:"check_interval"`
}

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	var config Config

	viper.SetConfigFile(configPath)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 将配置映射到结构体
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &config, nil
}
