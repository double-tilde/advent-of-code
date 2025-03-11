package ui

import (
	"aoc-2024/model"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ClearScreen       = "\033[H\033[2J"
	ShowCursor        = "\033[?25h"
	HideCursor        = "\033[?25l"
	Reset             = "\033[0m"
	GreenBgBlackText  = "\033[42;30m"
	YellowBgBlackText = "\033[43;30m"
)

func StringColor(s, c string) string {
	s = c + s + Reset
	return s
}

func Setup(tick time.Duration) (chan os.Signal, *time.Ticker) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(time.Second / tick)
	fmt.Print(HideCursor)

	return sigChan, ticker
}

func Draw(text string) {
	fmt.Print(text)
}

func Matrix(
	matrix [][]string,
	searchLen int,
	word []model.WordPosition,
	sigChan chan os.Signal,
	ticker *time.Ticker,
) {
	highlightMap := make(map[[2]int]string)

	var sb []model.WordPosition
	for _, char := range word {
		var uiMatrix string
		sb = append(sb, char)
		if len(sb) == searchLen {
			for _, b := range sb {
				highlightMap[[2]int{b.Row, b.Col}] = GreenBgBlackText
			}
		} else {
			highlightMap[[2]int{char.Row, char.Col}] = YellowBgBlackText
		}

		for j := range matrix {
			for k := range matrix[j] {
				if color, exists := highlightMap[[2]int{j, k}]; exists {
					uiMatrix += StringColor(matrix[j][k], color)
				} else {
					uiMatrix += matrix[j][k]
				}
			}
			uiMatrix += "\n"
		}
		Create(uiMatrix, sigChan, ticker)
	}
}

func Create(text string, sigChan chan os.Signal, ticker *time.Ticker) {
	select {
	case <-ticker.C:
		fmt.Print(ClearScreen)
		Draw(text)
	case <-sigChan:
		fmt.Println(ShowCursor)
		fmt.Print(ClearScreen)
		os.Exit(0)
	}
}
