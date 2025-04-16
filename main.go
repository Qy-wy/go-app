package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type ResultResponse struct {
	Result int `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{Error: message}
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":    err.Error(),
			"method":   req.Method,
			"endpoint": req.URL.Path,
		}).Error("Error when encoding JSON")
		return
	}
}

func writeJSONResult(w http.ResponseWriter, result int, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := ResultResponse{Result: result}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":    err.Error(),
			"method":   req.Method,
			"endpoint": req.URL.Path,
		}).Error("Error when encoding JSON")
		return
	}
}

func parseParameters(req *http.Request) (int, int, error) {
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
	a, b, err := parseParameters(req)

	if err != nil {
		sendJSONError(w, "Invalid input. Please provide valid numeric values for 'a' and 'b'.", http.StatusBadRequest, req)
		return
	}

	result := a + b

	writeJSONResult(w, result, req)
}

func Subtract(w http.ResponseWriter, req *http.Request) {
	a, b, err := parseParameters(req)

	if err != nil {
		sendJSONError(w, "Invalid input. Please provide valid numeric values for 'a' and 'b'.", http.StatusBadRequest, req)
		return
	}

	result := a - b

	writeJSONResult(w, result, req)
}

func Multiply(w http.ResponseWriter, req *http.Request) {
	a, b, err := parseParameters(req)

	if err != nil {
		sendJSONError(w, "Invalid input. Please provide valid numeric values for 'a' and 'b'.", http.StatusBadRequest, req)
		return
	}

	result := a * b

	writeJSONResult(w, result, req)
}

func Divide(w http.ResponseWriter, req *http.Request) {
	a, b, err := parseParameters(req)

	if err != nil {
		sendJSONError(w, "Invalid input. Please provide valid numeric values for 'a' and 'b'.", http.StatusBadRequest, req)
		return
	}

	if b == 0 {
		sendJSONError(w, "Division by zero is not allowed.", http.StatusBadRequest, req)
		return
	}

	result := a / b

	writeJSONResult(w, result, req)
}

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	http.HandleFunc("/sum", Add)
	http.HandleFunc("/minus", Subtract)
	http.HandleFunc("/multiply", Multiply)
	http.HandleFunc("/divide", Divide)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("Error when starting the server")
	}
}
