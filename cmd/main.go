package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/axlts/go-telegram-bot-template/internal/config"
	"github.com/axlts/go-telegram-bot-template/internal/telegram"
)

var (
	configPath = flag.String("c", "", "path to config file")
)

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if *configPath == "" {
		panic("missing config path")
	}
	cfg, err := config.Parse(*configPath)
	if err != nil {
		panic(err)
	}

	bot, err := telegram.New(cfg.Bot)
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go bot.Run(ctx, wg)

	<-quit
	cancel()
	wg.Wait()
}
