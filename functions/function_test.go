package functions

import (
	"fmt"
	"reflect"
	"testing"
)

func runTest(t *testing.T, testFunction interface{}, tests []struct {
	input    interface{}
	expected interface{}
}) {
	for _, test := range tests {
		result := reflect.ValueOf(testFunction).Call([]reflect.Value{reflect.ValueOf(test.input)})[0].Interface()
		fmt.Printf("Expected: %v, Got: %v\n", test.expected, result)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestCheckForNotAllowedSymbols(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{interface{}("Hello, World!\nWelcome to ASCII Art"), true},
		{interface{}("Hello, World! ©"), false},
	}
	runTest(t, CheckForNotAllowedSymbols, tests)
}

func TestSplitStringByNewline(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{"Hello, World!\nWelcome to ASCII Art", []string{"Hello, World!", "Welcome to ASCII Art"}},
		{"Hello, World!", []string{"Hello, World!"}},
		{"", []string{}},
	}
	for _, test := range tests {
		result := splitStringByNewline(test.input.(string))
		fmt.Printf("Expected: %v, Got: %v\n", test.expected, result)
		if !slicesMatch(result, test.expected.([]string)) {
			t.Errorf("For input %v, expected %v, got %v", test.input, test.expected, result)
		}
	}
}

func slicesMatch(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, element := range slice1 {
		if element != slice2[i] {
			return false
		}
	}
	return true
}
