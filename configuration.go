package mapping

import "github.com/semichkin-gopkg/configurator"

type Configuration[L, R comparable] struct {
	DefaultLeft  L
	DefaultRight R

	LeftComparator  func(L, L) bool
	RightComparator func(R, R) bool
}

func WithDefaultLeft[L, R comparable](left L) configurator.Updater[Configuration[L, R]] {
	return func(c *Configuration[L, R]) {
		c.DefaultLeft = left
	}
}

func WithDefaultRight[L, R comparable](right R) configurator.Updater[Configuration[L, R]] {
	return func(c *Configuration[L, R]) {
		c.DefaultRight = right
	}
}

func WithLeftComparator[L, R comparable](comparator func(L, L) bool) configurator.Updater[Configuration[L, R]] {
	return func(c *Configuration[L, R]) {
		c.LeftComparator = comparator
	}
}

func WithRightComparator[L, R comparable](comparator func(R, R) bool) configurator.Updater[Configuration[L, R]] {
	return func(c *Configuration[L, R]) {
		c.RightComparator = comparator
	}
}
