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

	for _, m := range modules {
		itemWord := "items"
		if m.ItemsCount == 1 {
			itemWord = "item"
		}
		switch m.State {
		case "completed":
			green.Printf("[✓] %-40s (%d %s)\n", m.Name, m.ItemsCount, itemWord)
		case "started":
			yellow.Printf("[~] %-40s (%d %s)\n", m.Name, m.ItemsCount, itemWord)
		case "locked":
			red.Printf("[🔒] %-40s (%d %s)\n", m.Name, m.ItemsCount, itemWord)
		default:
			white.Printf("[ ] %-40s (%d %s)\n", m.Name, m.ItemsCount, itemWord)
		}
	}
}
