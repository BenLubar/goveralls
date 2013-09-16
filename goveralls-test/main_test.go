package main

import (
	"testing"
)

func TestIsSix(t *testing.T) {
	for i := 0; i < 10; i++ {
		if is, err := IsSix(i); err != nil {
			t.Errorf("IsSix(%d) returned an error: %v", err)
		} else if is != (i == 6) {
			t.Errorf("IsSix(%d) returned %b but it should have returned %b", i, is, !is)
		}
	}
}

func TestEndsWithSix(t *testing.T) {
	if ok, err := EndsWithSix(42); err != nil {
		t.Errorf("42 %v", err)
	} else if ok {
		t.Errorf("42 ends with six?!?!?!?!")
	}
}
