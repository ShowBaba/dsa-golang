package twocrystalballs

import (
	"math"
)

// Given two crystal balls that will break if dropped from high enough
// distance, determine the exact spot in which it will break in the most
// optimized way.

func TwoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(breaks))))
	i := jumpAmount
	for ; i < len(breaks); i += jumpAmount { // jump sqrt of N
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount

	for j := 0; j < jumpAmount && i < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}
	return -1
}
