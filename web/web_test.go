package web

import (
	"testing"
)

func TestToInt(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected int
	}{
		{int(9), 9},
		{int64(121), 121},
		{string("23"), 23},
		{float64(99.4), 99},
	}

	for _, testCase := range testCases {
		answer := ToInt(testCase.value)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.value, testCase.expected, answer)
		}
	}
}

func TestToString(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected string
	}{
		{int(9), "9"},
		{int64(121), "121"},
		{string("23"), "23"},
		{float64(99.4), "99.4"},
	}

	for _, testCase := range testCases {
		answer := ToString(testCase.value)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.value, testCase.expected, answer)
		}
	}
}

func TestToFloat64(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected float64
	}{
		{int(9), 9},
		{int64(121), 121},
		{string("23"), 23},
		{float64(99.4), 99.4},
	}

	for _, testCase := range testCases {
		answer := ToFloat64(testCase.value)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.value, testCase.expected, answer)
		}
	}
}
