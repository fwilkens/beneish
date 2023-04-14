package beneish

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestDsri(t *testing.T) {
	result := dsri(90, 100, 150, 200)
	expected := decimal.NewFromFloat(0.8333)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestGmi(t *testing.T) {
	result := gmi(90, 100, 60, 50)
	expected := decimal.NewFromFloat(0.6667)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestAqi(t *testing.T) {
	result := aqi(100, 60, 600, 130, 60, 700)
	expected := decimal.NewFromFloat(0.9935)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestSgi(t *testing.T) {
	result := sgi(1200, 1300)
	expected := decimal.NewFromFloat(1.0833)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestDepi(t *testing.T) {
	result := depi(1000, 600, 1200, 900)
	expected := decimal.NewFromFloat(1.0938)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestLvgi(t *testing.T) {
	result := lvgi(1200, 400, 1800, 1300, 450, 1700)
	expected := decimal.NewFromFloat(1.1581)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestTata(t *testing.T) {
	result := tata(300, 50, 1200)
	expected := decimal.NewFromFloat(0.2083)

	if !result.Equals(expected) {
		fmt.Println(result)
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// Example data from:
// https://apps.kelley.iu.edu/Beneish/Content/Examples/SunbeamExample.pdf
func TestMscore(t *testing.T) {
	mscore := mScoreCalc(
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
	if !mscore.dsri.Equals(dsriExpected) {
		t.Errorf("expected dsri of %q, got %q", dsriExpected, mscore.dsri)
	}

	gmiExpected := decimal.NewFromFloat(0.3002)
	if !mscore.gmi.Equals(gmiExpected) {
		t.Errorf("expected gmi of %q, got %q", gmiExpected, mscore.gmi)
	}

	aqiExpected := decimal.NewFromFloat(0.9282)
	if !mscore.aqi.Equals(aqiExpected) {
		t.Errorf("expected aqi of %q, got %q", aqiExpected, mscore.aqi)
	}

	sgiExpected := decimal.NewFromFloat(1.1870)
	if !mscore.sgi.Equals(sgiExpected) {
		t.Errorf("expected sgi of %q, got %q", sgiExpected, mscore.sgi)
	}

	sgaiExpected := decimal.NewFromFloat(0.5157)
	if !mscore.sgai.Equals(sgaiExpected) {
		t.Errorf("expected sgai of %q, got %q", sgaiExpected, mscore.sgai)
	}

	depiExpected := decimal.NewFromFloat(1.2831)
	if !mscore.depi.Equals(depiExpected) {
		t.Errorf("expected depi of %q, got %q", depiExpected, mscore.depi)
	}

	tataExpected := decimal.NewFromFloat(0.1050)
	if !mscore.tata.Equals(tataExpected) {
		t.Errorf("expected tata of %q, got %q", tataExpected, mscore.tata)
	}

	lvgiExpected := decimal.NewFromFloat(0.7955)
	if !mscore.lvgi.Equals(lvgiExpected) {
		t.Errorf("expected lvgi of %q, got %q", lvgiExpected, mscore.lvgi)
	}

	mExpected := decimal.NewFromFloat(-1.8840)

	if !mscore.score.Equals(mExpected) {
		fmt.Println(mscore.score)
		t.Errorf("expected %q, got %q", mExpected, mscore.score)
	}
}
