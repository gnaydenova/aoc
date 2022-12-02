package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		part     int
		expected string
	}{
		{1, "15"},
		{2, "12"},
	}

	for _, testCase := range testCases {
		output, err := exec.Command("go", "run", "main.go", "-file", "test_input", "-part", fmt.Sprintf("%d", testCase.part)).Output()
		if err != nil {
			t.Errorf(err.Error())
		}

		if strings.Trim(string(output), "\n") != testCase.expected {
			t.Errorf("expected: %s, got: %s", testCase.expected, output)
		}
	}
}
