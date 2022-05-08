package utils

import (
	"testing"
)

func TestParseFloat(t *testing.T) {

	got := ParseFloat("1234")
	want := 1234.0

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestUint(t *testing.T) {

	got := ParseUint("1234")
	var want uint64 = 1234

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
