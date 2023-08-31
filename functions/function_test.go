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

func TestGetSymbolsMapVerticalRepresentation(t *testing.T) {
	test := []struct {
		input    interface{}
		expected interface{}
	}{
		{map[rune][]string{'a': {"a1", "b2", "c3"}}, map[rune][]string{'a': {"abc", "123"}}},
		{map[rune][]string{'a': {"a1", "b2", "c3"}, 'b': {"a1", "b2", "c3"}}, map[rune][]string{'a': {"abc", "123"}, 'b': {"abc", "123"}}},
	}
	runTest(t, GetSymbolsMapVerticalRepresentation, test)
}

func TestGenerateVerticalRepresentation(t *testing.T) {
	test := []struct {
		input    interface{}
		expected interface{}
	}{
		{[]string{"a1", "b2", "c3"}, []string{"abc", "123"}},
		{[]string{"a12", "b23", "c34"}, []string{"abc", "123", "234"}},
		{[]string{"a147", "b258", "c369"}, []string{"abc", "123", "456", "789"}},
	}
	runTest(t, generateVerticalRepresentation, test)
}
