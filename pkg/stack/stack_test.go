package stack

import "testing"

type testCase struct {
	array []int
	size  int
}

func TestStack(t *testing.T) {
	testCases := []testCase{
		{
			array: []int{1, 2, 3},
			size:  3,
		},
		{
			array: []int{1},
			size:  1,
		},
		{
			array: []int{},
			size:  0,
		},
	}

	for _, test := range testCases {
		stack := NewLinkedListStack[int]()

		runAddTestCase(t, stack, test)
	}

	for _, test := range testCases {
		stack := NewArrayStack[int](len(test.array))

		runAddTestCase(t, stack, test)
	}
}

func runAddTestCase(t *testing.T, stack Stack[int], test testCase) {
	for _, num := range test.array {
		stack.Add(num)
	}

	if stack.Size() != test.size {
		t.Errorf("Expected size %d, but got %d", test.size, stack.Size())
	}

	for i := len(test.array) - 1; i >= 0; i-- {
		val := stack.Pop()

		if val != test.array[i] {
			t.Errorf("Expected value %d, but got %d", test.array[i], val)
		}
	}

	if stack.Size() != 0 {
		t.Errorf("Expected size %d, but got %d", 0, stack.Size())
	}
}
