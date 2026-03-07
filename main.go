package main

import (
	"fmt"
	"log"
	"os"

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

	courses, err := fetchCourses(token, baseURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, course := range courses {
		fmt.Printf("%d: %s (%s)\n", course.ID, course.Name, course.CourseCode)
	}
}
