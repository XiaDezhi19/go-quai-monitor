package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-quai-monitor/config"
	"go-quai-monitor/internal/pkg/models"
	"io"
	"net/http"
	"time"
)

type Sender struct {
	cfg *config.LarkConfig
}

func NewSender(cfg *config.LarkConfig) *Sender {
	return &Sender{
		cfg: cfg,
	}
}

func (s *Sender) SendLarkAlert(message string) {
	larkMsg := models.Message{
		Msg:  message,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}

	jsonData, _ := json.Marshal(larkMsg)

	resp, err := http.Post(s.cfg.WebhookUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("send msg to lark failed: %v\n", err)
		return
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read response body failed: %v", err)
	}
	println(string(body))
	defer resp.Body.Close()
}
