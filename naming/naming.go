package naming

import (
	"fmt"
	"math/rand"
	"time"
)

var maleNames = []string{"Haoran", "Derek", "Mike", "Tim", "Jianguo"}
var femaleNames = []string{"Iris", "Jenifer", "Ella"}

func CreateBabyName(male bool, minLen int) (string, error) {
	rand.Seed(time.Now().UnixNano())

	var result string
	var names []string

	if male {
		names = maleNames
	} else {
		names = femaleNames
	}

	nameLen := len(names)

	rand.Shuffle(nameLen, func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	for i := 0; i < nameLen; i++ {
		if len(names[i]) >= minLen {
			result = names[i]
			return result, nil
		}
	}

	return "", fmt.Errorf("no baby found with min length %v", minLen)
}
