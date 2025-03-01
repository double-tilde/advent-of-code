package ui

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ClearScreen = "\033[H\033[2J"
	ShowCursor  = "\033[?25h"
	HideCursor  = "\033[?25l"
)

func Draw(text string) {
	fmt.Print(ClearScreen)
	fmt.Print(text)
}

func Create(text string, sigChan chan os.Signal, ticker *time.Ticker) {
	select {
	case <-ticker.C:
		Draw(text)
	case <-sigChan:
		fmt.Println(ShowCursor)
		fmt.Print(ClearScreen)
		os.Exit(0)
	}
}

func Setup(tick time.Duration) (chan os.Signal, *time.Ticker) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(time.Second / tick)
	fmt.Print(HideCursor)

	return sigChan, ticker
}
