package main

import (
	"github.com/kyokomi/emoji"
	"golang-united-school-homework-1/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	if solution.GetMessage() != emoji.Sprint("Hello :world_map:!") {
		t.Errorf("Wrong String")
	}
}
