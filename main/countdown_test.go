package main

import (
	"bytes"
	"errors"
	"testing"
	"reflect"
)


type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}


type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write([]byte) (int, error) {
	s.Calls = append(s.Calls, write)
	return 1, errors.New("")
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
			buffer := &bytes.Buffer{}
			Countdown(buffer, &SpyCountdownOperations{})

			got := buffer.String()
			want := `3
2
1
Go!`

			if got != want {
					t.Errorf("got %q want %q", got, want)
			}
	})

	t.Run("sleep before every print", func(t *testing.T) {
			spySleepPrinter := &SpyCountdownOperations{}
			Countdown(spySleepPrinter, spySleepPrinter)

			want := []string{
					sleep,
					write,
					sleep,
					write,
					sleep,
					write,
					sleep,
					write,
			}

			if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
					t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
			}
	})
}