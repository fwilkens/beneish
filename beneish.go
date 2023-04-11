package beneish

import (
	"github.com/shopspring/decimal"
)

type mscore struct {
	dsri  decimal.Decimal
	gmi   decimal.Decimal
	aqi   decimal.Decimal
	sgi   decimal.Decimal
	depi  decimal.Decimal
	sgai  decimal.Decimal
	lvgi  decimal.Decimal
	tata  decimal.Decimal
	score decimal.Decimal
}

/*
https://en.wikipedia.org/wiki/Beneish_M-score
The Beneish M-score is calculated using 8 variables (financial ratios)
*/

// Days Sales in Receivables Index
// (DSRI) DSRI = (Net Receivables t / Sales t) / (Net Receivables t-1 / Sales t-1)
func dsri(prevNetReceivables int64, currNetReceivables int64, prevSales int64, currSales int64) decimal.Decimal {
	numerator := decimal.New(currNetReceivables, 0).Div(decimal.New(currSales, 0))
	denominator := decimal.New(prevNetReceivables, 0).Div(decimal.New(prevSales, 0))

	return numerator.DivRound(denominator, 4)
}

// Gross Margin Index (GMI)
// GMI = [(Sales t-1 - COGS t-1) / Sales t-1] / [(Sales t - COGSt) / Sales t]
func gmi(prevSales int64, currSales int64, prevCogs int64, currCogs int64) decimal.Decimal {
	numerator := decimal.New(prevSales-prevCogs, 0).Div(decimal.New(prevSales, 0))
	denominator := decimal.New(currSales-currCogs, 0).Div(decimal.New(currSales, 0))

	return numerator.DivRound(denominator, 4)
}

// Asset Quality Index (AQI)
// AQI = [1 - (Current Assets t + PP&E t) / Total Assets t] / = [1 - ((Current Assets t-1 + PP&E t-1) / Total Assets t-1)]
func aqi(
	prevAssets int64,
	prevPPE int64,
	prevTotalAssets int64,
	currAssets int64,
	currPPE int64,
	currTotalAssets int64,
) decimal.Decimal {
	currSum := currAssets + currPPE
	currIndex := decimal.NewFromInt(1).Sub(decimal.NewFromInt(currSum).Div(decimal.NewFromInt(currTotalAssets)))
	prevSum := prevAssets + prevPPE
	prevIndex := decimal.NewFromInt(1).Sub(decimal.NewFromInt(prevSum).Div(decimal.NewFromInt(prevTotalAssets)))

	return currIndex.DivRound(prevIndex, 4)
}

// Sales Growth Index (SGI)
// SGI = Sales t / Sales t-1
func sgi(prevSales int64, currSales int64) decimal.Decimal {
	return decimal.New(currSales, 0).DivRound(decimal.New(prevSales, 0), 4)
}

// Depreciation Index (DEPI)
// DEPI = (Depreciation t-1/ (PP&E t-1 + Depreciation t-1)) / (Depreciationt / (PP&Et + Depreciationt))
func depi(prevDepr int64, prevPPE int64, currDepr int64, currPPE int64) decimal.Decimal {
	numerator := decimal.New(prevDepr, 0).Div(decimal.New(prevPPE+prevDepr, 0))
	denominator := decimal.New(currDepr, 0).Div(decimal.New(currPPE+currDepr, 0))

	return numerator.DivRound(denominator, 4)
}

// Sales General and Administrative Expenses Index (SGAI)
// SGAI = (SG&A Expense t / Sales t) / (SG&A Expense t-1 / Sales t-1)
func sgai(prevSGA int64, prevSales int64, currSGA int64, currSales int64) decimal.Decimal {
	numerator := decimal.New(currSGA, 0).Div(decimal.New(currSales, 0))
	denominator := decimal.New(prevSGA, 0).Div(decimal.New(prevSales, 0))

	return numerator.DivRound(denominator, 4)
}

// Leverage Index (LVGI)
// LVGI = [(Current Liabilities t + Total Long Term Debt t) / Total Assets t] / [(Current Liabilities t-1 + Total Long Term Debt t-1) / Total Assets t-1]
func lvgi(
	prevLiabilities int64,
	prevTLTD int64,
	prevTotalAssets int64,
	currLiabilities int64,
	currTLTD int64,
	currTotalAssets int64,
) decimal.Decimal {
	numerator := decimal.New(currLiabilities+currTLTD, 0).Div(decimal.New(currTotalAssets, 0))
	denominator := decimal.New(prevLiabilities+prevTLTD, 0).Div(decimal.New(prevTotalAssets, 0))

	return numerator.DivRound(denominator, 4)
}

// Total Accruals to Total Assets (TATA)
// TATA = (Income from Continuing Operations t - Cash Flows from Operations t) / Total Assets t
func tata(currICO int64, currCFO int64, currTotalAssets int64) decimal.Decimal {
	return decimal.New(currICO+currCFO, 0).DivRound(decimal.New(currTotalAssets, 0), 4)
}

/*
M-score =
−4.84 +
0.92 × DSRI +
0.528 × GMI +
0.404 × AQI +
0.892 × SGI +
0.115 × DEPI −
0.172 × SGAI +
4.679 × TATA −
0.327 × LVGI
*/
const BaseMScore = -4.84
const DSRImod = 0.92
const GMImod = 0.528
const AQImod = 0.404
const SGImod = 0.892
const DEPImod = 0.115
const SGAImod = 0.172
const TATAmod = 4.679
const LVGImod = 0.327

func mScoreCalc(
	prevNetReceivables int64,
	prevAssets int64,
	prevCogs int64,
	prevPPE int64,
	prevTotalAssets int64,
	prevSales int64,
	prevDepr int64,
	prevSGA int64,
	prevLiabilities int64,
	prevTLTD int64,
	currNetReceivables int64,
	currAssets int64,
	currCogs int64,
	currPPE int64,
	currTotalAssets int64,
	currSales int64,
	currDepr int64,
	currSGA int64,
	currLiabilities int64,
	currTLTD int64,
	currICO int64,
	currCFO int64,
) *mscore {
	dsri := dsri(prevNetReceivables, currNetReceivables, prevSales, currSales)
	gmi := gmi(prevSales, currSales, prevCogs, currCogs)
	aqi := aqi(prevAssets, prevPPE, prevTotalAssets, currAssets, currPPE, currTotalAssets)
	sgi := sgi(prevSales, currSales)
	sgai := sgai(prevSGA, prevSales, currSGA, currSales)
	depi := depi(prevDepr, prevPPE, currDepr, currPPE)
	tata := tata(currICO, currCFO, currTotalAssets)
	lvgi := lvgi(prevLiabilities, prevTLTD, prevTotalAssets, currLiabilities, currTLTD, currTotalAssets)

	score := decimal.NewFromFloat(BaseMScore)
	score = score.Add(decimal.NewFromFloat(DSRImod).Mul(dsri))
	score = score.Add(decimal.NewFromFloat(GMImod).Mul(gmi))
	score = score.Add(decimal.NewFromFloat(AQImod).Mul(aqi))
	score = score.Add(decimal.NewFromFloat(SGImod).Mul(sgi))
	score = score.Sub(decimal.NewFromFloat(SGAImod).Mul(sgai))
	score = score.Add(decimal.NewFromFloat(DEPImod).Mul(depi))
	score = score.Add(decimal.NewFromFloat(TATAmod).Mul(tata))
	score = score.Sub(decimal.NewFromFloat(LVGImod).Mul(lvgi))

	m := mscore{
		dsri:  dsri,
		gmi:   gmi,
		aqi:   aqi,
		sgi:   sgi,
		depi:  depi,
		sgai:  sgai,
		lvgi:  lvgi,
		tata:  tata,
		score: score,
	}
	return &m
}
