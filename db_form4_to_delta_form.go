package insidervizforms

import "github.com/Matterhorn-Studios/insidervizforms/iv_models"

func DbForm4ToDeltaForm(dbForm4 iv_models.DB_BaseForm4) (iv_models.DB_DeltaForm, bool) {
	flagged := false

	// create the delta form
	deltaForm := iv_models.DB_DeltaForm{}

	// add basic fields
	deltaForm.AccessionNumber = dbForm4.AccessionNumber
	deltaForm.FormClass = "Insider"
	deltaForm.PeriodOfReport = dbForm4.PeriodOfReport
	deltaForm.Url = dbForm4.Url
	deltaForm.Issuer = dbForm4.Issuer
	deltaForm.Reporters = dbForm4.Reporters

	// compute helper fields
	deltaForm.BuyOrSell = "Buy"
	count := 0
	rawNetTotal := 0.0
	rawAmount := 0.0
	curLastDate := ""
	for i := range dbForm4.NonDerivativeTransactions {
		curTransaction := dbForm4.NonDerivativeTransactions[i]
		if curTransaction.TransactionCode == "S" || curTransaction.TransactionCode == "P" {
			count++
			if curLastDate < curTransaction.TransactionDate {
				curLastDate = curTransaction.TransactionDate
				deltaForm.PostTransactionShares = curTransaction.PostTransactionShares
			}
			if curTransaction.TransactionAcquiredDisposedCode == "D" {
				// sell
				deltaForm.NetTotal -= curTransaction.TransactionShares * curTransaction.TransactionPricePerShare
				deltaForm.SharesTraded -= curTransaction.TransactionShares
			} else {
				// buy
				deltaForm.NetTotal += curTransaction.TransactionShares * curTransaction.TransactionPricePerShare
				deltaForm.SharesTraded += curTransaction.TransactionShares
			}
			rawNetTotal += float64(curTransaction.TransactionShares) * float64(curTransaction.TransactionPricePerShare)
			rawAmount += float64(curTransaction.TransactionShares)
		}
	}

	// cleanup and flag
	if deltaForm.NetTotal < 0 {
		deltaForm.BuyOrSell = "Sell"
		deltaForm.NetTotal = -deltaForm.NetTotal
	}
	if deltaForm.SharesTraded < 0 {
		deltaForm.SharesTraded = -deltaForm.SharesTraded
	}
	if count > 0 && rawAmount > 0 {
		deltaForm.AveragePricePerShare = float32(rawNetTotal / rawAmount)
	}
	if deltaForm.AveragePricePerShare > 5000 {
		flagged = true
	}
	if deltaForm.NetTotal == 0.0 {
		flagged = true
	}
	if deltaForm.NetTotal > 100000000000 {
		flagged = true
	}

	return deltaForm, flagged
}
