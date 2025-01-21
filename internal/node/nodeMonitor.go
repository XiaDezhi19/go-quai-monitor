package node

import (
	"encoding/json"
	"fmt"
	"go-quai-monitor/config"
	error2 "go-quai-monitor/internal/pkg/error"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Monitor struct {
	cfg *config.MonitorConfig
}

type RPCResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int    `json:"id"`
}

func NewMonitor(cfg *config.MonitorConfig) *Monitor {
	return &Monitor{cfg: cfg}
}

func (m *Monitor) CheckNodeSync() error {
	localNumber := blockNumber(m.cfg.LocalUrl)
	fmt.Println("Local number: ", localNumber)
	Number := blockNumber(m.cfg.NodeUrl)
	fmt.Println("correct number: ", Number)
	if Number-localNumber >= 3 {
		return error2.ErrNodeSync
	}
	return nil
}

func blockNumber(url string) uint64 {
	payload := strings.NewReader(`{
        "id": 1,
        "jsonrpc": "2.0",
        "method": "quai_blockNumber",
        "params": []
    }`)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var response RPCResponse
	_ = json.Unmarshal(body, &response)

	blockNum, _ := strconv.ParseInt(strings.TrimPrefix(response.Result, "0x"), 16, 64)
	return uint64(blockNum)
}
