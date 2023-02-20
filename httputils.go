package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// QueryString returns a string from a query like ?key=value
func QueryString(r *http.Request, key string) string {
	value := r.URL.Query().Get(key)
	if value == "" {
		return ""
	}
	return value
}

// QueryStrings returns a slice of strings from a query like ?key=value1,value2
func QueryStrings(r *http.Request, key string) []string {
	values := r.URL.Query()[key]
	if len(values) == 0 {
		return nil
	}
	return strings.Split(values[0], ",")
}

// QueryInt returns an int from a query like ?key=value
func QueryInt(r *http.Request, key string) int {
	value, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil {
		return 0
	}
	return value
}

// QueryInts returns a slice of ints from a query like ?key=value1,value2
func QueryInts(r *http.Request, key string) []int {
	values := r.URL.Query()[key]
	if len(values) == 0 {
		return nil
	}
	split := strings.Split(values[0], ",")

	ints := make([]int, len(split))
	for i, v := range split {
		ints[i], _ = strconv.Atoi(v)
	}
	return ints
}

// QueryBool returns a bool from a query like ?key=value
func QueryBool(r *http.Request, key string) bool {
	value, err := strconv.ParseBool(r.URL.Query().Get(key))
	if err != nil {
		return false
	}
	return value
}

// WriteJSON writes a JSON response to the http.ResponseWriter
func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			_, _ = fmt.Fprintf(w, `
			{
				"code": 500,
				"message": "Internal Server Error",
			}`)
		}
	}
}
