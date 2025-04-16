package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testCase struct {
	query string
	want  string // TODO: Use map[string]interface{}
} // may be add here endpoint string

func TestAdd(t *testing.T) {
	tests := []testCase{
		{"?a=5&b=3", `{"result":8}`},
		{"?a=5&b=abc", `{"error":"Invalid input. Please provide valid numeric values for 'a' and 'b'."}`},
		{"?a=&b=6", `{"error":"Invalid input. Please provide valid numeric values for 'a' and 'b'."}`},
		{"?a=5&b=7", `{"result":12}`},
	}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			// TODO: Extract this code into a separate func and invoke it with endpoint parameters.
			req, _ := http.NewRequest("GET", fmt.Sprintf("/sum%s", tc.query), nil) // TODO: Add err handling
			recorder := httptest.NewRecorder()

			Add(recorder, req)

			var got, want any
			json.Unmarshal(recorder.Body.Bytes(), &got) // TODO: Add err handling
			json.Unmarshal([]byte(tc.want), &want)      // TODO: Add err handling

			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) { // TODO: try to use another func to qualing
				t.Errorf("%s response was expected for %s request, but %s was received", tc.query, tc.want, recorder.Body.String())
			}
		})
	}
}
func TestSubtract(t *testing.T) {
	tests := []testCase{
		{"?a=5&b=3", `{"result":2}`},
		{"?a=5&b=abc", `{"error":"Invalid input. Please provide valid numeric values for 'a' and 'b'."}`},
		{"?a=&b=6", `{"error":"Invalid input. Please provide valid numeric values for 'a' and 'b'."}`},
		{"?a=5&b=7", `{"result":-2}`},
	}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/sum%s", tc.query), nil) // /sum?
			recorder := httptest.NewRecorder()

			Subtract(recorder, req)

			var got, want any
			json.Unmarshal(recorder.Body.Bytes(), &got)
			json.Unmarshal([]byte(tc.want), &want)

			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
				t.Errorf("%s response was expected for %s request, but %s was received", tc.query, tc.want, recorder.Body.String())
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []testCase{
		{"?a=5&b=3", `{"result":15}`},
		{"?a=5&b=abc", `{"error":"Invalid input. Please provide valid numeric values for 'a' and 'b'."}`},
		{"?a=&b=6", `{"error":"Invalid input. Please provide valid numeric values for 'a' and 'b'."}`},
		{"?a=5&b=7", `{"result":35}`},
	}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/sum%s", tc.query), nil)
			recorder := httptest.NewRecorder()

			Multiply(recorder, req)

			var got, want any
			json.Unmarshal(recorder.Body.Bytes(), &got)
			json.Unmarshal([]byte(tc.want), &want)

			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
				t.Errorf("%s response was expected for %s request, but %s was received", tc.query, tc.want, recorder.Body.String())
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []testCase{
		{"?a=15&b=3", `{"result":5}`},
		{"?a=5&b=0", `{"error":"Division by zero is not allowed."}`},
		{"?a=&b=6", `{"error":"Invalid input. Please provide valid numeric values for 'a' and 'b'."}`},
		{"?a=7&b=7", `{"result":1}`},
	}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/sum%s", tc.query), nil)
			recorder := httptest.NewRecorder()

			Divide(recorder, req)

			var got, want any
			json.Unmarshal(recorder.Body.Bytes(), &got)
			json.Unmarshal([]byte(tc.want), &want)

			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
				t.Errorf("%s response was expected for %s request, but %s was received", tc.query, tc.want, recorder.Body.String())
			}
		})
	}
}
