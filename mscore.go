package beneish

import (
	"github.com/shopspring/decimal"
)

type Mscore struct {
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

// https://en.wikipedia.org/wiki/Beneish_M-score
// M-score = −4.84 + 0.92 × DSRI + 0.528 × GMI + 0.404 × AQI + 0.892 × SGI + 0.115 × DEPI − 0.172 × SGAI + 4.679 × TATA − 0.327 × LVGI

// Days Sales in Receivables Index
// (DSRI) DSRI = (Net Receivables t / Sales t) / (Net Receivables t-1 / Sales t-1)
func Dsri(prevNetReceivables int64, currNetReceivables int64, prevSales int64, currSales int64) decimal.Decimal {
	numerator := decimal.NewFromInt(currNetReceivables).Div(decimal.NewFromInt(currSales))
	denominator := decimal.NewFromInt(prevNetReceivables).Div(decimal.NewFromInt(prevSales))

	return numerator.DivRound(denominator, 4)
}

// Gross Margin Index (GMI)
// GMI = [(Sales t-1 - COGS t-1) / Sales t-1] / [(Sales t - COGSt) / Sales t]
func Gmi(prevSales int64, currSales int64, prevCogs int64, currCogs int64) decimal.Decimal {
	numerator := decimal.NewFromInt(prevSales - prevCogs).Div(decimal.NewFromInt(prevSales))
	denominator := decimal.NewFromInt(currSales - currCogs).Div(decimal.NewFromInt(currSales))

	return numerator.DivRound(denominator, 4)
}

// Asset Quality Index (AQI)
// AQI = [1 - (Current Assets t + PP&E t) / Total Assets t] / = [1 - ((Current Assets t-1 + PP&E t-1) / Total Assets t-1)]
func Aqi(
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
func Sgi(prevSales int64, currSales int64) decimal.Decimal {
	return decimal.NewFromInt(currSales).DivRound(decimal.NewFromInt(prevSales), 4)
}

// Depreciation Index (DEPI)
// DEPI = (Depreciation t-1/ (PP&E t-1 + Depreciation t-1)) / (Depreciationt / (PP&Et + Depreciationt))
func Depi(prevDepr int64, prevPPE int64, currDepr int64, currPPE int64) decimal.Decimal {
	numerator := decimal.NewFromInt(prevDepr).Div(decimal.NewFromInt(prevPPE + prevDepr))
	denominator := decimal.NewFromInt(currDepr).Div(decimal.NewFromInt(currPPE + currDepr))

	return numerator.DivRound(denominator, 4)
}

// Sales General and Administrative Expenses Index (SGAI)
// SGAI = (SG&A Expense t / Sales t) / (SG&A Expense t-1 / Sales t-1)
func Sgai(prevSGA int64, prevSales int64, currSGA int64, currSales int64) decimal.Decimal {
	numerator := decimal.NewFromInt(currSGA).Div(decimal.NewFromInt(currSales))
	denominator := decimal.NewFromInt(prevSGA).Div(decimal.NewFromInt(prevSales))

	return numerator.DivRound(denominator, 4)
}

// Leverage Index (LVGI)
// LVGI = [(Current Liabilities t + Total Long Term Debt t) / Total Assets t] / [(Current Liabilities t-1 + Total Long Term Debt t-1) / Total Assets t-1]
func Lvgi(
	prevLiabilities int64,
	prevTLTD int64,
	prevTotalAssets int64,
	currLiabilities int64,
	currTLTD int64,
	currTotalAssets int64,
) decimal.Decimal {
	numerator := decimal.NewFromInt(currLiabilities + currTLTD).Div(decimal.NewFromInt(currTotalAssets))
	denominator := decimal.NewFromInt(prevLiabilities + prevTLTD).Div(decimal.NewFromInt(prevTotalAssets))

	return numerator.DivRound(denominator, 4)
}

// Total Accruals to Total Assets (TATA)
// TATA = (Income from Continuing Operations t - Cash Flows from Operations t) / Total Assets t
func Tata(currICO int64, currCFO int64, currTotalAssets int64) decimal.Decimal {
	return decimal.NewFromInt(currICO-currCFO).DivRound(decimal.NewFromInt(currTotalAssets), 4)
}

const BaseMScore = -4.84
const DSRIcoef = 0.92
const GMIcoef = 0.528
const AQIcoef = 0.404
const SGIcoef = 0.892
const DEPIcoef = 0.115
const SGAIcoef = 0.172
const TATAcoef = 4.679
const LVGIcoef = 0.327

func MscoreCalc(
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
) *Mscore {
	dsri := Dsri(prevNetReceivables, currNetReceivables, prevSales, currSales)
	gmi := Gmi(prevSales, currSales, prevCogs, currCogs)
	aqi := Aqi(prevAssets, prevPPE, prevTotalAssets, currAssets, currPPE, currTotalAssets)
	sgi := Sgi(prevSales, currSales)
	sgai := Sgai(prevSGA, prevSales, currSGA, currSales)
	depi := Depi(prevDepr, prevPPE, currDepr, currPPE)
	tata := Tata(currICO, currCFO, currTotalAssets)
	lvgi := Lvgi(prevLiabilities, prevTLTD, prevTotalAssets, currLiabilities, currTLTD, currTotalAssets)

	score := decimal.NewFromFloat(BaseMScore)
	score = score.Add(decimal.NewFromFloat(DSRIcoef).Mul(dsri))
	score = score.Add(decimal.NewFromFloat(GMIcoef).Mul(gmi))
	score = score.Add(decimal.NewFromFloat(AQIcoef).Mul(aqi))
	score = score.Add(decimal.NewFromFloat(SGIcoef).Mul(sgi))
	score = score.Sub(decimal.NewFromFloat(SGAIcoef).Mul(sgai))
	score = score.Add(decimal.NewFromFloat(DEPIcoef).Mul(depi))
	score = score.Add(decimal.NewFromFloat(TATAcoef).Mul(tata))
	score = score.Sub(decimal.NewFromFloat(LVGIcoef).Mul(lvgi)).Round(4)

	m := Mscore{
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
