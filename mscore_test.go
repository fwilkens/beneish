package beneish

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestDsri(t *testing.T) {
	result := Dsri(90, 100, 150, 200)
	expected := decimal.NewFromFloat(0.8333)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestGmi(t *testing.T) {
	result := Gmi(90, 100, 60, 50)
	expected := decimal.NewFromFloat(0.6667)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestAqi(t *testing.T) {
	result := Aqi(100, 60, 600, 130, 60, 700)
	expected := decimal.NewFromFloat(0.9935)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestSgi(t *testing.T) {
	result := Sgi(1200, 1300)
	expected := decimal.NewFromFloat(1.0833)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestDepi(t *testing.T) {
	result := Depi(1000, 600, 1200, 900)
	expected := decimal.NewFromFloat(1.0938)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestLvgi(t *testing.T) {
	result := Lvgi(1200, 400, 1800, 1300, 450, 1700)
	expected := decimal.NewFromFloat(1.1581)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestTata(t *testing.T) {
	result := Tata(300, 50, 1200)
	expected := decimal.NewFromFloat(0.2083)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// Example data from:
// https://apps.kelley.iu.edu/Beneish/Content/Examples/SunbeamExample.pdf
func TestMscore(t *testing.T) {
	mscore := MscoreCalc(
		2134,  // prevNetReceivables
		6242,  // prevAssets
		9006,  // prevCogs
		2201,  // prevPPE
		10727, // prevTotalAssets
		9842,  // prevSales
		474,   // prevDepr
		2140,  // prevSGA
		2716,  // prevLiabilities
		2011,  // prevTLTD
		2956,  // currNetReceivables
		6580,  // currAssets
		8377,  // currCogs
		2409,  // currPPE
		11203, // currTotalAssets
		11682, // currSales
		386,   // currDepr
		1310,  // currSGA
		1981,  // currLiabilities
		1946,  // currTLTD
		1094,  // currICO
		-82,   // currCFO
	)

	dsriExpected := decimal.NewFromFloat(1.1670)
	if !mscore.Dsri.Equals(dsriExpected) {
		t.Errorf("expected dsri of %q, got %q", dsriExpected, mscore.Dsri)
	}

	gmiExpected := decimal.NewFromFloat(0.3002)
	if !mscore.Gmi.Equals(gmiExpected) {
		t.Errorf("expected gmi of %q, got %q", gmiExpected, mscore.Gmi)
	}

	aqiExpected := decimal.NewFromFloat(0.9282)
	if !mscore.Aqi.Equals(aqiExpected) {
		t.Errorf("expected aqi of %q, got %q", aqiExpected, mscore.Aqi)
	}

	sgiExpected := decimal.NewFromFloat(1.1870)
	if !mscore.Sgi.Equals(sgiExpected) {
		t.Errorf("expected sgi of %q, got %q", sgiExpected, mscore.Sgi)
	}

	sgaiExpected := decimal.NewFromFloat(0.5157)
	if !mscore.Sgai.Equals(sgaiExpected) {
		t.Errorf("expected sgai of %q, got %q", sgaiExpected, mscore.Sgai)
	}

	depiExpected := decimal.NewFromFloat(1.2831)
	if !mscore.Depi.Equals(depiExpected) {
		t.Errorf("expected depi of %q, got %q", depiExpected, mscore.Depi)
	}

	tataExpected := decimal.NewFromFloat(0.1050)
	if !mscore.Tata.Equals(tataExpected) {
		t.Errorf("expected tata of %q, got %q", tataExpected, mscore.Tata)
	}

	lvgiExpected := decimal.NewFromFloat(0.7955)
	if !mscore.Lvgi.Equals(lvgiExpected) {
		t.Errorf("expected lvgi of %q, got %q", lvgiExpected, mscore.Lvgi)
	}

	mExpected := decimal.NewFromFloat(-1.8840)

	if !mscore.Score.Equals(mExpected) {
		fmt.Println(mscore.Score)
		t.Errorf("expected %q, got %q", mExpected, mscore.Score)
	}
}
