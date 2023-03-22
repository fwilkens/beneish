package beneish

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestDsri(t *testing.T) {
	expected, err := decimal.NewFromString("0.8333")

	if err != nil {
		t.Error(err)
	}

	result := dsri(100, 90, 200, 150)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}

	expected, err = decimal.NewFromString("1.92")

	if err != nil {
		t.Error(err)
	}

	result = dsri(120, 50, 500, 400)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}
