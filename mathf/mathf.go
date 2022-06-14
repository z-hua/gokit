package mathf

import (
	"github.com/z-hua/gokit/types"
	"golang.org/x/exp/constraints"
)

// Clamp 将v的值限制在[min,max]之间
func Clamp[T types.Number](v, min, max T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Clamp01 将v的值限制在[0,1]之间
func Clamp01[T constraints.Float](v T) T {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

// Max 求两者之间的较大值
func Max[T types.Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min 求两者之间的较小值
func Min[T types.Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Ternary 三元运算
func Ternary[T any](cond bool, f1, f2 func() T) T {
	if cond {
		return f1()
	}
	return f2()
}
