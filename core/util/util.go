package util

import "math"

func Clamp(val, min, max int) int {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

func DegToRad(angle float64) float64 {
	return angle * math.Pi / 180
}
