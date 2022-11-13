package twocrystalballs

import (
	"math/rand"
	"testing"
	"time"
)

func TestTwoCrystalBalls(t *testing.T) {
	// set seed
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(10000)
	data := make([]bool, 10000)
	for i := idx; i < 10000; i += 1 {
		data[i] = true
	}
	response := TwoCrystalBalls(data)
	if response != idx {
		t.Errorf("error")
	}

	response_ := TwoCrystalBalls( make([]bool, 10000))
	if response_ != -1 {
		t.Errorf("error")
	}
}
