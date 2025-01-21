package main

import (
	"go-quai-monitor/config"
	"go-quai-monitor/internal"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	app := internal.NewApp(cfg)
	app.Start()
}
