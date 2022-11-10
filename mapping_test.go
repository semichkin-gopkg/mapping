package mapping

import (
	"fmt"
	"testing"
)

var testCases = map[int]string{
	0: "0",
	1: "1",
	2: "2",
}

func TestMapping_ToRight(t *testing.T) {
	defaultRight := "default"

	mapping := New(testCases, WithDefaultRight[int, string](defaultRight))

	for left, right := range testCases {
		if got := mapping.ToRight(left); got != right {
			t.Error(
				fmt.Sprintf(
					"Method ToRight with argument %v must return %v, got %v",
					left,
					right,
					got,
				),
			)
		}
	}

	if got := mapping.ToRight(3); got != defaultRight {
		t.Error(
			fmt.Sprintf(
				"Method ToRight with unknown argument %v must return default %v, got %v",
				3,
				defaultRight,
				got,
			),
		)
	}
}

func TestMapping_ToLeft(t *testing.T) {
	defaultLeft := 3

	mapping := New(testCases, WithDefaultLeft[int, string](defaultLeft))

	for left, right := range testCases {
		if got := mapping.ToLeft(right); got != left {
			t.Error(
				fmt.Sprintf(
					"Method ToLeft with argument %v must return %v, got %v",
					right,
					left,
					got,
				),
			)
		}
	}

	if got := mapping.ToLeft("{unk}"); got != defaultLeft {
		t.Error(
			fmt.Sprintf(
				"Method ToLeft with unknown argument %v must return default %v, got %v",
				"unknown",
				defaultLeft,
				got,
			),
		)
	}
}
