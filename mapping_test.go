package mapping

import (
	"fmt"
	"strings"
	"testing"
)

var cases = map[int]string{
	0: "0",
	1: "1",
	2: "2",
}

func TestMapping_ToRight(t *testing.T) {
	mapping := New(cases)

	for left, right := range cases {
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
}

func TestMapping_ToLeft(t *testing.T) {
	mapping := New(cases)

	for left, right := range cases {
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
}

func TestWithDefaultRight(t *testing.T) {
	defaultRight := "default"

	mapping := New(cases, WithDefaultRight[int, string](defaultRight))

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

func TestWithDefaultLeft(t *testing.T) {
	defaultLeft := 100

	mapping := New(cases, WithDefaultLeft[int, string](defaultLeft))

	if got := mapping.ToLeft("{unk}"); got != defaultLeft {
		t.Error(
			fmt.Sprintf(
				"Method ToLeft with unknown argument %v must return default %v, got %v",
				"{unk}",
				defaultLeft,
				got,
			),
		)
	}
}

func TestWithLeftComparator(t *testing.T) {
	mapping := New(cases, WithLeftComparator[int, string](func(a, b int) bool {
		if a == b {
			return true
		}

		return (a == 3 && b == 0) || (a == 4 && b == 1) || (a == 5 && b == 2)
	}))

	testCases := map[int]int{
		0: 0,
		1: 1,
		2: 2,
		3: 0,
		4: 1,
		5: 2,
	}

	for a, b := range testCases {
		aGot := mapping.ToRight(a)
		bGot := mapping.ToRight(b)

		if aGot != bGot {
			t.Error(
				fmt.Sprintf(
					"By configurator WithLeftComparator ToRight for %v must be equal with ToRight for %v; got %v and %v",
					a,
					b,
					aGot,
					bGot,
				),
			)
		}
	}
}

func TestWithRightComparator(t *testing.T) {
	mapping := New(cases, WithRightComparator[int, string](func(a, b string) bool {
		if a == b {
			return true
		}

		return strings.Contains(a, b)
	}))

	testCases := map[string]string{
		" 1 ":       "1",
		"test1test": "1",
		"1":         "1",
	}

	for a, b := range testCases {
		aGot := mapping.ToLeft(a)
		bGot := mapping.ToLeft(b)

		if aGot != bGot {
			t.Error(
				fmt.Sprintf(
					"By configurator WithRightComparator ToLeft for %v must be equal with ToLeft for %v; got %v and %v",
					a,
					b,
					aGot,
					bGot,
				),
			)
		}
	}
}
