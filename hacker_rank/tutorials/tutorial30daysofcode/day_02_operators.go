package tutorial30daysofcode

import (
	"math"
)

func operators(mealCost float64, tipPercent, taxPercent int32) {
	result := mealCost * float64(100+tipPercent+taxPercent) / 100
	_, _ = funcPrint(math.Round(result))
}
