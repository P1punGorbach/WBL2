package main

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	testTable := []struct {
		input    string
		expected string
		err      error
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde", err: nil},
		{input: "abcd", expected: "abcd", err: nil},
		{input: "45", expected: "", err: errors.New("некорректная строка")},
		{input: "", expected: "", err: nil},
		{input: `qwe\4\5`, expected: `qwe45`, err: nil},
		{input: `qwe\45`, expected: `qwe44444`, err: nil},
		{input: `qwe\\5`, expected: `qwe\\\\\`, err: nil},
	}

	for _, testCase := range testTable {
		result, err := Unpack(testCase.input)

		if result != testCase.expected {
			t.Errorf("Expected string: %s, got: %s", testCase.expected, result)
		}

		if err != nil || testCase.err != nil {
			if testCase.err == nil || err == nil {
				t.Error("несоответствие ошибок:", err, testCase.err)
			} else if testCase.err.Error() != err.Error() {
				t.Error("несоответствие ошибок:", testCase.err, err.Error())
			}
		}
	}
}
