package mapping

import (
	"github.com/semichkin-gopkg/configurator"
)

type Mapping[L, R comparable] struct {
	configuration Configuration[L, R]

	left  map[L]R
	right map[R]L
}

func New[L, R comparable](
	left map[L]R,
	updaters ...configurator.Updater[Configuration[L, R]],
) *Mapping[L, R] {
	var defaultL L
	var defaultR R

	configuration := configurator.New[Configuration[L, R]]().
		Append(
			WithDefaultLeft[L, R](defaultL),
			WithDefaultRight[L, R](defaultR),
		).
		Append(updaters...).
		Apply()

	right := map[R]L{}

	for l, r := range left {
		right[r] = l
	}

	return &Mapping[L, R]{
		configuration: configuration,
		left:          left,
		right:         right,
	}
}

func (m *Mapping[L, R]) Set(left L, right R) {
	m.left[left] = right
	m.right[right] = left
}

func (m *Mapping[L, R]) ToRight(from L) R {
	if m.configuration.LeftComparator == nil {
		if right, ok := m.left[from]; ok {
			return right
		}

		return m.configuration.DefaultRight
	}

	for left, right := range m.left {
		if m.configuration.LeftComparator(from, left) {
			return right
		}
	}

	return m.configuration.DefaultRight
}

func (m *Mapping[L, R]) ToLeft(from R) L {
	if m.configuration.RightComparator == nil {
		if left, ok := m.right[from]; ok {
			return left
		}

		return m.configuration.DefaultLeft
	}

	for right, left := range m.right {
		if m.configuration.RightComparator(from, right) {
			return left
		}
	}

	return m.configuration.DefaultLeft
}

func (m *Mapping[L, R]) Lefts() []L {
	lefts := make([]L, 0, len(m.left))
	for l := range m.left {
		lefts = append(lefts, l)
	}
	return lefts
}

func (m *Mapping[L, R]) Rights() []R {
	rights := make([]R, 0, len(m.right))
	for r := range m.right {
		rights = append(rights, r)
	}
	return rights
}
