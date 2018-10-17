package main

import (
	"testing"
)

func TestGetPassScores(t *testing.T) {
	outputs := getPassScores(scores, pass)
	for _, output := range outputs {
		if output.Score <= 59 {
			t.Error("TestGetPassScores failed")
		}
	}
}
func TestGetLengthName(t *testing.T) {
	outputs := getLengthName(scores, length)
	for _, output := range outputs {
		if len(output.Name) != 11 {
			t.Error("TestgetLengthName failed")
		}
	}
}

func TestGetSpecificClass(t *testing.T) {
	outputs := getSpecificClass(scores, certainClass)
	for _, output := range outputs {
		if output.Class != "C1" {
			t.Error("TestgetSpecificClass failed")
		}
	}
}
