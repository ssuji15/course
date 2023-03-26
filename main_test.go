package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCourseOrder(t *testing.T) {
	tests := []struct {
		name          string
		numOfCourse   int
		prerequisite  [][]int
		expectedOrder [][]int
		expectedErr   error
	}{
		{
			name:          "Case 1",
			numOfCourse:   2,
			prerequisite:  [][]int{{1, 0}},
			expectedOrder: [][]int{{0, 1}},
			expectedErr:   nil,
		},
		{
			name:          "Case 2",
			numOfCourse:   4,
			prerequisite:  [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expectedOrder: [][]int{{0, 1, 2, 3}, {0, 2, 1, 3}},
			expectedErr:   nil,
		},
		{
			name:          "Case 3",
			numOfCourse:   1,
			prerequisite:  [][]int{},
			expectedOrder: [][]int{{0}},
			expectedErr:   nil,
		},
		{
			name:          "Case 4",
			numOfCourse:   3,
			prerequisite:  [][]int{},
			expectedOrder: [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}},
			expectedErr:   nil,
		},
		{
			name:          "Case 5",
			numOfCourse:   3,
			prerequisite:  [][]int{{1, 0}},
			expectedOrder: [][]int{{0, 1, 2}, {0, 2, 1}, {2, 0, 1}},
			expectedErr:   nil,
		},
		{
			name:          "cycle in prerequisites",
			numOfCourse:   4,
			prerequisite:  [][]int{{1, 0}, {2, 1}, {3, 2}, {0, 3}},
			expectedOrder: [][]int{{}},
			expectedErr:   fmt.Errorf("Cannot complete all courses!"),
		},
		{
			name:          "invalid course index in prerequisites",
			numOfCourse:   4,
			prerequisite:  [][]int{{1, 5}, {2, -1}},
			expectedOrder: [][]int{{}},
			expectedErr:   fmt.Errorf("invalid prerequisite course index 5"),
		},
		{
			name:          "invalid number of courses",
			numOfCourse:   -1,
			prerequisite:  [][]int{{1, 0}},
			expectedOrder: [][]int{{}},
			expectedErr:   fmt.Errorf("Invalid number of courses"),
		},
		{
			name:          "invalid number of courses",
			numOfCourse:   2001,
			prerequisite:  [][]int{{1, 0}},
			expectedOrder: [][]int{{}},
			expectedErr:   fmt.Errorf("Invalid number of courses"),
		},
		{
			name:          "invalid prerequisite length",
			numOfCourse:   2,
			prerequisite:  [][]int{{1, 0, 2}},
			expectedOrder: [][]int{{}},
			expectedErr:   fmt.Errorf("prerequisite length must be 2"),
		},
		{
			name:          "mapping prerequisite to itself",
			numOfCourse:   2,
			prerequisite:  [][]int{{1, 1}},
			expectedOrder: [][]int{{}},
			expectedErr:   fmt.Errorf("prerequisite course cannot map to itself"),
		},
		{
			name:          "too many prerequisites",
			numOfCourse:   2,
			prerequisite:  [][]int{{1, 0}, {2, 0}, {2, 1}, {2, 0}},
			expectedOrder: [][]int{{}},
			expectedErr:   fmt.Errorf("number of prerequisites exceeds the maximum possible number"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, err := courseOrder(tt.numOfCourse, tt.prerequisite)

			if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("expected error %v, but got %v", tt.expectedErr, err)
			}

			if !isValidOrder(order, tt.expectedOrder) {
				t.Errorf("expected order %v, but got %v", tt.expectedOrder, order)
			}
		})
	}
}

func isValidOrder(order []int, validOrders [][]int) bool {
	for _, validOrder := range validOrders {
		match := true
		for i := 0; i < len(order); i++ {
			if order[i] != validOrder[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
