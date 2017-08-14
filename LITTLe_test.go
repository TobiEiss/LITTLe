package LITTLe_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TobiEiss/LITTLe"
)

func testServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
}

func TestATestStep(t *testing.T) {
	// create test-server
	testServer := testServer()
	defer testServer.Close()

	// create TestStep
	testStep := LITTLe.TestStep{}
	request, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	testStep.ExpectedStatus = 200
	testStep.Request = request

	// run testStep
	err = testStep.Run()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestFailATestStep(t *testing.T) {
	// create test-server
	testServer := testServer()
	defer testServer.Close()

	// create TestStep
	testStep := LITTLe.TestStep{}
	request, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	testStep.ExpectedStatus = 201
	testStep.Request = request

	// run testStep
	err = testStep.Run()
	if err == nil {
		log.Println("Expected a failure")
		t.Fail()
	}
}

func TestATestCase(t *testing.T) {
	// create test-server
	testServer := testServer()
	defer testServer.Close()

	// create TestStep
	testStep := LITTLe.TestStep{}
	request, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	testStep.ExpectedStatus = 200
	testStep.Request = request

	// create TestCase
	testCase := LITTLe.TestCase{}
	testCase.Before = []LITTLe.TestUnit{testCase}
	testCase.TestUnits = []LITTLe.TestUnit{testCase}
	testCase.After = []LITTLe.TestUnit{testCase}

	// run testCase
	err = testCase.Run()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestATestSuite(t *testing.T) {
	// create test-server
	testServer := testServer()
	defer testServer.Close()

	// create TestStep
	testStep := LITTLe.TestStep{}
	request, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	testStep.ExpectedStatus = 200
	testStep.Request = request

	// create TestCase
	testCase := LITTLe.TestCase{}
	testCase.Before = []LITTLe.TestUnit{testCase}
	testCase.TestUnits = []LITTLe.TestUnit{testCase}
	testCase.After = []LITTLe.TestUnit{testCase}

	// create testSuite
	testSuite := LITTLe.TestSuite{TestCases: []LITTLe.TestCase{testCase}}

	// run testCase
	testSuite.RunTestSuite()
}
