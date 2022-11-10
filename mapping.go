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

func (m *Mapping[L, R]) ToRight(from L) R {
	if r, ok := m.left[from]; ok {
		return r
	}

	return m.configuration.DefaultRight
}

func (m *Mapping[L, R]) ToLeft(from R) L {
	if l, ok := m.right[from]; ok {
		return l
	}

	return m.configuration.DefaultLeft
}
