package helloV3

import (
	"rsc.io/quote/v3"
)
func Hello() string {
    return quote.HelloV3() + " hellov3"
}

func Proverb() string {
    return quote.Concurrency()+ " hellov3"

}
