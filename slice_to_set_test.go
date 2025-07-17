package commonz_test

import (
	"testing"

	"github.com/goosz/commonz"
	"github.com/stretchr/testify/assert"
)

func TestSliceToSet_NoDuplicates(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	set, err := commonz.SliceToSet(slice, false)
	assert.NoError(t, err)
	assert.Len(t, set, 4)
	for _, v := range slice {
		assert.Contains(t, set, v)
	}
}

func TestSliceToSet_WithDuplicates_StrictFalse(t *testing.T) {
	slice := []string{"a", "b", "a", "c"}
	set, err := commonz.SliceToSet(slice, false)
	assert.NoError(t, err)
	assert.Len(t, set, 3)
	for _, v := range []string{"a", "b", "c"} {
		assert.Contains(t, set, v)
	}
}

func TestSliceToSet_WithDuplicates_StrictTrue(t *testing.T) {
	slice := []int{1, 2, 2, 3}
	_, err := commonz.SliceToSet(slice, true)
	assert.Error(t, err)
}

func TestSliceToSet_EmptySlice(t *testing.T) {
	set, err := commonz.SliceToSet([]int{}, false)
	assert.NoError(t, err)
	assert.Empty(t, set)
}

func TestSliceToSet_SingleElement(t *testing.T) {
	set, err := commonz.SliceToSet([]string{"x"}, false)
	assert.NoError(t, err)
	assert.Len(t, set, 1)
	assert.Contains(t, set, "x")
}

func TestSliceToSet_CustomStruct(t *testing.T) {
	type pair struct{ A, B int }
	slice := []pair{{1, 2}, {3, 4}, {1, 2}}
	set, err := commonz.SliceToSet(slice, false)
	assert.NoError(t, err)
	assert.Len(t, set, 2)
}
