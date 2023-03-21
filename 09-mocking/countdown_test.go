package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
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

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
