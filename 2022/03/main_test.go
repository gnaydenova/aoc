package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		file     string
		part     int
		expected string
	}{
		{"test_input", 1, "157"},
		{"test_input2", 2, "70"},
	}

	for _, testCase := range testCases {
		output, err := exec.Command("go", "run", "main.go", "-file", testCase.file, "-part", fmt.Sprintf("%d", testCase.part)).Output()
		if err != nil {
			t.Errorf(err.Error())
		}

		if string(output) != testCase.expected {
			t.Errorf("expected: %s, got: %s", testCase.expected, output)
		}
	}
}
