package main

import "testing"

func TestTargetAreaHit(t *testing.T) {
	testCases := []struct {
		inputArea  Area
		inputPoint Point
		expected   bool
	}{
		{
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{28, -7},
			expected:   true,
		},
		{ // exact borders
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{20, -10},
			expected:   true,
		},
		{ // exact borders
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{30, -5},
			expected:   true,
		},
		{ // fail by +X
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{31, -7},
			expected:   false,
		},
		{ // fail by -X
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{19, -7},
			expected:   false,
		},
		{ // fail by +Y
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{28, -4},
			expected:   false,
		},
		{ // fail by -Y
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{28, -11},
			expected:   false,
		},
	}
	for _, tc := range testCases {
		t.Run("Target hit test", func(t *testing.T) {
			if targetAreaHit(tc.inputArea, tc.inputPoint) != tc.expected {
				t.Error("input not matching expected")
			}
		})
	}
}

func TestInsideBoundaries(t *testing.T) {
	testCases := []struct {
		inputArea  Area
		inputPoint Point
		expected   bool
	}{
		{ // random point inside
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{0, 0},
			expected:   true,
		},
		{ // another random point inside with positives
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{15, 15},
			expected:   true,
		},
		{ // another random point inside with negatives
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{22, -2},
			expected:   true,
		},
		{ // exact borders
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{30, -10},
			expected:   true,
		},
		{ // fail by +X going too right
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{31, -10},
			expected:   false,
		},
		{ // fail by +X going too right
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{100, -5},
			expected:   false,
		},
		{ // fail by +Y going too down
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{22, -11},
			expected:   false,
		},
		{ // fail by +Y going too down
			inputArea:  Area{20, 30, -10, -5},
			inputPoint: Point{2, -222},
			expected:   false,
		},
	}
	for _, tc := range testCases {
		t.Run("Target hit test", func(t *testing.T) {
			if insideBoundaries(tc.inputArea, tc.inputPoint) != tc.expected {
				t.Error("input not matching expected")
			}
		})
	}
}

func TestStep(t *testing.T) {
	testCases := []struct {
		input    Point
		xVel     int
		yVel     int
		expected Point
		expX     int
		expY     int
	}{
		{ // 0,0 > 7, 2
			input:    Point{0, 0},
			xVel:     7,
			yVel:     2,
			expected: Point{7, 2},
			expX:     6,
			expY:     1,
		},
		{ // 7,2 > 13, 3
			input:    Point{7, 2},
			xVel:     6,
			yVel:     1,
			expected: Point{13, 3},
			expX:     5,
			expY:     0,
		},
		{ // 13, 3 >
			input:    Point{13, 3},
			xVel:     5,
			yVel:     0,
			expected: Point{18, 3},
			expX:     4,
			expY:     -1,
		},
		{
			input:    Point{18, 3},
			xVel:     4,
			yVel:     -1,
			expected: Point{22, 2},
			expX:     3,
			expY:     -2,
		},
		{
			input:    Point{22, 2},
			xVel:     3,
			yVel:     -2,
			expected: Point{25, 0},
			expX:     2,
			expY:     -3,
		},
		{
			input:    Point{25, 0},
			xVel:     2,
			yVel:     -3,
			expected: Point{27, -3},
			expX:     1,
			expY:     -4,
		},
		{
			input:    Point{27, -3},
			xVel:     1,
			yVel:     -4,
			expected: Point{28, -7},
			expX:     0,
			expY:     -5,
		},
		{
			input:    Point{28, -7},
			xVel:     0,
			yVel:     -5,
			expected: Point{28, -12},
			expX:     0,
			expY:     -6,
		},
	}
	for _, tc := range testCases {
		t.Run("Target hit test", func(t *testing.T) {
			newPoint, newX, newY := step(tc.input, tc.xVel, tc.yVel)
			if newPoint != tc.expected {
				t.Errorf("expected: %v got %v\n", tc.expected, newPoint)
			}
			if newX != tc.expX {
				t.Errorf("expected: %d got %d\n", tc.expX, newX)
			}
			if newY != tc.expY {
				t.Errorf("expected: %d got %d\n", tc.expY, newY)
			}
		})
	}
}
