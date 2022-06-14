package mathf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClamp(t *testing.T) {
	assert.Equal(t, 50, Clamp(50, 1, 100))
	assert.Equal(t, 1, Clamp(0, 1, 100))
	assert.Equal(t, 100, Clamp(101, 1, 100))
}

func TestClamp01(t *testing.T) {
	assert.Equal(t, 0.5, Clamp01(0.5))
	assert.Equal(t, 0.0, Clamp01(-1.0))
	assert.Equal(t, 1.0, Clamp01(2.0))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, Max(1, 2))
	assert.Equal(t, 2, Max(2, 1))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, 1, Min(2, 1))
}

func TestTernary(t *testing.T) {
	var a int
	a = 3
	assert.Equal(t, 1, Ternary(a > 1, func() int { return 1 }, func() int { return 2 }))
	assert.Equal(t, 2, Ternary(a < 1, func() int { return 1 }, func() int { return 2 }))
}
