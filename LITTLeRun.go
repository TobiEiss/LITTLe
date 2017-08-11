package LITTLe

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

// Run is the implenetation to run a TestStep.
func (ts TestStep) Run() error {
	response, err := (&http.Client{}).Do(ts.Request)
	if err != nil {
		return err
	}

	// check response
	if response.StatusCode != ts.ExpectedStatus {
		return errors.New("ExpectedStatus: " + strconv.Itoa(ts.ExpectedStatus) + " But response is: " + strconv.Itoa(response.StatusCode))
	}

	return nil
}

// Run is the implementation to run a TestCase.
func (tc TestCase) Run() error {
	// run before
	for _, tc := range tc.Before {
		err := tc.Run()
		if err != nil {
			return err
		}
	}

	// run testUnits
	for _, tc := range tc.TestUnits {
		err := tc.Run()
		if err != nil {
			return err
		}
	}

	// run After
	for _, tc := range tc.After {
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
