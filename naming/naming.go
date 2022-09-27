package naming

import (
	"math/rand"
	"time"
)

var maleNames = []string{"Haoran", "Derek", "Mike", "Tim", "Jianguo"}
var femaleNames = []string{"Iris", "Jenifer", "Ella"}

func CreateBabyName(male bool) string {
	rand.Seed(time.Now().UnixNano())

	if male {
		index := rand.Intn(len(maleNames))
		return maleNames[index]
	} else {
		index := rand.Intn(len(femaleNames))
		return femaleNames[index]
	}
}
