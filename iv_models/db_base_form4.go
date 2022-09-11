package iv_models

type DB_BaseForm4 struct {
	Parsed                    bool                          `json:"parsed" bson:"parsed"`
	PeriodOfReport            string                        `json:"periodOfReport" bson:"periodOfReport"`
	AccessionNumber           string                        `json:"accessionNumber" bson:"accessionNumber"`
	Url                       string                        `json:"url" bson:"url"`
	Issuer                    DB_Issuer                     `json:"issuer" bson:"issuer"`
	Reporters                 []DB_Reporter                 `json:"reporters" bson:"reporters"`
	DerivativeTransactions    []DB_DerivativeTransaction    `json:"derivativeTransactions" bson:"derivativeTransactions"`
	NonDerivativeTransactions []DB_NonDerivativeTransaction `json:"nonDerivativeTransactions" bson:"nonDerivativeTransactions"`
}

type DB_NonDerivativeTransaction struct {
	SecurityTitle                   string  `json:"securityTitle" bson:"securityTitle"`
	TransactionDate                 string  `json:"transactionDate" bson:"transactionDate"`
	TransactionCode                 string  `json:"transactionCode" bson:"transactionCode"`
	TransactionShares               float32 `json:"transactionShares" bson:"transactionShares"`
	TransactionPricePerShare        float32 `json:"transactionPricePerShare" bson:"transactionPricePerShare"`
	TransactionAcquiredDisposedCode string  `json:"transactionAcquiredDisposedCode" bson:"transactionAcquiredDisposedCode"`
	PostTransactionShares           float32 `json:"postTransactionShares" bson:"postTransactionShares"`
}

type DB_DerivativeTransaction struct {
	SecurityTitle                   string  `json:"securityTitle" bson:"securityTitle"`
	TransactionDate                 string  `json:"transactionDate" bson:"transactionDate"`
	ConversionOrExercisePrice       float32 `json:"conversionOrExercisePrice" bson:"conversionOrExercisePrice"`
	DeemedExecutionDate             string  `json:"deemedExecutionDate" bson:"deemedExecutionDate"`
	TransactionCode                 string  `json:"transactionCode" bson:"transactionCode"`
	TransactionShares               float32 `json:"transactionShares" bson:"transactionShares"`
	TransactionPricePerShare        float32 `json:"transactionPricePerShare" bson:"transactionPricePerShare"`
	TransactionAcquiredDisposedCode string  `json:"transactionAcquiredDisposedCode" bson:"transactionAcquiredDisposedCode"`
	ExerciseDate                    string  `json:"exerciseDate" bson:"exerciseDate"`
	PostTransactionShares           float32 `json:"postTransactionShares" bson:"postTransactionShares"`
	ExpirationDate                  string  `json:"expirationDate" bson:"expirationDate"`
}
