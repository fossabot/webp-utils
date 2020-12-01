package config

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	testCases := []struct {
		file           string
		expectedError  string
		expectedConfig map[string]interface{}
	}{
		{
			file:          "basic_valid.json",
			expectedError: "",
			expectedConfig: map[string]interface{}{
				"o": "{{ .source_file_name }}.webp",
			},
		},
		{
			file:          "syntactic_invalid.json",
			expectedError: "error parsing JSON bytes: invalid character '\\n' in string literal",
			expectedConfig: nil,
		},
		{
			file:          "basic_invalid.json",
			expectedError: "/: {\"quiet\":true} \"o\" value is required",
			expectedConfig: nil,
		},
		{
			file:          "non_existent.json",
			expectedError: "open testdata/non_existent.json: no such file or directory",
			expectedConfig: nil,
		},
	}

	for _, testCase := range testCases {
		config, err := Load("testdata/" + testCase.file)

		if testCase.expectedError == "" && err != nil {
			t.Error("Expected no error, but got ", err.Error())
		} else if testCase.expectedError != "" && (err == nil || testCase.expectedError != err.Error()) {
			t.Errorf("Error matches dont match expected %s, but got %v", testCase.expectedError, err)
		}

		if ! reflect.DeepEqual(testCase.expectedConfig, config) {
			t.Error("Configs dont match")
		}
	}
}
