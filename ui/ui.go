package ui

import (
	"fmt"
	"os"
	"time"
)

const (
	ClearScreen = "\033[H\033[2J"
	ShowCursor  = "\033[?25h"
	HideCursor  = "\033[?25l"
)

func DrawUI(text string) {
	fmt.Print(ClearScreen)
	fmt.Print(text)
}

func CreateUI(text string, sigChan chan os.Signal, ticker *time.Ticker) {
	select {
	case <-ticker.C:
		DrawUI(text)
	case <-sigChan:
		fmt.Println(ShowCursor)
		fmt.Print(ClearScreen)
		os.Exit(0)
	}
}
