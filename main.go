package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ResultResponse struct {
	Result int `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{Error: message}
	json.NewEncoder(w).Encode(errorResponse)
}

func writeJSONResult(w http.ResponseWriter, result int) {
	w.Header().Set("Content-Type", "application/json")

	response := ResultResponse{Result: result}
	json.NewEncoder(w).Encode(response)
}

func parseParametrs(req *http.Request) (int, int, error) {
	query := req.URL.Query()
	aStr := query.Get("a")
	bStr := query.Get("b")

	a, errA := strconv.Atoi(aStr)
	b, errB := strconv.Atoi(bStr)

	if errA != nil || errB != nil {
		return 0, 0, fmt.Errorf("Invalid input")
	}

	return a, b, nil
}

func Add(w http.ResponseWriter, req *http.Request) {
	a, b, err := parseParametrs(req)

	if err != nil {
		sendJSONError(w, "Invalid input. Please provide valid numeric values for 'a' and 'b'.", http.StatusBadRequest)
		return
	}

	result := a + b

	writeJSONResult(w, result)
}

func Subtract(w http.ResponseWriter, req *http.Request) {
	a, b, err := parseParametrs(req)

	if err != nil {
		sendJSONError(w, "Invalid input. Please provide valid numeric values for 'a' and 'b'.", http.StatusBadRequest)
		return
	}

	result := a - b

	writeJSONResult(w, result)
}

func Multiply(w http.ResponseWriter, req *http.Request) {
	a, b, err := parseParametrs(req)

	if err != nil {
		sendJSONError(w, "Invalid input. Please provide valid numeric values for 'a' and 'b'.", http.StatusBadRequest)
		return
	}

	result := a * b

	writeJSONResult(w, result)
}

func Divide(w http.ResponseWriter, req *http.Request) {
	a, b, err := parseParametrs(req)

	if err != nil {
		sendJSONError(w, "Invalid input. Please provide valid numeric values for 'a' and 'b'.", http.StatusBadRequest)
		return
	}

	if b == 0 {
		sendJSONError(w, "Division by zero is not allowed.", http.StatusBadRequest)
		return
	}

	result := a / b

	writeJSONResult(w, result)
}

func main() {

	http.HandleFunc("/sum", Add)
	http.HandleFunc("/minus", Subtract)
	http.HandleFunc("/multiply", Multiply)
	http.HandleFunc("/divide", Divide)

	http.ListenAndServe(":8080", nil)
}
