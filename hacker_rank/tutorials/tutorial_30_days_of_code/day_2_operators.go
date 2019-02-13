package tutorial_30_days_of_code

import (
	"fmt"
	"math"
)

func operators(mealCost float64, tipPercent int32, taxPercent int32) {
	result := mealCost * float64(100 +tipPercent+taxPercent) / 100
	fmt.Println(math.Round(result))
}

func Operators(mealCost float64, tipPercent int32, taxPercent int32) {
	operators(mealCost, tipPercent, taxPercent)
}