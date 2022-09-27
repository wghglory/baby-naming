package naming

import (
	"math/rand"
	"time"
)

var names = []string{"Haoran", "Derek", "Iris", "Mike", "Tim", "Jenifer", "Ella", "Jianguo"}

func CreateBabyName() string {
	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(names))
	return names[index]
}
