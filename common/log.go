package common

import (
	"fmt"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

//fmt.Println(string(colorRed), "test")
//fmt.Println(string(colorGreen), "test")
//fmt.Println(string(colorYellow), "test")
//fmt.Println(string(colorBlue), "test")
//fmt.Println(string(colorPurple), "test")
//fmt.Println(string(colorWhite), "test")
//fmt.Println(string(colorCyan), "test")

func LogInfo(s string) {
	fmt.Println(string(colorCyan), "[*]", s, string(colorReset))
}

func LogWarning(s string) {
	fmt.Println(string(colorYellow), "[*]", s, string(colorReset))
}

func LogError(s string) {
	fmt.Println(string(colorRed), "[!]", s, string(colorReset))
}

func LogFatal(s string) {
	fmt.Println(string(colorRed), "[!]", s, string(colorReset))
}
