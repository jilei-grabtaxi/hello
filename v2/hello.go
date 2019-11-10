package hello

import (
	"rsc.io/quote/v3"
)
func Hello() string {
    return quote.HelloV3() + " hello / v2"
}

func Proverb() string {
    return quote.Concurrency() + + " hello / v2"
}
