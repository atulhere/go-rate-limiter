package main

import (
	"fmt"
	"math"
	"time"
)

type TokenBucket struct {
	tokens         float64
	maxTokens      float64
	refillRate     float64
	lastRefillTime time.Time
}

func NewTokenBucket(maxTokens float64, refillrate float64) *TokenBucket {

	return &TokenBucket{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillRate:     refillrate,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) refill() {

	now := time.Now()
	fmt.Println("current Time is", now)
	duration := now.Sub(tb.lastRefillTime)

	fmt.Println("Duration ", duration)
	tokensToAdd := tb.refillRate * duration.Seconds()
	fmt.Println("tokensToAdd ", tokensToAdd)
	tb.tokens = math.Min(tb.tokens+tokensToAdd, tb.maxTokens)
	fmt.Println("Current tokens ", tb.tokens)
	tb.lastRefillTime = now
}

func (tb *TokenBucket) request(tokens float64) bool {

	tb.refill()
	if tokens <= tb.tokens {

		tb.tokens = tb.tokens - tokens

		return true
	} else {
		return false
	}
}

func main() {

	tb := NewTokenBucket(10, 1)

	for i := 0; i < 20; i++ {
		fmt.Printf("Request %d: %v\n", i+1, tb.request(1))

		time.Sleep(500 * time.Millisecond)

	}

}
