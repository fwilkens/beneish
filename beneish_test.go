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

func TestMscore(t *testing.T) {
	mscore := mScoreCalc(
		9692,   // prevNetReceivables
		60197,  // prevAssets
		103134, // prevCogs
		737,    // prevSecurities
		68573,  // prevPPE
		131310, // prevTotalAssets
		177866, // prevSales
		9166,   // prevDepr
		61612,  // prevSGA
		57883,  // prevLiabilities
		37926,  // prevTLTD
		13310,  // currNetReceivables
		75101,  // currAssets
		127056, // currCogs
		942,    // currSecurities
		95770,  // currPPE
		162648, // currTotalAssets
		232887, // currSales
		12575,  // currDepr
		81014,  // currSGA
		63391,  // currLiabilities
		39787,  // currTLTD
		10073,  // currICO
		30723,  // currCFO
	)

	dsriExpected, _ := decimal.NewFromString("1.0488")
	if !mscore.dsri.Equals(dsriExpected) {
		t.Errorf("expected dsri of %q, got %q", dsriExpected, mscore.dsri)
	}

	gmiExpected, _ := decimal.NewFromString("0.9246")
	if !mscore.gmi.Equals(gmiExpected) {
		t.Errorf("expected gmi of %q, got %q", gmiExpected, mscore.gmi)
	}

	aqiExpected, _ := decimal.NewFromString("1.07")
	if !mscore.aqi.Equals(aqiExpected) {
		t.Errorf("expected aqi of %q, got %q", aqiExpected, mscore.aqi)
	}

	sgiExpected, _ := decimal.NewFromString("1.3093")
	if !mscore.sgi.Equals(sgiExpected) {
		t.Errorf("expected sgi of %q, got %q", sgiExpected, mscore.sgi)
	}

	sgaiExpected, _ := decimal.NewFromString("1.0043")
	if !mscore.sgai.Equals(sgaiExpected) {
		t.Errorf("expected sgai of %q, got %q", sgaiExpected, mscore.sgai)
	}

	depiExpected, _ := decimal.NewFromString("1.0159")
	if !mscore.depi.Equals(depiExpected) {
		t.Errorf("expected depi of %q, got %q", depiExpected, mscore.depi)
	}

	tataExpected, _ := decimal.NewFromString("-0.13")
	if !mscore.tata.Equals(tataExpected) {
		t.Errorf("expected tata of %q, got %q", tataExpected, mscore.tata)
	}

	lvgiExpected, _ := decimal.NewFromString("0.91")
	if !mscore.lvgi.Equals(lvgiExpected) {
		t.Errorf("expected lvgi of %q, got %q", lvgiExpected, mscore.lvgi)
	}

	mExpected, _ := decimal.NewFromString("-2.734")

	if !mscore.score.Equals(mExpected) {
		fmt.Println(mscore.score)
		t.Errorf("expected %q, got %q", mExpected, mscore.score)
	}
}
