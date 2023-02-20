package httputils_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tinytoolkit/httputils"
)

func TestQueryString(t *testing.T) {
	r, _ := http.NewRequest("GET", "/?key=value", nil)

	if httputils.QueryString(r, "key") != "value" {
		t.Error("QueryString should return value")
	}
}

func TestQueryStrings(t *testing.T) {
	r, _ := http.NewRequest("GET", "/?key=value1,value2", nil)

	if len(httputils.QueryStrings(r, "key")) != 2 {
		t.Error("QueryStrings should return a slice of strings")
	}

	if httputils.QueryStrings(r, "key")[0] != "value1" {
		t.Error("QueryStrings should return a slice of strings")
	}

	if httputils.QueryStrings(r, "key")[1] != "value2" {
		t.Error("QueryStrings should return a slice of strings")
	}
}

func TestQueryInt(t *testing.T) {
	r, _ := http.NewRequest("GET", "/?key=1", nil)

	if httputils.QueryInt(r, "key") != 1 {
		t.Error("QueryInt should return an int")
	}
}

func TestQueryInts(t *testing.T) {
	r, _ := http.NewRequest("GET", "/?key=1,2", nil)

	if len(httputils.QueryInts(r, "key")) != 2 {
		t.Error("QueryInts should return a slice of ints")
	}

	if httputils.QueryInts(r, "key")[0] != 1 {
		t.Error("QueryInts should return a slice of ints")
	}

	if httputils.QueryInts(r, "key")[1] != 2 {
		t.Error("QueryInts should return a slice of ints")
	}
}

func TestQueryBool(t *testing.T) {
	r, _ := http.NewRequest("GET", "/?key=true", nil)

	if httputils.QueryBool(r, "key") != true {
		t.Error("QueryBool should return a bool")
	}
}

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()

	testData := map[string]string{
		"name": "John",
		"age":  "30",
	}

	httputils.WriteJSON(w, http.StatusOK, testData)

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("WriteJSON() failed to set Content-Type header")
	}

	if w.Code != http.StatusOK {
		t.Errorf("WriteJSON() failed with status code: %v", w.Code)
	}

	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf("WriteJSON() failed to decode JSON response")
	}

	if response["name"] != testData["name"] || response["age"] != testData["age"] {
		t.Errorf("WriteJSON() produced unexpected response data")
	}
}
