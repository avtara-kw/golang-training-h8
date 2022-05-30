package test

import (
	"fmt"
	"testing"
	"unit-test/helpers"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	nums := []struct {
		req []int
		res int
		msg string
	}{
		{
			req: []int{10, 20, 30},
			res: 60,
		},
		{
			req: []int{30, 20, 30},
			res: 80,
		},
	}

	for k, num := range nums {
		t.Run(fmt.Sprintf("test-%d", k+1), func(t *testing.T) {
			require.Equal(t, num.res, helpers.Sum(num.req...))
		})
	}
}

func TestSumBy0(t *testing.T) {
	nums := []int{0, 0, 0}
	expected := -1

	result := helpers.Sum(nums...)

	require.Equal(t, expected, result)
}
