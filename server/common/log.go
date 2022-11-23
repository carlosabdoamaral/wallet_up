package common

import (
	"fmt"
	"log"
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

func PrintStartMethod(s string) {
	log.Println("")
	PrintInfo(fmt.Sprintf("Received %s request", s))
}

func PrintInfo(s string) {
	s = fmt.Sprintf("%s%s%s", string(colorBlue), s, string(colorReset))
	log.Println(s)
}

func PrintSuccess(s string) {
	s = fmt.Sprintf("%s%s%s", string(colorGreen), s, string(colorReset))
	log.Println(s)
}

func PrintWarning(s string) {
	s = fmt.Sprintf("%s%s%s", string(colorYellow), s, string(colorReset))
	log.Println(s)
}

func PrintError(s string) {
	s = fmt.Sprintf("%s%s%s", string(colorRed), s, string(colorReset))
	log.Println(s)
}

func PrintFatal(s string) {
	fmt.Printf("%s%s%s\n", string(colorRed), s, string(colorReset))
	log.Fatalln()
}
