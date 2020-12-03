package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	got := ParseInput(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`)

	want := []Password{
		{"abcde", 'a', 1, 3},
		{"cdefg", 'b', 1, 3},
		{"ccccccccc", 'c', 2, 9},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountValidPassword(t *testing.T) {
	got := CountValidPasswords([]Password{
		{"abcde", 'a', 1, 3},
		{"cdefg", 'b', 1, 3},
		{"ccccccccc", 'c', 2, 9},
	}, Password.isValid)

	if got != 2 {
		t.Errorf("expected to be 2 valid passwords")
	}
}

func TestValidPassword(t *testing.T) {
	p := Password{"abcde", 'a', 1, 3}

	if !p.isValid() {
		t.Errorf("expected to be valid: %v", p)
	}
}

func TestInvalidPassword(t *testing.T) {
	p := Password{"cdefg", 'b', 1, 3}

	if p.isValid() {
		t.Errorf("expected to be invalid: %v", p)
	}
}

func TestReallyValidPassword(t *testing.T) {
	p := Password{"abcde", 'a', 1, 3}

	if !p.isReallyValid() {
		t.Errorf("expected to be valid: %v", p)
	}
}

func TestReallyInvalidPassword1(t *testing.T) {
	p := Password{"cdefg", 'b', 1, 3}

	if p.isReallyValid() {
		t.Errorf("expected to be invalid: %v", p)
	}
}

func TestReallyInvalidPassword2(t *testing.T) {
	p := Password{"ccccccccc", 'c', 2, 9}

	if p.isReallyValid() {
		t.Errorf("expected to be invalid: %v", p)
	}
}
