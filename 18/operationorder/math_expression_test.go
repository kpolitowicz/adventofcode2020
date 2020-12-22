package operationorder

import (
	"reflect"
	"testing"
)

func TestCalculateSimple(t *testing.T) {
	expr := newExpression("1 + 2 * 3 + 4 * 5 + 6")
	got := expr.Calculate()
	want := uint64(71)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCalculateParen(t *testing.T) {
	expr := newExpression("1 + (2 * 3)")
	got := expr.Calculate()
	want := uint64(7)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCalculateParenComplex(t *testing.T) {
	expr := newExpression("1 + (2 * 3) + (4 * (5 + 6))")
	got := expr.Calculate()
	want := uint64(51)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCalculateParenEvenMoreComplex(t *testing.T) {
	expr := newExpression("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")
	got := expr.Calculate()
	want := uint64(13632)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
