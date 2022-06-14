package lists

import (
	"github.com/z-hua/gokit/types"
	"golang.org/x/exp/constraints"
)

// Copy 拷贝列表
func Copy[T any](l []T) []T {
	ll := make([]T, len(l))
	copy(ll, l)
	return ll
}

// All 判断列表中的所有元素是否都满足pred函数
func All[T any](l []T, pred func(T) bool) bool {
	for _, e := range l {
		if !pred(e) {
			return false
		}
	}
	return true
}

// Any 判断列表中是否存在元素满足pred函数
func Any[T any](l []T, pred func(T) bool) bool {
	for _, e := range l {
		if pred(e) {
			return true
		}
	}
	return false
}

// Insert 元素不存在时插入
func Insert[T constraints.Ordered](l []T, v T) ([]T, bool) {
	for _, e := range l {
		if e == v {
			return l, false
		}
	}
	l = append(l, v)
	return l, true
}

// Remove 移除列表成员
func Remove[T constraints.Ordered](l []T, v T) ([]T, bool) {
	for i, e := range l {
		if v == e {
			l = append(l[:i], l[i+1:]...)
			return l, true
		}
	}
	return l, false
}

// Member 判断是否为列表成员
func Member[T types.Number](l []T, v T) bool {
	for _, e := range l {
		if v == e {
			return true
		}
	}
	return false
}

// Min 返回列表中的最小值
func Min[T types.Number](l []T) T {
	result := l[0]
	for _, e := range l[1:] {
		if e < result {
			result = e
		}
	}
	return result
}

// Max 返回列表中的最大值
func Max[T types.Number](l []T) T {
	result := l[0]
	for _, e := range l[1:] {
		if e > result {
			result = e
		}
	}
	return result
}

// Sum 列表求和
func Sum[T types.Number](l []T) T {
	var sum T = 0
	for _, e := range l {
		sum += e
	}
	return sum
}

// Map 列表映射
func Map[T1, T2 any](l []T1, mapper func(T1) T2) []T2 {
	result := make([]T2, len(l))
	for i, e := range l {
		result[i] = mapper(e)
	}
	return result
}

// Filter 过滤掉列表中不满足pred函数的元素
func Filter[T any](l []T, pred func(T) bool) []T {
	result := make([]T, 0, len(l))
	for _, e := range l {
		if pred(e) {
			result = append(result, e)
		}
	}
	return result
}

// FilterMap 过滤列表并映射
func FilterMap[T1, T2 any](l []T1, fun func(T1) (T2, bool)) []T2 {
	result := make([]T2, 0, len(l))
	for _, e := range l {
		if v, ok := fun(e); ok {
			result = append(result, v)
		}
	}
	return result
}

// Partition 列表分割
func Partition[T any](l []T, pred func(T) bool) ([]T, []T) {
	satisfy := make([]T, 0, len(l))
	notSatisfy := make([]T, 0, len(l))
	for _, e := range l {
		if pred(e) {
			satisfy = append(satisfy, e)
		} else {
			notSatisfy = append(notSatisfy, e)
		}
	}
	return satisfy, notSatisfy
}

// Union 返回两个列表的并集
func Union[T constraints.Ordered](l1, l2 []T) []T {
	m := make(map[T]struct{})
	for _, e := range l1 {
		m[e] = struct{}{}
	}
	for _, e := range l2 {
		m[e] = struct{}{}
	}
	result := make([]T, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// Intersection 返回两个列表的交集
func Intersection[T constraints.Ordered](l1, l2 []T) []T {
	m1 := make(map[T]struct{})
	for _, e := range l1 {
		m1[e] = struct{}{}
	}
	m2 := make(map[T]struct{})
	for _, e := range l2 {
		m2[e] = struct{}{}
	}
	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	result := make([]T, 0, len(m1))
	for k := range m1 {
		if _, ok := m2[k]; ok {
			result = append(result, k)
		}
	}
	return result
}

// Difference 返回两个列表的差集
func Difference[T constraints.Ordered](l1, l2 []T) []T {
	m1 := make(map[T]struct{})
	for _, e := range l1 {
		m1[e] = struct{}{}
	}
	m2 := make(map[T]struct{})
	for _, e := range l2 {
		m2[e] = struct{}{}
	}
	result := make([]T, 0, len(m1))
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			result = append(result, k)
		}
	}
	return result
}
