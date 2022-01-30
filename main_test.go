package main

import (
	"reflect"
	"testing"
)

func TestIsAnagram(t *testing.T) {

	testCases := []struct {
		input1         string
		input2         string
		expectedResult bool
		description    string
	}{
		{
			description:    "case senstive anagram strings",
			input1:         "Listen",
			input2:         "Silent",
			expectedResult: false,
		}, {
			description:    "simple anagrams",
			input1:         "triangle",
			input2:         "integral",
			expectedResult: true,
		}, {
			description:    "invalid anagrams",
			input1:         "aavvbbccd",
			input2:         "dffeed",
			expectedResult: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.description, func(t *testing.T) {
			actual := isAnagram(tC.input1, tC.input2)
			if actual != tC.expectedResult {
				t.Fatalf("Input1 %s, Input2 %s | expected %v got %v", tC.input1, tC.input2, tC.expectedResult, actual)
			}
		})
	}
}

func TestGetChracterOccurenceMapFromString(t *testing.T) {

	testCases := []struct {
		input       string
		occurence   map[string]int
		description string
	}{
		{
			description: "case 1",
			input:       "Listened",
			occurence:   map[string]int{"L": 1, "i": 1, "s": 1, "t": 1, "e": 2, "n": 1, "d": 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.description, func(t *testing.T) {
			actual := getChracterOccurenceMapFromString(tC.input)
			if !reflect.DeepEqual(actual, tC.occurence) {
				t.Fatalf("Input %s | expected %v got %v", tC.input, tC.occurence, actual)
			}
		})
	}
}
