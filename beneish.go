package beneish

import (
	"github.com/shopspring/decimal"
)

/*
https://en.wikipedia.org/wiki/Beneish_M-score
The Beneish M-score is calculated using 8 variables (financial ratios):[1][2]

The formula to calculate the M-score is:[1]

M-score = −4.84 + 0.92 × DSRI + 0.528 × GMI + 0.404 × AQI + 0.892 × SGI + 0.115 × DEPI −0.172 × SGAI + 4.679 × TATA − 0.327 × LVGI
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
// AQI = [1 - (Current Assets t + PP&E t + Securities t) / Total Assets t] / = [1 - ((Current Assets t-1 + PP&E t-1 + Securities t-1) / Total Assets t-1)]
func aqi(
	prevAssets int64,
	prevSecurities int64,
	prevPPE int64,
	prevTotalAssets int64,
	currAssets int64,
	currSecurities int64,
	currPPE int64,
	currTotalAssets int64,
) decimal.Decimal {
	currSum := currAssets + currSecurities + currPPE
	currRatio := decimal.New(currSum, 0).Div(decimal.New(currTotalAssets, 0))
	prevSum := prevAssets + prevPPE + prevSecurities
	prevRatio := decimal.New(prevSum, 0).Div(decimal.New(prevTotalAssets, 0))
	numerator := decimal.New(1, 0).Sub(currRatio)
	denominator := decimal.New(1, 0).Sub(prevRatio)

	return numerator.DivRound(denominator, 4)
}

// Sales Growth Index (SGI)
// SGI = Sales t / Sales t-1

// Depreciation Index (DEPI)
// DEPI = (Depreciation t-1/ (PP&E t-1 + Depreciation t-1)) / (Depreciationt / (PP&Et + Depreciationt))

// Sales General and Administrative Expenses Index (SGAI)
// SGAI = (SG&A Expense t / Sales t) / (SG&A Expense t-1 / Sales t-1)

// Leverage Index (LVGI)
// LVGI = [(Current Liabilities t + Total Long Term Debt t) / Total Assets t] / [(Current Liabilities t-1 + Total Long Term Debt t-1) / Total Assets t-1]

// Total Accruals to Total Assets (TATA)
// TATA = (Income from Continuing Operations t - Cash Flows from Operations t) / Total Assets t
