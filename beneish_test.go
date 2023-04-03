package beneish

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestDsri(t *testing.T) {
	result := dsri(90, 100, 150, 200)
	expected, err := decimal.NewFromString("0.8333")
	if err != nil {
		t.Error(err)
	}

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestGmi(t *testing.T) {
	result := gmi(90, 100, 60, 50)
	expected, err := decimal.NewFromString("0.6667")
	if err != nil {
		t.Error(err)
	}

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestAqi(t *testing.T) {
	result := aqi(100, 200, 60, 600, 130, 210, 60, 700)
	expected, err := decimal.NewFromString("1.0714")
	if err != nil {
		t.Error(err)
	}

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestSgi(t *testing.T) {
	result := sgi(1200, 1300)
	expected, err := decimal.NewFromString("1.0833")
	if err != nil {
		t.Error(err)
	}

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestDepi(t *testing.T) {
	result := depi(1000, 600, 1200, 900)
	expected, err := decimal.NewFromString("1.0938")
	if err != nil {
		t.Error(err)
	}

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestLvgi(t *testing.T) {
	result := lvgi(1200, 400, 1800, 1300, 450, 1700)
	expected, err := decimal.NewFromString("1.1581")
	if err != nil {
		t.Error(err)
	}

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestTata(t *testing.T) {
	result := tata(300, 400, 1200)
	expected, err := decimal.NewFromString("0.5833")
	if err != nil {
		t.Error(err)
	}

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}
