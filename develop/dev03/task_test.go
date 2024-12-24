package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestReadLines(t *testing.T) {
	data := "line1\nline2\nline3\n"
	expected := []string{"line1", "line2", "line3"}

	r := bytes.NewBufferString(data)
	lines, err := readLines(r)
	if err != nil {
		t.Errorf("readLines() returned error: %v", err)
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("readLines() returned %v, expected %v", lines, expected)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	lines := []string{"line1", "line2", "line1", "line3", "line2"}
	expected := []string{"line1", "line2", "line3"}

	result := removeDuplicates(lines)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeDuplicates() returned %v, expected %v", result, expected)
	}
}

func TestExtractNumericSuffix(t *testing.T) {
	s := "abc123def456"
	expectedVal := "abc123def"
	expectedSuffix := "456"

	resultVal, resultSuffix := extractNumericSuffix(s)
	if resultVal != expectedVal || resultSuffix != expectedSuffix {
		t.Errorf("extractNumericSuffix() returned (%s, %s), expected (%s, %s)", resultVal, resultSuffix, expectedVal, expectedSuffix)
	}
}

func TestIsNumeric(t *testing.T) {
	c := byte('9')

	result := isNumeric(c)
	if result != true {
		t.Errorf("isNumeric() returned %v, expected %v", result, true)
	}
}

func TestParseMonth(t *testing.T) {
	s := "January"
	expected, _ := time.Parse("January", s)

	result, err := parseMonth(s)
	if err != nil {
		t.Errorf("parseMonth() returned error: %v", err)
	}

	if result != expected {
		t.Errorf("parseMonth() returned %v, expected %v", result, expected)
	}
}

func TestIsSorted(t *testing.T) {
	lines := []string{"line1", "line2", "line3"}
	lessFunc := func(i, j int) bool {
		return lines[i] < lines[j]
	}

	result := isSorted(lines, lessFunc)
	if result != true {
		t.Errorf("isSorted() returned %v, expected %v", result, true)
	}
}