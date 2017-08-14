package LITTLe

import (
	"fmt"
)

// TestUnit can be a TestStep or a TestCase. Both is runable.
type TestUnit interface {
	Run() error
}

// TestStep is the smallest unit of a test. It's just a little test step.
// It represent one request. Reuse this in all your TestCases.
type TestStep struct {
	Request        Request `json:"request,omitempty"`
	Title          string  `json:"title,omitempty"`
	Description    string  `json:"description,omitempty"`
	ExpectedStatus int     `json:"expectedStatus"`
}

// TestCase hold a batch of TestCases
// It represent a test-case. Reuse this in all your TestSuites
// Use "Before" and "After" to
type TestCase struct {
	Before      []TestUnit
	TestUnits   []TestUnit
	After       []TestUnit
	Title       string
	Description string
}

// TestSuite represent a batch of TestCases
type TestSuite struct {
	TestCases   []TestCase
	Title       string
	Description string
}

// ReportError ths error is used for an TestUnit-Fail
type ReportError struct {
	TestStep         *TestStep
	ActualStatusCode int
}

type Request struct {
	Methode string            `json:"methode,omitemtpy"`
	URL     string            `json:"url,omitempty"`
	Body    interface{}       `json:"body,omitempty"`
	Header  map[string]string `json:"header,omitempty"`
}

func (e ReportError) Error() string {
	return fmt.Sprintf("TestCase '%s' failed - Expected StatusCode: '%d' Actual StatusCode: '%d'", e.TestStep.Title, e.TestStep.ExpectedStatus, e.ActualStatusCode)
}
