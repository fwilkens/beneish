## Beneish Go package

A package to calculate [M scores](https://en.wikipedia.org/wiki/Beneish_M-score), as in Beneish, Lee Nichols 2013, Financial Analysts Journal.

## Usage

### Calculate an M-Score

```go
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
```

The returned struct will have the complete M score, as well as the incremental indexes.

There are functions for each incremental index as well. For example:

```go
dsri := dsri(prevNetReceivables, currNetReceivables, prevSales, currSales)
```

