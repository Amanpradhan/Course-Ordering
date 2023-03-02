package main

import (
	"testing"
)

func TestFindOrder(t *testing.T) {
	// Test case 1: valid course order.
	n1 := 4
	prereq1 := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	expectedOrder1 := []int{0, 1, 2, 3}
	actualOrder1 := findOrder(n1, prereq1)
	isCorrect1 := len(actualOrder1) == len(actualOrder1)
	if isCorrect1 {
		for i, v := range expectedOrder1 {
			if v != actualOrder1[i] {
				isCorrect1 = false
				break
			}
		}
	}
	if !isCorrect1 {
		t.Errorf("Test case 1 failed: expected %v but got %v", expectedOrder1, actualOrder1)
	}

	// Test case 2: No valid course order.
	n2 := 2
	prereq2 := [][]int{{1, 0}}
	expectedOrder2 := []int{0, 1}
	actualOrder2 := findOrder(n2, prereq2)
	isCorrect2 := len(actualOrder2) == len(actualOrder2)
	if isCorrect2 {
		for i, v := range expectedOrder2 {
			if v != actualOrder2[i] {
				isCorrect2 = false
				break
			}
		}
	}
	if !isCorrect2 {
		t.Errorf("Test case 2 failed: expected %v but got %v", expectedOrder2, actualOrder2)
	}

	// Test case 3: empty
	n3 := 1
	prereq3 := [][]int{}
	expectedOrder3 := []int{}
	actualOrder3 := findOrder(n3, prereq3)
	isCorrect3 := len(actualOrder3) == len(actualOrder3)
	if isCorrect3 {
		for i, v := range expectedOrder3 {
			if v != actualOrder3[i] {
				isCorrect3 = false
				break
			}
		}
	}
	if !isCorrect3 {
		t.Errorf("Test case 3 failed: expected %v but got %v", expectedOrder3, actualOrder3)
	}
}
