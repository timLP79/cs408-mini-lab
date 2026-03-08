package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func progressBar(completed, total int) string {
	if total == 0 {
		return "[]"
	}
	bar := "["
	for i := 0; i < total; i++ {
		if i < completed {
			bar += "█"
		} else {
			bar += "-"
		}
	}
	bar += "]"
	return bar
}

func displayModules(courseName string, modules []Module, completedCounts map[int]int, trackableCounts map[int]int) {
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)

	fmt.Printf("\n%s\n", courseName)
	fmt.Println(strings.Repeat("-", len(courseName)))

	maxLen := 0
	for _, m := range modules {
		if len(m.Name) > maxLen {
			maxLen = len(m.Name)
		}
	}

	for _, m := range modules {
		// Only show a progress bar if there are trackable items, or if the module
		// isn't marked completed (avoids showing empty bars for non-trackable modules).
		var progress string
		if completedCounts[m.ID] > 0 || m.State != "completed" {
			progress = fmt.Sprintf("%s %d/%d", progressBar(completedCounts[m.ID], trackableCounts[m.ID]), completedCounts[m.ID], trackableCounts[m.ID])
		}
		format := fmt.Sprintf("%%s %%-%ds %%s\n", maxLen+2)

		switch m.State {
		case "completed":

			if completedCounts[m.ID] > 0 {
				green.Printf(format, "[✓]", m.Name, progress)
			} else {
				white.Printf(format, "[ ]", m.Name, progress)
			}
		case "started":
			yellow.Printf(format, "[~]", m.Name, progress)
		case "locked":
			red.Printf(format, "[🔒]", m.Name, progress)
		default:
			white.Printf(format, "[ ]", m.Name, progress)
		}
	}
}
