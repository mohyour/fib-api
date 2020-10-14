package fibonacci

import (
	"fmt"
	"math/big"
)

var (
	// Cache to store fibonacci values
	fibCache map[int]*big.Int = make(map[int]*big.Int)
)

// FibResponse - struct for json response
type FibResponse struct {
	FibTerm          int      `json:"term,omitempty"`
	FibSequenceValue *big.Int `json:"fib_sequence_value,omitempty"`
	Message          string   `json:"message,omitempty"`
}

// Fibonacci - gets fibonacci sequence of nth term and add to cache
// Return value as big.Int to accomodate getting sequence for over 93
func Fibonacci(num int) *big.Int {
	first := big.NewInt(0)
	second := big.NewInt(1)

	if num <= 1 {
		return big.NewInt(int64(num))
	}

	// Retrieve from cache if available
	value, ok := fibCache[num]
	if ok {
		fmt.Println(fibCache)
		return value
	}

	for i := 1; i <= num; i++ {
		second.Add(second, first)
		first, second = second, first
	}
	// Add to cache
	fibCache[num] = first
	fmt.Println(first)
	return first
}
