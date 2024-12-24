package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	expected := map[string][]string{
		"акптя":  {"пятак", "пятка", "тяпка"},
		"иклост": {"листок", "слиток", "столик"},
	}
	anagramSets := findAnagramSets(&words)

	if !reflect.DeepEqual(*anagramSets, expected) {
		t.Errorf("Expected %v, but got %v", expected, *anagramSets)
	}
}

func TestSortString(t *testing.T) {
	word := "пятка"
	expected := "акптя"
	result := sortString(word)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSortStringEmpty(t *testing.T) {
	word := ""
	expected := ""
	result := sortString(word)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSortStringSingleCharacter(t *testing.T) {
	word := "а"
	expected := "а"
	result := sortString(word)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}