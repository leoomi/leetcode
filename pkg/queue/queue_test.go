package queueu

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		values []int
	}{
		{values: []int{0}},
		{values: []int{1, 2}},
		{values: []int{3, 4, 5, 6}},
	}

	for _, testCase := range tests {
		queue := Queue[int]{}

		for _, val := range testCase.values {
			queue.Add(val)
		}

		if len(queue.nodes) != len(testCase.values) {
			t.Errorf("Expected queue with length %v but got %v", len(testCase.values), len(queue.nodes))
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		values []int
		result bool
	}{
		{values: []int{}, result: true},
		{values: []int{0}, result: false},
		{values: []int{1, 2}, result: false},
		{values: []int{3, 4, 5, 6}, result: false},
	}

	for _, testCase := range tests {
		queue := Queue[int]{}

		for _, val := range testCase.values {
			queue.Add(val)
		}

		if queue.IsEmpty() != testCase.result {
			t.Errorf("Expected IsEmpty() to return %v but got %v", testCase.result, queue.IsEmpty())
		}
	}
}

func TestRemoveFilledQueues(t *testing.T) {
	var tests = []struct {
		values []int
	}{
		{values: []int{0}},
		{values: []int{1, 2}},
		{values: []int{3, 4, 5, 6}},
	}

	for _, testCase := range tests {
		queue := Queue[int]{}

		for _, val := range testCase.values {
			queue.Add(val)
		}

		for _, expectedVal := range testCase.values {
			val, _ := queue.Remove()

			if val != expectedVal {
				t.Errorf("Expected to remove value %v but got %v", expectedVal, val)
			}
		}

		if !queue.IsEmpty() {
			t.Errorf("Expected queue to be empty")
		}
	}
}

func TestRemoveEmptyQueue(t *testing.T) {
	queue := Queue[int]{}

	val, err := queue.Remove()

	if err == nil {
		t.Errorf("Expected error to be returned when queue is empty")
	}

	if val != 0 {
		t.Errorf("Expected zero value to be returned when an error happens")
	}
}

func TestAddRemove(t *testing.T) {
	queue := Queue[int]{}

	queue.Add(1)
	queue.Add(2)
	queue.Remove()
	queue.Add(3)

}
