package rand

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
)

// Random 从[min,max]之间随机出一个数
func Random[T constraints.Integer](min, max T) T {
	if min > max {
		min, max = max, min
	}
	return T(rand.Int63n(int64(max-min+1)) + int64(min))
}

// Choose 从列表中随机出一个元素
func Choose[T any](l []T) T {
	return l[rand.Intn(len(l))]
}

// ChooseN 从列表中随机出m个元素
func ChooseN[T any](l []T, n int) []T {
	if len(l) < n {
		panic(fmt.Sprintf("lenght of %v is smaller than %d", l, n))
	}
	m := len(l)
	if m == n {
		return l
	}
	// 蓄水池算法
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = l[i]
	}
	for i := n; i < m; i++ {
		r := Random(0, i)
		if r < n {
			result[r] = l[i]
		}
	}
	return result
}

// Shuffle 对列表进行乱序
func Shuffle[T any](l []T) {
	for i := len(l) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		l[i], l[j] = l[j], l[i]
	}
}
