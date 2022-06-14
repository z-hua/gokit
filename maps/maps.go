package maps

import (
	"github.com/z-hua/gokit/types"
	"golang.org/x/exp/constraints"
)

// Copy 拷贝map
func Copy[K constraints.Ordered, V any](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

// Keys 返回字典的键列表
func Keys[K constraints.Ordered, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// Values 返回字典的值列表
func Values[K constraints.Ordered, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// With 保留字典中键在keys列表中的项
func With[K constraints.Ordered, V any](m map[K]V, keys []K) map[K]V {
	result := make(map[K]V)
	for _, k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}
	return result
}

// Without 保留字典中键不在keys列表中的项
func Without[K constraints.Ordered, V any](m map[K]V, keys []K) map[K]V {
	result := Copy(m)
	for _, k := range keys {
		delete(result, k)
	}
	return result
}

// Map 字典映射
func Map[K constraints.Ordered, V1, V2 any](m map[K]V1, fun func(K, V1) V2) map[K]V2 {
	result := make(map[K]V2)
	for k, v := range m {
		result[k] = fun(k, v)
	}
	return result
}

// Filter 字典过滤，移除不满足pred函数的元素
func Filter[K constraints.Ordered, V any](m map[K]V, pred func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if pred(k, v) {
			result[k] = v
		}
	}
	return result
}

// FilterMap 过滤字典并映射
func FilterMap[K constraints.Ordered, V1, V2 any](m map[K]V1, pred func(K, V1) (V2, bool)) map[K]V2 {
	result := make(map[K]V2)
	for k, v := range m {
		if vv, ok := pred(k, v); ok {
			result[k] = vv
		}
	}
	return result
}

// Merge 合并两个map，存在相同key时使用m2中的值
func Merge[K constraints.Ordered, V any](m1, m2 map[K]V) map[K]V {
	result := Copy(m1)
	for k, v2 := range m2 {
		result[k] = v2
	}
	return result
}

// MergeWith 合并两个map，存在相同key时使用combiner函数的返回值
func MergeWith[K constraints.Ordered, V any](m1, m2 map[K]V, combiner func(k K, v1, v2 V) V) map[K]V {
	result := Copy(m1)
	for k, v2 := range m2 {
		if v1, ok := result[k]; ok {
			result[k] = combiner(k, v1, v2)
		} else {
			result[k] = v2
		}
	}
	return result
}

// MergeWithAdd 合并两个map，存在相同key时将其值相加
func MergeWithAdd[K constraints.Ordered, V types.Number](m1, m2 map[K]V) map[K]V {
	result := Copy(m1)
	for k, v2 := range m2 {
		result[k] += v2
	}
	return result
}

func MergeWithSub[K constraints.Ordered, V types.Number](m1, m2 map[K]V) map[K]V {
	result := Copy(m1)
	for k, v2 := range m2 {
		result[k] -= v2
	}
	return result
}

// ListAppend 列表扩展
func ListAppend[K constraints.Ordered, V any](m map[K][]V, k K, v V, exist bool) bool {
	l, ok := m[k]
	if !ok {
		if exist {
			return false
		}
		m[k] = []V{v}
		return true
	}
	m[k] = append(l, v)
	return true
}

// ListInsert 列表插入
func ListInsert[K constraints.Ordered, V constraints.Ordered](m map[K][]V, k K, v V, exist bool) bool {
	l, ok := m[k]
	if !ok {
		if exist {
			return false
		}
		m[k] = []V{v}
		return true
	}
	for _, e := range l {
		if v == e {
			return false
		}
	}
	m[k] = append(l, v)
	return true
}

// ListRemove 列表移除
func ListRemove[K constraints.Ordered, V constraints.Ordered](m map[K][]V, k K, v V) bool {
	l, ok := m[k]
	if !ok {
		return false
	}
	for i, e := range l {
		if v == e {
			if len(l) == 1 {
				delete(m, k)
			} else {
				m[k] = append(l[:i], l[i+1:]...)
			}
			return true
		}
	}
	return false
}
