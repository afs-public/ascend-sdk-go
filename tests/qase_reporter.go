package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

// QaseReporter handles reporting test results to Qase.io
type QaseReporter struct {
	enabled     bool
	apiToken    string
	apiBaseURL  string
	projectCode string
	client      *http.Client
	testRunID   int64
	mu          sync.Mutex
	results     []qaseTestResult
}

type qaseTestResult struct {
	TestName   string
	Status     string // passed, failed, skipped
	TimeMS     int64
	Comment    string
	Stacktrace string
}

var (
	qaseReporterInstance *QaseReporter
	qaseReporterOnce     sync.Once
)

// InitQaseReporter creates and initializes the singleton Qase reporter.
// Should be called from TestMain.
func InitQaseReporter() *QaseReporter {
	qaseReporterOnce.Do(func() {
		mode := os.Getenv("QASE_MODE")
		token := os.Getenv("QASE_API_TOKEN")
		project := os.Getenv("QASE_PROJECT")
		baseURL := os.Getenv("QASE_API_BASE_URL")

		if baseURL == "" {
			baseURL = "https://api.qase.io/v1"
		}
		if project == "" {
			project = "CDX"
		}

		enabled := mode == "testops" && token != ""

		qaseReporterInstance = &QaseReporter{
			enabled:     enabled,
			apiToken:    token,
			apiBaseURL:  baseURL,
			projectCode: project,
			client: &http.Client{
				Timeout: 30 * time.Second,
			},
			results: make([]qaseTestResult, 0),
		}

		if enabled {
			if err := qaseReporterInstance.initializeTestRun(); err != nil {
				fmt.Printf("Qase: failed to initialize test run: %v\n", err)
				qaseReporterInstance.enabled = false
			}
		}
	})
	return qaseReporterInstance
}

// GetQaseReporter returns the singleton Qase reporter instance.
// Returns a no-op reporter if InitQaseReporter has not been called.
func GetQaseReporter() *QaseReporter {
	if qaseReporterInstance == nil {
		return &QaseReporter{enabled: false}
	}
	return qaseReporterInstance
}

// RecordTestResult records the result of a test. Should be called via defer.
func (r *QaseReporter) RecordTestResult(t *testing.T, startTime time.Time) {
	if !r.enabled {
		return
	}

	elapsed := time.Since(startTime).Milliseconds()

	status := "passed"
	comment := ""
	if t.Skipped() {
		status = "skipped"
	} else if t.Failed() {
		status = "failed"
		comment = "Test failed"
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.results = append(r.results, qaseTestResult{
		TestName: t.Name(),
		Status:   status,
		TimeMS:   elapsed,
		Comment:  comment,
	})
}

// PublishResults uploads all recorded test results to Qase and completes the run.
func (r *QaseReporter) PublishResults() {
	if !r.enabled || r.testRunID == 0 {
		return
	}

	r.mu.Lock()
	results := make([]qaseTestResult, len(r.results))
	copy(results, r.results)
	r.mu.Unlock()

	fmt.Printf("Qase: publishing %d test results to run #%d\n", len(results), r.testRunID)

	for _, result := range results {
		caseID, err := r.getOrCreateTestCase(result.TestName)
		if err != nil {
			fmt.Printf("Qase: failed to get/create test case for %s: %v\n", result.TestName, err)
			continue
		}
		if err := r.publishSingleResult(caseID, result); err != nil {
			fmt.Printf("Qase: failed to publish result for %s: %v\n", result.TestName, err)
		}
	}

	if err := r.completeTestRun(); err != nil {
		fmt.Printf("Qase: failed to complete test run: %v\n", err)
	}

	fmt.Printf("Qase: results published - https://app.qase.io/run/%s/dashboard/%d\n", r.projectCode, r.testRunID)
}

func (r *QaseReporter) initializeTestRun() error {
	payload := map[string]interface{}{
		"title":       "Ascend SDK Go Tests",
		"description": "Automated test run for Ascend SDK Go",
		"cases":       []int{},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal test run payload: %w", err)
	}

	url := fmt.Sprintf("%s/run/%s", r.apiBaseURL, r.projectCode)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("create test run request: %w", err)
	}

	req.Header.Set("Token", r.apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("create test run: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("create test run failed, status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var result struct {
		Status bool `json:"status"`
		Result struct {
			ID int64 `json:"id"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("decode test run response: %w", err)
	}

	r.testRunID = result.Result.ID
	fmt.Printf("Qase: created test run #%d\n", r.testRunID)
	return nil
}

func (r *QaseReporter) getOrCreateTestCase(testName string) (int64, error) {
	// Search for existing test case by title
	url := fmt.Sprintf("%s/case/%s?search=%s&limit=1", r.apiBaseURL, r.projectCode, testName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Token", r.apiToken)

	resp, err := r.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var searchResult struct {
			Status bool `json:"status"`
			Result struct {
				Entities []struct {
					ID    int64  `json:"id"`
					Title string `json:"title"`
				} `json:"entities"`
			} `json:"result"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&searchResult); err == nil {
			for _, entity := range searchResult.Result.Entities {
				if entity.Title == testName {
					return entity.ID, nil
				}
			}
		}
	}

	// Create new test case
	return r.createTestCase(testName)
}

func (r *QaseReporter) createTestCase(title string) (int64, error) {
	// Derive suite name from test function name (e.g., TestAssets_... -> Assets)
	suiteName := "Ascend SDK Go"
	parts := strings.SplitN(title, "_", 2)
	if len(parts) >= 1 {
		name := strings.TrimPrefix(parts[0], "Test")
		if name != "" {
			suiteName = name
		}
	}

	payload := map[string]interface{}{
		"title":       title,
		"suite_title": suiteName,
		"automation":  2, // Automated
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("%s/case/%s", r.apiBaseURL, r.projectCode)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Token", r.apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("create test case failed, status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var result struct {
		Status bool `json:"status"`
		Result struct {
			ID int64 `json:"id"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Result.ID, nil
}

func (r *QaseReporter) publishSingleResult(caseID int64, result qaseTestResult) error {
	payload := map[string]interface{}{
		"case_id": caseID,
		"status":  result.Status,
		"time_ms": result.TimeMS,
		"comment": result.Comment,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/result/%s/%d", r.apiBaseURL, r.projectCode, r.testRunID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Token", r.apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("publish result failed, status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

func (r *QaseReporter) completeTestRun() error {
	url := fmt.Sprintf("%s/run/%s/%d/complete", r.apiBaseURL, r.projectCode, r.testRunID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Token", r.apiToken)

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Qase: test run #%d completed\n", r.testRunID)
		return nil
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("complete test run failed, status: %d, body: %s", resp.StatusCode, string(bodyBytes))
}
