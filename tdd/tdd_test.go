package tdd

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMultiples(t *testing.T) {
	if !reflect.DeepEqual([]int{3, 5}, Multiples(1, 5)) {
		t.Error()
	}
	assert.Equal(t, []int{3, 5}, Multiples(1, 5))

	
}

func TestMultiples_02(t *testing.T) {
	assert.Equal(t, []int{5, 6, 9, 10, 12, 15}, Multiples(5, 15))
}

func TestSomething(t *testing.T) {
	assert.True(t, true, "True is true!")
}

// func TestMultiples(t *testing.T) {
// 	assert.Equal(t, []int{3}, Multiples(1,3))
// }
