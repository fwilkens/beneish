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

	result := dsri(90, 100, 150, 200)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}

	expected, err = decimal.NewFromString("1.92")

	if err != nil {
		t.Error(err)
	}

	result = dsri(50, 120, 400, 500)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestGmi(t *testing.T) {
	expected, err := decimal.NewFromString("0.6667")

	if err != nil {
		t.Error(err)
	}

	result := gmi(90, 100, 60, 50)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}

	expected, err = decimal.NewFromString("1.5000")

	if err != nil {
		t.Error(err)
	}

	result = gmi(100, 90, 50, 60)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestAqi(t *testing.T) {
	expected, err := decimal.NewFromString("1.0714")

	if err != nil {
		t.Error(err)
	}

	result := aqi(100, 200, 60, 600, 130, 210, 60, 700)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestDepi(t *testing.T) {
	expected, err := decimal.NewFromString("1.0938")

	if err != nil {
		t.Error(err)
	}

	result := depi(1000, 600, 1200, 900)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestLvgi(t *testing.T) {
	expected, err := decimal.NewFromString("1.1581")

	if err != nil {
		t.Error(err)
	}

	result := lvgi(1200, 400, 1800, 1300, 450, 1700)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestTata(t *testing.T) {
	expected, err := decimal.NewFromString("0.5833")

	if err != nil {
		t.Error(err)
	}

	result := tata(300, 400, 1200)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}
