package mapping

import "github.com/semichkin-gopkg/configurator"

type Configuration[L, R comparable] struct {
	DefaultLeft  L
	DefaultRight R
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
