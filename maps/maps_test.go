package maps

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"testing"
)

func TestCopy(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	m2 := Copy(m1)
	assert.Equal(t, m1, m2)
	assert.NotEqual(t, fmt.Sprintf("%p", &m1), fmt.Sprintf("%p", &m2))
}

func TestKeys(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	keys := Keys(m)
	sort.Ints(keys)
	assert.Equal(t, []int{1, 2, 3}, keys)
}

func TestValues(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	values := Values(m)
	sort.Strings(values)
	assert.Equal(t, []string{"a", "b", "c"}, values)
}

func TestWith(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	assert.Equal(t, map[int]string{1: "a"}, With(m1, []int{1}))

	m2 := map[int]string{1: "a", 2: "b", 3: "c"}
	assert.Equal(t, map[int]string{1: "a", 2: "b"}, With(m2, []int{1, 2}))
}

func TestWithout(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	assert.Equal(t, map[int]string{2: "b", 3: "c"}, Without(m1, []int{1}))

	m2 := map[int]string{1: "a", 2: "b", 3: "c"}
	assert.Equal(t, map[int]string{3: "c"}, Without(m2, []int{1, 2}))
}

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	m2 := map[int]string{1: "1", 2: "2", 3: "3"}
	assert.Equal(t, m2, Map(m1, func(k int, v int) string { return strconv.Itoa(v) }))
}

func TestFilter(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	m2 := map[int]int{2: 2}
	assert.Equal(t, m2, Filter(m1, func(k int, v int) bool { return v%2 == 0 }))
}

func TestFilterMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	m2 := map[int]string{2: "2"}
	assert.Equal(t, m2, FilterMap(m1, func(k int, v int) (string, bool) {
		if v%2 == 0 {
			return strconv.Itoa(v), true
		} else {
			return "", false
		}
	}))
}

func TestMerge(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b"}
	m2 := map[int]string{1: "aa", 3: "c"}
	assert.Equal(t, map[int]string{1: "aa", 2: "b", 3: "c"}, Merge(m1, m2))
}

func TestMergeWith(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b"}
	m2 := map[int]string{1: "aa", 3: "c"}
	m3 := MergeWith(m1, m2, func(k int, v1, v2 string) string {
		return v1 + v2
	})
	assert.Equal(t, map[int]string{1: "aaa", 2: "b", 3: "c"}, m3)
}

func TestMergeWithAdd(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 10, "c": 3}
	assert.Equal(t, map[string]int{"a": 11, "b": 2, "c": 3}, MergeWithAdd(m1, m2))
}

func TestMergeWithSub(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 10, "c": 3}
	assert.Equal(t, map[string]int{"a": -9, "b": 2, "c": -3}, MergeWithSub(m1, m2))
}

func TestListAppend(t *testing.T) {
	m := make(map[int][]string)
	b1 := ListAppend(m, 1, "a1", false)
	assert.Equal(t, map[int][]string{1: {"a1"}}, m)
	assert.Equal(t, true, b1)
	b2 := ListAppend(m, 1, "a2", false)
	assert.Equal(t, map[int][]string{1: {"a1", "a2"}}, m)
	assert.Equal(t, true, b2)
	b3 := ListAppend(m, 2, "b1", false)
	assert.Equal(t, map[int][]string{1: {"a1", "a2"}, 2: {"b1"}}, m)
	assert.Equal(t, true, b3)
	b4 := ListAppend(m, 3, "c1", true)
	assert.Equal(t, map[int][]string{1: {"a1", "a2"}, 2: {"b1"}}, m)
	assert.Equal(t, false, b4)
	b5 := ListAppend(m, 2, "b2", true)
	assert.Equal(t, map[int][]string{1: {"a1", "a2"}, 2: {"b1", "b2"}}, m)
	assert.Equal(t, true, b5)
}

func TestListInsert(t *testing.T) {
	m := make(map[int][]string)
	b1 := ListInsert(m, 1, "a1", false)
	assert.Equal(t, map[int][]string{1: {"a1"}}, m)
	assert.Equal(t, true, b1)
	b2 := ListInsert(m, 1, "a1", false)
	assert.Equal(t, map[int][]string{1: {"a1"}}, m)
	assert.Equal(t, false, b2)
	b3 := ListInsert(m, 2, "b1", true)
	assert.Equal(t, map[int][]string{1: {"a1"}}, m)
	assert.Equal(t, false, b3)
}

func TestListRemove(t *testing.T) {
	m := map[int][]string{1: {"a1", "a2", "a3"}, 2: {"b1"}}
	b1 := ListRemove(m, 1, "a2")
	assert.Equal(t, map[int][]string{1: {"a1", "a3"}, 2: {"b1"}}, m)
	assert.Equal(t, true, b1)
	b2 := ListRemove(m, 2, "b1")
	assert.Equal(t, map[int][]string{1: {"a1", "a3"}}, m)
	assert.Equal(t, true, b2)
	b3 := ListRemove(m, 2, "b1")
	assert.Equal(t, map[int][]string{1: {"a1", "a3"}}, m)
	assert.Equal(t, false, b3)
}
