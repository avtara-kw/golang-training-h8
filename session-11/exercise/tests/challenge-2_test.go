package test

import (
	"exercise/helpers"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFight(t *testing.T) {
	person := []struct {
		enemies       []int
		enemies_power []int
		power         int
		res           int
		msg           string
	}{
		{
			enemies:       []int{5, 3, 9, 8},
			enemies_power: []int{2, 2, 3, 1},
			power:         3,
			res:           7,
		},
		{
			enemies:       []int{2, 6, 3, 9},
			enemies_power: []int{2, 2, 3, 5},
			power:         2,
			res:           14,
		},
	}

	for k, people := range person {
		t.Run(fmt.Sprintf("test-%d", k+1), func(t *testing.T) {
			require.Equal(t, people.res, helpers.Fight(people.enemies, people.enemies_power, people.power))
		})
	}
}
