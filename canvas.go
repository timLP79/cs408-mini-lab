package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func fetchCourses(token, baseURL string) ([]Course, error) {
	url := baseURL + "/api/v1/courses?per_page=100&enrollment_state=active"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var courses []Course
	err = json.NewDecoder(resp.Body).Decode(&courses)
	return courses, err
}

func fetchModules(token, baseURL string, courseID int) ([]Module, error) {
	url := fmt.Sprintf("%s/api/v1/courses/%d/modules?per_page=100", baseURL, courseID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var modules []Module
	err = json.NewDecoder(resp.Body).Decode(&modules)
	return modules, err
}

/*func getNextPage(linkHeader string) string {

}
*/
