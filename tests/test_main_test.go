package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	reporter := InitQaseReporter()
	exitCode := m.Run()
	if reporter.enabled && reporter.apiToken != "" {
		reporter.PublishResults()
	}
	os.Exit(exitCode)
}
