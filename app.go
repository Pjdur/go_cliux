package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func main() {
	Boxed("Google made Go", "Cliux", 40)
	Label("Cliux", "success")
	Label("Cliux", "error")
	Label("Cliux", "info")
	Section("Cliux", "Go is made by Google", 30)
	Divider("0", 25)
}

func Boxed(text string, name string, width int) {
	text = strings.TrimSpace(text)
	name = strings.TrimSpace(name)

	boxWidth := width
	fmt.Println(strings.Repeat("-", boxWidth))
	fmt.Println("¦  ", name, strings.Repeat(" ", boxWidth-len(name)-9), "  ¦")
	fmt.Println(strings.Repeat("-", boxWidth))
	fmt.Println("¦  ", text, strings.Repeat(" ", boxWidth-len(text)-9), "  ¦")
	fmt.Println(strings.Repeat("-", boxWidth))
}

func Label(content string, style string) {
	switch style {
	case "info":
		color.Blue(content)
	case "error":
		color.Red(content)
	case "success":
		color.Green(content)
	}
}

func Section(content string, name string, width int) {
	fmt.Println(name)
	fmt.Println(strings.Repeat("-", width))
	fmt.Println(content)
}

func Divider(style string,width int) {
	fmt.Println(strings.Repeat(style, width))
}