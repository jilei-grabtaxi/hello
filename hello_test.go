package hello

import (
    "testing"
)

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

func TestProverb(t *testing.T) {
    want := "Concurrency is not parallelism."
    if got := Proverb(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}

func TestCallV2(t *testing.T){
	want := "Hello, world. hello / v2"
    if got := CallV2(); got != want {
        t.Errorf("CallV2() = %q, want %q", got, want)
    }
}

func TestCallV3(t *testing.T){
	want := "Hello, world. hellov3"
    if got := CallV3(); got != want {
        t.Errorf("CallV3() = %q, want %q", got, want)
    }
}
