package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type testCase struct {
	query    string
	want     map[string]interface{}
	endpoint string
}

func runTest(tc testCase, function http.HandlerFunc) (error, map[string]interface{}) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", tc.endpoint, tc.query), nil)
	recorder := httptest.NewRecorder()

	function(recorder, req)

	var got map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &got)

	if err != nil {
		return err, got
	}

	return err, got
}

func TestFunc(t *testing.T) {
	tests := []testCase{
		{"?a=5&b=3", map[string]interface{}{"result": 8}, "/sum"},
		{"?a=5&b=abc", map[string]interface{}{"error": "Invalid input. Please provide valid numeric values for 'a' and 'b'."}, "/sum"},
		{"?a=&b=6", map[string]interface{}{"error": "Invalid input. Please provide valid numeric values for 'a' and 'b'."}, "/sum"},
		{"?a=5&b=7", map[string]interface{}{"result": -2}, "/minus"},
		{"?a=5&b=3", map[string]interface{}{"result": 15}, "/multiply"},
		{"?a=15&b=3", map[string]interface{}{"result": 5}, "/divide"},
		{"?a=5&b=0", map[string]interface{}{"error": "Division by zero is not allowed."}, "/divide"},
	}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			var err error
			var got map[string]interface{}

			switch tc.endpoint {
			case "/sum":
				err, got = runTest(tc, Add)
			case "/minus":
				err, got = runTest(tc, Subtract)
			case "/multiply":
				err, got = runTest(tc, Multiply)
			case "/divide":
				err, got = runTest(tc, Divide)
			}

			if err != nil {
				t.Fatal("Failed to unmarshal response")
			}

			if want, ok := tc.want["result"].(int); ok {
				tc.want["result"] = float64(want)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%s response was expected for %s request, but %s was received", tc.query, tc.want, got)
			}
		})
	}
}
