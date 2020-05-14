package basic

import (
	"errors"
	"math"
)

// Export function has to be Captilized
func Sum(x int, y int) int {
	return x + y
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined  for negative numebrs")
	}

	return math.Sqrt(x), nil
}
