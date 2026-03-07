package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Enrollment struct {
	Type            string `json:"type"`
	EnrollmentState string `json:"enrollment_state"`
}
type Course struct {
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	CourseCode    string       `json:"course_code"`
	WorkFlowState string       `json:"workflow_state"`
	Enrollments   []Enrollment `json:"enrollments"`
}

type Module struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Position    int    `json:"position"`
	State       string `json:"state"`
	CompletedAt string `json:"completed_at"`
	ItemsCount  int    `json:"items_count"`
}

func fetchPage(token, url string, target interface{}) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return "", err
	}

	return getNextPage(resp.Header.Get("Link")), nil
}

func getNextPage(linkHeader string) string {
	parts := strings.Split(linkHeader, ",")
	for _, part := range parts {
		segments := strings.Split(strings.TrimSpace(part), ";")
		if len(segments) == 2 && strings.TrimSpace(segments[1]) == `rel="next"` {
			url := strings.TrimSpace(segments[0])
			url = strings.Trim(url, "<>")
			return url
		}
	}
	return ""
}

func fetchCourses(token, baseURL string) ([]Course, error) {
	url := baseURL + "/api/v1/courses?per_page=100&enrollment_state=active"
	var allCourses []Course

	for url != "" {
		var page []Course
		nextURL, err := fetchPage(token, url, &page)
		if err != nil {
			return nil, err
		}
		allCourses = append(allCourses, page...)
		url = nextURL
	}

	return allCourses, nil
}

func fetchModules(token, baseURL string, courseID int) ([]Module, error) {
	url := fmt.Sprintf("%s/api/v1/courses/%d/modules?per_page=100", baseURL, courseID)
	var allModules []Module

	for url != "" {
		var page []Module
		nextURL, err := fetchPage(token, url, &page)
		if err != nil {
			return nil, err
		}
		allModules = append(allModules, page...)
		url = nextURL
	}

	return allModules, nil
}
