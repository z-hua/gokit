package lists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestCopy(t *testing.T) {
	l1 := []int{1, 2, 3}
	l2 := Copy(l1)
	assert.Equal(t, l2, l1)
	assert.NotEqual(t, fmt.Sprintf("%p", &l1), fmt.Sprintf("%p", &l2))
}

func TestAll(t *testing.T) {
	pred := func(e int) bool {
		return e%2 == 0
	}

	l1 := []int{2, 4, 6}
	assert.Equal(t, true, All(l1, pred))

	l2 := []int{2, 4, 6, 7}
	assert.Equal(t, false, All(l2, pred))
}

func TestAny(t *testing.T) {
	pred := func(e int) bool {
		return e%2 == 0
	}

	l1 := []int{1, 2, 3}
	assert.Equal(t, true, Any(l1, pred))

	l2 := []int{1, 3, 5}
	assert.Equal(t, false, Any(l2, pred))
}

func TestInsert(t *testing.T) {
	l := []int{1, 2, 3}
	l1, b1 := Insert(l, 3)
	assert.Equal(t, []int{1, 2, 3}, l1)
	assert.Equal(t, false, b1)
	l2, b2 := Insert(l, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, l2)
	assert.Equal(t, true, b2)
}

func TestRemove(t *testing.T) {
	l1, b1 := Remove([]int{1, 2, 3}, 1)
	assert.Equal(t, true, b1)
	assert.Equal(t, []int{2, 3}, l1)

	l2, b2 := Remove([]int{1, 2, 3}, 4)
	assert.Equal(t, false, b2)
	assert.Equal(t, []int{1, 2, 3}, l2)

	l3, b3 := Remove([]int{1}, 1)
	assert.Equal(t, true, b3)
	assert.Equal(t, []int{}, l3)
}

func TestMember(t *testing.T) {
	l := []int{1, 2, 3}
	assert.Equal(t, true, Member(l, 1))
	assert.Equal(t, false, Member(l, 4))
}

func TestMin(t *testing.T) {
	l := []int{1, 2, 3}
	assert.Equal(t, 1, Min(l))
}

func TestMax(t *testing.T) {
	l := []int{1, 2, 3}
	assert.Equal(t, 3, Max(l))
}

func TestSum(t *testing.T) {
	l := []int{1, 2, 3}
	assert.Equal(t, 6, Sum(l))
}

func TestMap(t *testing.T) {
	l := []int{1, 2, 3}
	assert.Equal(t, []int{10, 20, 30}, Map(l, func(v int) int {
		return v * 10
	}))
}

func TestFilter(t *testing.T) {
	l := []int{1, 2, 3, 4, 5}
	assert.Equal(t, []int{1, 3, 5}, Filter(l, func(e int) bool { return e%2 == 1 }))
}

func TestFilterMap(t *testing.T) {
	l := []int{1, 2, 3, 4, 5}
	assert.Equal(t, []int{10, 30, 50}, FilterMap(l, func(e int) (int, bool) {
		if e%2 == 1 {
			return e * 10, true
		}
		return 0, false
	}))
}

func TestPartition(t *testing.T) {
	l := []int{1, 2, 3}
	l1, l2 := Partition(l, func(t int) bool {
		return t%2 == 0
	})
	assert.Equal(t, []int{2}, l1)
	assert.Equal(t, []int{1, 3}, l2)
}

func TestUnion(t *testing.T) {
	l1 := []int{1, 2, 3}
	l2 := []int{2, 3, 4}
	l := Union(l1, l2)
	sort.Ints(l)
	assert.Equal(t, []int{1, 2, 3, 4}, l)
}

func TestIntersection(t *testing.T) {
	l1 := []int{1, 2, 3}
	l2 := []int{2, 3, 4}
	l := Intersection(l1, l2)
	sort.Ints(l)
	assert.Equal(t, []int{2, 3}, l)
}

func TestDifference(t *testing.T) {
	l1 := []int{1, 2, 3}
	l2 := []int{2, 3, 4}
	l := Difference(l1, l2)
	sort.Ints(l)
	assert.Equal(t, []int{1}, l)
}
