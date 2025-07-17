package commonz_test

import (
	"testing"

	"github.com/goosz/commonz"
	"github.com/stretchr/testify/assert"
)

type customStruct struct {
	A int
	B string
}

func TestZero_BasicTypes(t *testing.T) {
	assert.Equal(t, 0, commonz.Zero[int]())
	assert.Equal(t, "", commonz.Zero[string]())
	assert.Equal(t, false, commonz.Zero[bool]())
}

func TestZero_Struct(t *testing.T) {
	var want customStruct
	assert.Equal(t, want, commonz.Zero[customStruct]())
}

func TestZero_NilTypes(t *testing.T) {
	assert.Nil(t, commonz.Zero[[]int]())
	assert.Nil(t, commonz.Zero[*customStruct]())
	assert.Nil(t, commonz.Zero[map[string]int]())
	assert.Nil(t, commonz.Zero[chan int]())
	assert.Nil(t, commonz.Zero[func()]())
}
