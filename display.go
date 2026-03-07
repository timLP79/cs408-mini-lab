package main

import (
	"fmt"

	"github.com/fatih/color"
)

func displayModules(courseName string, modules []Module) {
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)

	fmt.Printf("\n%s\n", courseName)
	fmt.Println("------------------------------")

	maxLen := 0
	for _, m := range modules {
		if len(m.Name) > maxLen {
			maxLen = len(m.Name)
		}
	}

	for _, m := range modules {
		itemWord := "items"
		if m.ItemsCount == 1 {
			itemWord = "item"
		}
		format := fmt.Sprintf("%%s %%-%ds (%%d %%s)\n", maxLen+2)

		switch m.State {
		case "completed":
			green.Printf(format, "[✓]", m.Name, m.ItemsCount, itemWord)
		case "started":
			yellow.Printf(format, "[~]", m.Name, m.ItemsCount, itemWord)
		case "locked":
			red.Printf(format, "[🔒]", m.Name, m.ItemsCount, itemWord)
		default:
			white.Printf(format, "[ ]", m.Name, m.ItemsCount, itemWord)
		}
	}
}
