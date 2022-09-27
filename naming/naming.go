package naming

import (
	"math/rand"
)

var names = []string{"Haoran", "Derek", "Iris", "Mike", "Tim", "Jenifer", "Ella", "Jianguo"}

func CreateBabyName() string {
	index := rand.Intn(len(names) - 1)
	return names[index]
}
