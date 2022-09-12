package insidervizforms

import "github.com/Matterhorn-Studios/insidervizforms/iv_models"

func RawForm4ToDbForm4(rawForm4 iv_models.RawForm4) iv_models.DB_BaseForm4 {
	// create the db form4
	dbForm4 := iv_models.DB_BaseForm4{}

	// main parts
	dbForm4.Parsed = false
	dbForm4.PeriodOfReport = rawForm4.PeriodOfReport
	dbForm4.Issuer.IssuerCik = rawForm4.Issuer.IssuerCIK
	dbForm4.Issuer.IssuerName = rawForm4.Issuer.IssuerName
	dbForm4.Issuer.IssuerTicker = rawForm4.Issuer.IssuerTradingSymbol
	dbForm4.AccessionNumber = rawForm4.AccessionNumber
	dbForm4.Url = rawForm4.Url

	// reporters
	dbForm4.Reporters = make([]iv_models.DB_Reporter, 0)
	for _, reporter := range rawForm4.ReportingOwners {
		curReporter := iv_models.DB_Reporter{}
		curReporter.ReporterName = reporter.ReportingOwnerId.RptOwnerName
		curReporter.ReporterCik = reporter.ReportingOwnerId.RptOwnerCik
		// title
		ownerRelation := reporter.ReportingOwnerRelationship
		title := ""
		if ownerRelation.OfficerTitle != "" {
			title += ", " + ownerRelation.OfficerTitle
		}
		if ownerRelation.IsDirector {
			title += ", Director"
		}
		if ownerRelation.IsOfficer {
			title += ", Officer"
		}
		if ownerRelation.IsTenPercentOwner {
			title += ", 10% Owner"
		}
		if ownerRelation.IsOther {
			title += ", " + ownerRelation.OtherText
		}

		// remove the first 2 characters from title
		title = title[2:]
		curReporter.ReporterTitle = title

		// address
		ownerAddress := reporter.ReportingOwnerAddress
		address := ownerAddress.RptOwnerStreet1 + " " + ownerAddress.RptOwnerCity + ", " + ownerAddress.RptOwnerState + " " + ownerAddress.RptOwnerZipCode
		curReporter.ReporterAddress = address

		dbForm4.Reporters = append(dbForm4.Reporters, curReporter)
	}

	// Derivative Transactions
	derivTransactions := make([]iv_models.DB_DerivativeTransaction, 0)
	for _, transaction := range rawForm4.DerivativeTable.DerivativeTransactions {
		curDerivTransaction := iv_models.DB_DerivativeTransaction{}
		curDerivTransaction.SecurityTitle = transaction.SecurityTitle.Value
		curDerivTransaction.TransactionDate = transaction.TransactionDate.Value
		curDerivTransaction.ConversionOrExercisePrice = transaction.ConversionOrExercisePrice.Value
		curDerivTransaction.DeemedExecutionDate = transaction.DeemedExecutionDate.Value
		curDerivTransaction.TransactionCode = transaction.TransactionCoding.TransactionCode
		curDerivTransaction.TransactionShares = transaction.TransactionAmounts.TransactionShares.Value
		curDerivTransaction.TransactionPricePerShare = transaction.TransactionAmounts.TransactionPricePerShare.Value
		curDerivTransaction.ExerciseDate = transaction.ExerciseDate.Value
		curDerivTransaction.PostTransactionShares = transaction.PostTransactionAmounts.SharesOwnedFollowingTransaction.Value
		curDerivTransaction.ExpirationDate = transaction.ExpirationDate.Value
		curDerivTransaction.TransactionAcquiredDisposedCode = transaction.TransactionAmounts.TransactionAcquiredDisposedCode.Value
		derivTransactions = append(derivTransactions, curDerivTransaction)
	}
	dbForm4.DerivativeTransactions = derivTransactions

	// NonDerivative Transactions
	nonDerivTransactions := make([]iv_models.DB_NonDerivativeTransaction, 0)
	for _, transaction := range rawForm4.NonDerivativeTable.NonDerivativeTransactions {
		curNonDerivTransaction := iv_models.DB_NonDerivativeTransaction{}
		curNonDerivTransaction.SecurityTitle = transaction.SecurityTitle.Value
		curNonDerivTransaction.TransactionDate = transaction.TransactionDate.Value
		curNonDerivTransaction.TransactionCode = transaction.TransactionCoding.TransactionCode
		curNonDerivTransaction.TransactionShares = transaction.TransactionAmounts.TransactionShares.Value
		curNonDerivTransaction.TransactionPricePerShare = transaction.TransactionAmounts.TransactionPricePerShare.Value
		curNonDerivTransaction.TransactionAcquiredDisposedCode = transaction.TransactionAmounts.TransactionAcquiredDisposedCode.Value
		curNonDerivTransaction.PostTransactionShares = transaction.PostTransactionAmounts.SharesOwnedFollowingTransaction.Value
		nonDerivTransactions = append(nonDerivTransactions, curNonDerivTransaction)
	}
	dbForm4.NonDerivativeTransactions = nonDerivTransactions

	return dbForm4
}
