package LITTLe

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Run is the implenetation to run a TestStep.
func (ts TestStep) Run() error {
	// build request
	bodyByte, err := json.Marshal(ts.Request.Body)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(ts.Request.Methode, ts.Request.URL, bytes.NewReader(bodyByte))
	if err != nil {
		return err
	}

	// do request
	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return err
	}

	// check response
	if response.StatusCode != ts.ExpectedStatus {
		return ReportError{ActualStatusCode: response.StatusCode, TestStep: &ts}
	}

	return nil
}

// Run is the implementation to run a TestCase.
func (tc TestCase) Run() error {
	for _, tc := range append(tc.Before, append(tc.TestUnits, tc.After...)...) {
		err := tc.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// RunTestSuite runs a TestSuite and create a report
func (ts TestSuite) RunTestSuite() {
	var err error
	for _, testCase := range ts.TestCases {
		err = testCase.Run()
		if err != nil {
			break
		}
	}
	log.Println(err)
}
