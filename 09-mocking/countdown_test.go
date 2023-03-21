package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpyCountdownOperations{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := []string{write, sleep, write, sleep, write, sleep, write}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}

}

type SpyCountdownOperations struct {
	calls []string
}

const sleep = "sleep"
const write = "write"

func (s *SpyCountdownOperations) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *SpyCountdownOperations) Write(_ []byte) (_ int, _ error) {
	s.calls = append(s.calls, write)
	return
}
