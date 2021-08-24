package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaths(t *testing.T) {
	t.Run("1 and 1 is 2", func(t *testing.T) {
		channel := make(chan Equation)
		go Add(1, 1, channel)
		sum := <- channel
		assert.Equal(t, float64(2), sum)
	})
}
