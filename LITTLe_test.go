package LITTLe_test

import (
	"encoding/json"
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
	testStep.ExpectedStatus = 200
	testStep.Request = LITTLe.Request{
		Methode: "GET",
		URL:     testServer.URL,
	}

	// run testStep
	err := testStep.Run()
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
	testStep.ExpectedStatus = 201
	testStep.Request = LITTLe.Request{
		Methode: "GET",
		URL:     testServer.URL,
	}

	// run testStep
	err := testStep.Run()
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
	testStep.ExpectedStatus = 200
	testStep.Request = LITTLe.Request{
		Methode: "GET",
		URL:     testServer.URL,
	}

	// create TestCase
	testCase := LITTLe.TestCase{}
	testCase.Before = []LITTLe.TestUnit{testCase}
	testCase.TestUnits = []LITTLe.TestUnit{testCase}
	testCase.After = []LITTLe.TestUnit{testCase}

	// run testCase
	err := testCase.Run()
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
	testStep.ExpectedStatus = 200
	testStep.Request = LITTLe.Request{
		Methode: "GET",
		URL:     testServer.URL,
	}

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

func TestMarshal(t *testing.T) {
	// create TestStep
	testStep := LITTLe.TestStep{}
	testStep.ExpectedStatus = 200
	testStep.Request = LITTLe.Request{
		Methode: "GET",
		URL:     "www.url.de",
	}

	// marshal testStep
	jsonByte, err := json.Marshal(testStep)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	if string(jsonByte) != `{"request":{"methode":"GET","url":"www.url.de"},"expectedStatus":200}` {
		t.Fail()
	}
}
