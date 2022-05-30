package helpers

import (
	"sort"
)

func Fight(enemies []int, enemies_powers []int, power int) int {
	enemyPower := map[int]int{}
	if len(enemies) != len(enemies_powers) {
		return -1
	}
	for i := 0; i < len(enemies); i++ {
		enemyPower[enemies[i]] = enemies_powers[i]
	}

	sort.Ints(enemies)
	for _, enemy := range enemies {
		if power >= enemy {
			power += enemyPower[enemy]
		} else {
			break
		}
	}

	return power
}
