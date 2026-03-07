package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: .env file not found")
	}

	token := os.Getenv("CANVAS_API_TOKEN")
	baseURL := os.Getenv("CANVAS_BASE_URL")

	if token == "" {
		log.Fatal("Error: CANVAS_API_TOKEN is not set")
	}

	// Fetch display courses
	courses, err := fetchCourses(token, baseURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nYour Courses:")
	fmt.Println("-------------------------------")
	for i, course := range courses {
		fmt.Printf("%d: %s\n", i+1, course.Name)
	}

	// Get user input
	fmt.Print("\nEnter course number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(courses) {
		log.Fatal("Error: invalid course material")
	}

	selected := courses[choice-1]

	// Fetch modules
	modules, err := fetchModules(token, baseURL, selected.ID)
	if err != nil {
		log.Fatal(err)
	}

	displayModules(selected.Name, modules)
}
