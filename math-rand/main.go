package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// Rolls a dice and returns a random int d with 1 <= d <= 20.
	dice20 := rand.Intn(20-1) + 1
	fmt.Printf("Result of your D20 is: %d\n", dice20)

	// returns a random float64 f with 0.0 <= f < 12.0.
	randFloat := rand.Float64() * (12.0 - 0.0)
	fmt.Printf("Result of your random float between 12.0 and 0.0 is: %f\n", randFloat)

	// Shuffle a slice of strings
	cards := []string{"1-Swamp", "2-Mountain", "3-Island", "4-Forest", "5-Plains"}
	fmt.Printf("slice before shuffle: %s\n", cards)
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	fmt.Printf("slice after shuffle: %s\n", cards)

}
