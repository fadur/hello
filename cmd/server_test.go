package cmd_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
	"github.com/fadur/hello/internal/handlers"
)

func TestIndexHandler(t *testing.T) {
	// Create a new godog suite
	suite := godog.TestSuite{
		Name:                "IndexHandler",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "progress",
		},
	}

	// Run the godog suite
	status := suite.Run()

	// Check the status and fail the test if there are any failures
	if status != 0 {
		t.Fail()
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	// Define the steps for the scenario

	// Step 2: Send a GET request to the index endpoint
	ctx.Step(`^I send a GET request to the index endpoint$`, func() error {
		req, err := http.NewRequest("GET", "/index", nil)
		if err != nil {
			return err
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.IndexHandler)

		handler.ServeHTTP(rr, req)

		return nil
	})

	// Step 3: Verify the response status code
	rr := httptest.NewRecorder()
	ctx.Step(`^the response status code should be (\d+)$`, func(statusCode int) error {
		if rr.Code != statusCode {
			return fmt.Errorf("expected status code %d, but got %d", statusCode, rr.Code)
		}
		return nil
	})

}
