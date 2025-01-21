package internal

import (
	"fmt"
	"go-quai-monitor/config"
	"go-quai-monitor/internal/lark"
	"go-quai-monitor/internal/node"
	"go-quai-monitor/internal/proxy"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

type App struct {
	NodeMonitor  *node.Monitor
	ProxyMonitor *proxy.Monitor
	Sender       *lark.Sender
	ticker       *time.Ticker
	msgChan      chan string
	doneChan     chan struct{}
}

func NewApp(cfg *config.Config) *App {
	return &App{
		NodeMonitor:  node.NewMonitor(&cfg.Monitor),
		ProxyMonitor: &proxy.Monitor{},
		Sender:       lark.NewSender(&cfg.Lark),
		ticker:       time.NewTicker(time.Duration(cfg.Monitor.CheckInterval) * time.Minute),
		msgChan:      make(chan string, 10),
	}
}

func (a *App) StartMonitor() {
	defer a.ticker.Stop()
	for {
		select {
		case <-a.ticker.C:
			fmt.Println("start monitor")
			err := a.NodeMonitor.CheckNodeSync()
			if err != nil {
				a.msgChan <- err.Error()
			}
		}
	}
}

func (a *App) StartSend() {
	defer wg.Done()
	for {
		select {
		case msg := <-a.msgChan:
			a.Sender.SendLarkAlert(msg)
		}

	}
}

func (a *App) Start() {
	wg.Add(2)
	go a.StartSend()
	go a.StartMonitor()
	wg.Wait()
}

func (a *App) Stop() {

}
