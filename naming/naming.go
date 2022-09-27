package naming

import (
	"fmt"
	"math/rand"
)

var maleNames = []string{"Haoran", "Derek", "Mike", "Tim", "Jianguo"}
var femaleNames = []string{"Iris", "Jenifer", "Ella"}

// return a baby name by gender and min length
func CreateBabyName(male bool, minLen uint) (string, error) {
	names := getBabyNames(male)

	for i := 0; i < len(names); i++ {
		if uint(len(names[i])) >= minLen {
			return names[i], nil
		}
	}

	return "", fmt.Errorf("no baby found with min length %v", minLen)
}

// get baby names randomly (shuffled)
func getBabyNames(male bool) []string {
	var names []string

	if male {
		names = maleNames
	} else {
		names = femaleNames
	}

	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	return names
}
