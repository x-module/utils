package datastructure

import (
	"testing"

	"github.com/x-module/utils/utils/internal"
)

func TestSet_NewSetFromSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_NewSetFromSlice")

	s1 := NewSetFromSlice([]int{1, 2, 2, 3})
	assert.Equal(3, s1.Size())
	assert.Equal(true, s1.Contain(1))
	assert.Equal(true, s1.Contain(2))
	assert.Equal(true, s1.Contain(3))

	s2 := NewSetFromSlice([]int{})
	assert.Equal(0, s2.Size())
}

func TestSet_Add(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Add")

	set := NewSet[int]()
	set.Add(1, 2, 3)

	expected := NewSet(1, 2, 3)

	assert.Equal(true, set.Equal(expected))
}

func TestSet_AddIfNotExist(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_AddIfNotExist")

	set := NewSet[int]()
	set.Add(1, 2, 3)

	assert.Equal(false, set.AddIfNotExist(1))
	assert.Equal(true, set.AddIfNotExist(4))
	assert.Equal(NewSet(1, 2, 3, 4), set)
}

func TestSet_AddIfNotExistBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_AddIfNotExistBy")

	set := NewSet[int]()
	set.Add(1, 2)

	ok := set.AddIfNotExistBy(3, func(val int) bool {
		return val%2 != 0
	})

	notOk := set.AddIfNotExistBy(4, func(val int) bool {
		return val%2 != 0
	})

	assert.Equal(true, ok)
	assert.Equal(false, notOk)

	assert.Equal(true, set.Contain(3))
	assert.Equal(false, set.Contain(4))
}

func TestSet_Contain(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Contain")

	set := NewSet[int]()
	set.Add(1, 2, 3)

	assert.Equal(true, set.Contain(1))
	assert.Equal(false, set.Contain(4))
}

func TestSet_ContainAll(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_ContainAll")

	set1 := NewSet(1, 2, 3)
	set2 := NewSet(1, 2)
	set3 := NewSet(1, 2, 3, 4)

	assert.Equal(true, set1.ContainAll(set2))
	assert.Equal(false, set1.ContainAll(set3))
}

func TestSet_Clone(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Clone")

	set1 := NewSet(1, 2, 3)
	set2 := set1.Clone()

	assert.Equal(true, set1.Size() == set2.Size())
	assert.Equal(true, set1.ContainAll(set2))
}

func TestSet_Delete(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Delete")

	set := NewSet[int]()
	set.Add(1, 2, 3)
	set.Delete(3)

	expected := NewSet(1, 2)

	assert.Equal(true, set.Equal(expected))
}

func TestSet_Equal(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Equal")

	set1 := NewSet(1, 2, 3)
	set2 := NewSet(1, 2, 3)
	set3 := NewSet(1, 2, 3, 4)

	assert.Equal(true, set1.Equal(set2))
	assert.Equal(false, set1.Equal(set3))
}

func TestSet_Iterate(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Iterate")

	set := NewSet(1, 2, 3)
	arr := []int{}
	set.Iterate(func(value int) {
		arr = append(arr, value)
	})

	assert.Equal(3, len(arr))
}

func TestSet_IsEmpty(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_IsEmpty")

	set := NewSet[int]()
	assert.Equal(true, set.IsEmpty())
}

func TestSet_Size(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Size")

	set := NewSet(1, 2, 3)
	assert.Equal(3, set.Size())
}

func TestSet_Values(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Values")

	set := NewSet(1, 2, 3)
	values := set.Values()

	assert.Equal(3, len(values))
}

func TestSet_Union(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Union")

	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4, 5)
	expected := NewSet(1, 2, 3, 4, 5)
	unionSet := set1.Union(set2)

	assert.Equal(expected, unionSet)
}

func TestSet_Intersection(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Intersection")

	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4, 5)
	expected := NewSet(2, 3)
	intersectionSet := set1.Intersection(set2)

	assert.Equal(expected, intersectionSet)
}

func TestSet_SymmetricDifference(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_SymmetricDifference")

	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4, 5)

	assert.Equal(NewSet(1, 4, 5), set1.SymmetricDifference(set2))
}

func TestSet_Minus(t *testing.T) {
	assert := internal.NewAssert(t, "TestSet_Minus")

	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4, 5)
	set3 := NewSet(2, 3)

	assert.Equal(NewSet(1), set1.Minus(set2))
	assert.Equal(NewSet(4, 5), set2.Minus(set3))
}
