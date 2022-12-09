package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		part     int
		expected string
		file     string
	}{
		{1, "13", "test_input"},
		{2, "36", "test_input2"},
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
