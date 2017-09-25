package services

import "testing"

func TestRunOldStyle(t *testing.T) {
	testSt := "test string"
	someInt := 5

	expectedSt := "test string"
	expectedInt := 5

	someService := SomeService{}

	returnString, returnInt := someService.Run(testSt, someInt)
	if returnString != expectedSt || returnInt != expectedInt {
		t.Errorf("expected '%s', '%v' but got '%s', '%v'", expectedSt, expectedInt, returnString, returnInt)
	}
}

func TestRunNewStyle(t *testing.T) {
	testCases := []struct {
		name       string
		st         string
		in         int
		expectedSt string
		expectedIn int
	}{
		{
			"Return back exactly what was passed in if in is not 4",
			"test string",
			5,
			"test string",
			5,
		},
	}

	someService := SomeService{}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualSt, actualIn := someService.Run(testCase.st, testCase.in)
			if actualSt != testCase.expectedSt || actualIn != testCase.expectedIn {
				t.Errorf("expected '%s', '%v' but got '%s', '%v'", testCase.expectedSt, testCase.expectedIn, actualSt, actualIn)
			}
		})
	}
}
