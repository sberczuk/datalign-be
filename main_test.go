package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestEvalRoute(t *testing.T) {
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		description string

		// Test input
		route string
		body  string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "evalRoute success",
			body:          "3+2",
			route:         "/eval",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "5",
		},
		{
			description:   "evalRoute with embedded newlines",
			body:          "3+\n2",
			route:         "/eval",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "5",
		},
		{
			description:   "evalRoute with embedded and trailing newlines",
			body:          "3+\n2\n",
			route:         "/eval",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "5",
		},
		{
			// this error can come from the eval or the validation.
			// TODO: Add a message
			description:   "evalRoute error",
			body:          "3+&2",
			route:         "/eval",
			expectedError: false,
			expectedCode:  400,
			expectedBody:  "0",
		},
	}

	// Setup the app as it is done in the main function
	app := NewApp()
	app.app = SetupRoutes(app.app)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		payload := struct {
			Input string
		}{Input: test.body}

		jsonPayload, err := json.Marshal(payload)
		require.NoError(t, err)
		req, _ := http.NewRequest(
			"POST",
			test.route,
			bytes.NewReader(jsonPayload),
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := io.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		if test.expectedCode == 200 {
			// only test the response if we expected real data
			answer := struct {
				Answer string `json:"answer"`
			}{}

			json.Unmarshal(body, &answer)

			// Verify, that the reponse body equals the expected body
			response, err := strconv.ParseFloat(answer.Answer, 32)
			require.NoError(t, err)
			expected, err := strconv.ParseFloat(test.expectedBody, 32)
			require.NoError(t, err)
			assert.Equalf(t, expected, response, test.description)
		}
	}
}
