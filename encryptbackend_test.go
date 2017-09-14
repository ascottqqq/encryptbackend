package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBackEnd(t *testing.T) {
	type testCase struct {
		method string
		result int
		body   string
	}
	testCases := [...]testCase{{http.MethodGet, http.StatusMethodNotAllowed, ""},
		{http.MethodHead, http.StatusMethodNotAllowed, ""},
		{http.MethodPut, http.StatusMethodNotAllowed, ""},
		{http.MethodPatch, http.StatusMethodNotAllowed, ""},
		{http.MethodDelete, http.StatusMethodNotAllowed, ""},
		{http.MethodConnect, http.StatusMethodNotAllowed, ""},
		{http.MethodOptions, http.StatusMethodNotAllowed, ""},
		{http.MethodTrace, http.StatusMethodNotAllowed, ""},
		{http.MethodPost, http.StatusBadRequest, ""},
		{http.MethodPost, http.StatusBadRequest, `[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`},
		{http.MethodPost, http.StatusBadRequest,
			`{"key": [18,52], "counter": [0], "nonce": [67,33], "plaintext": "SGVsbG8="}`},
		{http.MethodPost, http.StatusBadRequest,
			`{"key": [18,52], "counter": 0, "nonce": [01,02,03,04,05,06,07,08,09,10,11,12,13], 
      "plaintext": "SGVsbG8="}`},
		{http.MethodPost, http.StatusBadRequest,
			`{"key": [18,52], "counter": 0, "nonce": [67,33]`},
		{http.MethodPost, http.StatusOK,
			`{"key": [18,52], "counter": 0, "nonce": [67,33], "plaintext": "SGVsbG8="}`}}
	for _, test := range testCases {
		w := httptest.NewRecorder()
		request := httptest.NewRequest(test.method, "http://localhost",
			strings.NewReader(test.body))
		EncryptHandler(w, request)
		result := w.Result()
		if result.StatusCode != test.result {
			t.Errorf("%s method has %d status code in result not %d", test,
				result.StatusCode, http.StatusMethodNotAllowed)
		}
	}
}
