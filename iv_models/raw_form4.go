package iv_models

import (
	"encoding/xml"
)

type RawForm4 struct {
	XMLName                  xml.Name           `xml:"ownershipDocument" json:"-"`
	SchemaVersion            string             `xml:"schemaVersion"`
	DocumentType             string             `xml:"documentType"`
	PeriodOfReport           string             `xml:"periodOfReport"`
	DateOfOriginalSubmission string             `xml:"dateOfOriginalSubmission"`
	NoSecuritiesOwned        bool               `xml:"noSecuritiesOwned"`
	AccessionNumber          string             `xml:"accessionNumber"`
	Url                      string             `xml:"url"`
	Issuer                   Issuer             `xml:"issuer"`
	ReportingOwners          []ReportingOwner   `xml:"reportingOwner" json:"ReportingOwners"`
	NonDerivativeTable       NonDerivativeTable `xml:"nonDerivativeTable"`
	DerivativeTable          DerivativeTable    `xml:"derivativeTable"`
	OwnerSignatures          []OwnerSignature   `xml:"ownerSignature" `
}

type Issuer struct {
	XMLName             xml.Name `xml:"issuer" json:"-"`
	IssuerCIK           string   `xml:"issuerCik"`
	IssuerName          string   `xml:"issuerName"`
	IssuerTradingSymbol string   `xml:"issuerTradingSymbol"`
}

type ReportingOwner struct {
	XMLName                    xml.Name                   `xml:"reportingOwner"`
	ReportingOwnerId           ReportingOwnerId           `xml:"reportingOwnerId"`
	ReportingOwnerAddress      ReportingOwnerAddress      `xml:"reportingOwnerAddress" `
	ReportingOwnerRelationship ReportingOwnerRelationship `xml:"reportingOwnerRelationship" `
}

type ReportingOwnerId struct {
	XMLName      xml.Name `xml:"reportingOwnerId"`
	RptOwnerCik  string   `xml:"rptOwnerCik"`
	RptOwnerCcc  string   `xml:"rptOwnerCcc"`
	RptOwnerName string   `xml:"rptOwnerName"`
}

type ReportingOwnerAddress struct {
	XMLName                  xml.Name `xml:"reportingOwnerAddress"`
	RptOwnerStreet1          string   `xml:"rptOwnerStreet1"`
	RptOwnerStreet2          string   `xml:"rptOwnerStreet2"`
	RptOwnerCity             string   `xml:"rptOwnerCity"`
	RptOwnerState            string   `xml:"rptOwnerState"`
	RptOwnerZipCode          string   `xml:"rptOwnerZipCode"`
	RptOwnerStateDescription string   `xml:"rptOwnerStateDescription"`
}

type ReportingOwnerRelationship struct {
	XMLName           xml.Name `xml:"reportingOwnerRelationship"`
	IsDirector        bool     `xml:"isDirector"`
	IsOfficer         bool     `xml:"isOfficer"`
	IsTenPercentOwner bool     `xml:"isTenPercentOwner"`
	IsOther           bool     `xml:"isOther"`
	OfficerTitle      string   `xml:"officerTitle"`
	OtherText         string   `xml:"otherText"`
}

type NonDerivativeTable struct {
	XMLName                   xml.Name                   `xml:"nonDerivativeTable"`
	NonDerivativeTransactions []NonDerivativeTransaction `xml:"nonDerivativeTransaction" `
	NonDerivativeHoldings     []NonDerivativeHolding     `xml:"nonDerivativeHolding"`
}

type NonDerivativeTransaction struct {
	XMLName                xml.Name               `xml:"nonDerivativeTransaction"`
	SecurityTitle          SecurityTitle          `xml:"securityTitle" `
	TransactionDate        TransactionDate        `xml:"transactionDate" `
	DeemedExecutionDate    DeemedExecutionDate    `xml:"deemedExecutionDate" `
	TransactionCoding      TransactionCoding      `xml:"transactionCoding" `
	TransactionTimeliness  TransactionTimeliness  `xml:"transactionTimeliness" `
	TransactionAmounts     TransactionAmounts     `xml:"transactionAmounts" `
	PostTransactionAmounts PostTransactionAmounts `xml:"postTransactionAmounts" `
	OwnershipNature        OwnershipNature        `xml:"ownershipNature" `
}

type NonDerivativeHolding struct {
	XMLName                xml.Name               `xml:"nonDerivativeHolding"`
	SecurityTitle          SecurityTitle          `xml:"securityTitle" `
	PostTransactionAmounts PostTransactionAmounts `xml:"postTransactionAmounts" `
	OwnershipNature        OwnershipNature        `xml:"ownershipNature" `
}

type DerivativeTable struct {
	XMLName                xml.Name                `xml:"derivativeTable"`
	DerivativeTransactions []DerivativeTransaction `xml:"derivativeTransaction" `
	DerivativeHoldings     []DerivativeHolding     `xml:"derivativeHolding" `
}

type DerivativeTransaction struct {
	XMLName                   xml.Name `xml:"derivativeTransaction"`
	DerivativeTableID         uint
	SecurityTitle             SecurityTitle             `xml:"securityTitle"`
	ConversionOrExercisePrice ConversionOrExercisePrice `xml:"conversionOrExercisePrice" `
	TransactionDate           TransactionDate           `xml:"transactionDate"`
	DeemedExecutionDate       DeemedExecutionDate       `xml:"deemedExecutionDate" `
	TransactionCoding         TransactionCoding         `xml:"transactionCoding"`
	TransactionTimeliness     TransactionTimeliness     `xml:"transactionTimeliness"`
	TransactionAmounts        TransactionAmounts        `xml:"transactionAmounts"`
	ExerciseDate              ExerciseDate              `xml:"exerciseDate"`
	ExpirationDate            ExpirationDate            `xml:"expirationDate" `
	UnderlyingSecurity        UnderlyingSecurity        `xml:"underlyingSecurity"`
	PostTransactionAmounts    PostTransactionAmounts    `xml:"postTransactionAmounts" `
	OwnershipNature           OwnershipNature           `xml:"ownershipNature"`
}

type DerivativeHolding struct {
	XMLName                   xml.Name `xml:"derivativeHolding"`
	DerivativeTableID         uint
	SecurityTitle             SecurityTitle             `xml:"securityTitle" `
	ConversionOrExercisePrice ConversionOrExercisePrice `xml:"conversionOrExercisePrice" `
	ExerciseDate              ExerciseDate              `xml:"exerciseDate" `
	ExpirationDate            ExpirationDate            `xml:"expirationDate" `
	UnderlyingSecurity        UnderlyingSecurity        `xml:"underlyingSecurity" `
	PostTransactionAmounts    PostTransactionAmounts    `xml:"postTransactionAmounts" `
	OwnershipNature           OwnershipNature           `xml:"ownershipNature" `
}

type UnderlyingSecurity struct {
	XMLName                  xml.Name                 `xml:"underlyingSecurity"`
	UnderlyingSecurityTitle  UnderlyingSecurityTitle  `xml:"underlyingSecurityTitle" `
	UnderlyingSecurityShares UnderlyingSecurityShares `xml:"underlyingSecurityShares" `
	UnderlyingSecurityValue  UnderlyingSecurityValue  `xml:"underlyingSecurityValue" `
}

type UnderlyingSecurityShares struct {
	XMLName xml.Name `xml:"underlyingSecurityShares"`
	Value   string   `xml:"value"`
}

type UnderlyingSecurityValue struct {
	XMLName xml.Name `xml:"underlyingSecurityValue"`
	Value   string   `xml:"value"`
}

type UnderlyingSecurityTitle struct {
	XMLName xml.Name `xml:"underlyingSecurityTitle"`
	Value   string   `xml:"value"`
}

type ExpirationDate struct {
	XMLName xml.Name `xml:"expirationDate"`
	Value   string   `xml:"value"`
}

type ExerciseDate struct {
	XMLName xml.Name `xml:"exerciseDate"`
	Value   string   `xml:"value"`
}

type ConversionOrExercisePrice struct {
	XMLName xml.Name `xml:"conversionOrExercisePrice"`
	Value   float32  `xml:"value"`
}

type SecurityTitle struct {
	XMLName xml.Name `xml:"securityTitle"`
	Value   string   `xml:"value"`
}

type TransactionDate struct {
	XMLName xml.Name `xml:"transactionDate"`
	Value   string   `xml:"value"`
}

type DeemedExecutionDate struct {
	XMLName xml.Name `xml:"deemedExecutionDate"`
	Value   string   `xml:"value"`
}

type TransactionCoding struct {
	XMLName             xml.Name `xml:"transactionCoding"`
	TransactionFormType string   `xml:"transactionFormType" `
	TransactionCode     string   `xml:"transactionCode"`
	EquitySwapInvolved  bool     `xml:"equitySwapInvolved" `
}

type TransactionTimeliness struct {
	XMLName xml.Name `xml:"transactionTimeliness"`
	Value   string   `xml:"value"`
}

type TransactionAmounts struct {
	XMLName                         xml.Name                        `xml:"transactionAmounts"`
	TransactionShares               TransactionShares               `xml:"transactionShares" `
	TransactionTotalValue           TransactionTotalValue           `xml:"transactionTotalValue" `
	TransactionPricePerShare        TransactionPricePerShare        `xml:"transactionPricePerShare" `
	TransactionAcquiredDisposedCode TransactionAcquiredDisposedCode `xml:"transactionAcquiredDisposedCode" `
}

type TransactionShares struct {
	XMLName xml.Name `xml:"transactionShares"`
	Value   float32  `xml:"value"`
}

type TransactionTotalValue struct {
	XMLName xml.Name `xml:"transactionTotalValue"`
	Value   string   `xml:"value"`
}

type TransactionPricePerShare struct {
	XMLName xml.Name `xml:"transactionPricePerShare"`
	Value   float32  `xml:"value"`
}

type TransactionAcquiredDisposedCode struct {
	XMLName xml.Name `xml:"transactionAcquiredDisposedCode"`
	Value   string   `xml:"value"`
}

type PostTransactionAmounts struct {
	XMLName                         xml.Name                        `xml:"postTransactionAmounts"`
	SharesOwnedFollowingTransaction SharesOwnedFollowingTransaction `xml:"sharesOwnedFollowingTransaction" `
	ValueOwnedFollowingTransaction  ValueOwnedFollowingTransaction  `xml:"valueOwnedFollowingTransaction" `
}

type SharesOwnedFollowingTransaction struct {
	XMLName xml.Name `xml:"sharesOwnedFollowingTransaction"`
	Value   float32  `xml:"value"`
}

type ValueOwnedFollowingTransaction struct {
	XMLName xml.Name `xml:"valueOwnedFollowingTransaction"`
	Value   string   `xml:"value"`
}

type OwnershipNature struct {
	XMLName                   xml.Name                  `xml:"ownershipNature"`
	DirectOrIndirectOwnership DirectOrIndirectOwnership `xml:"directOrIndirectOwnership" `
	NatureOfOwnership         NatureOfOwnership         `xml:"natureOfOwnership" `
}

type DirectOrIndirectOwnership struct {
	XMLName xml.Name `xml:"directOrIndirectOwnership"`
	Value   string   `xml:"value"`
}

type NatureOfOwnership struct {
	XMLName xml.Name `xml:"natureOfOwnership"`
	Value   string   `xml:"value"`
}

type OwnerSignature struct {
	XMLName       xml.Name `xml:"ownerSignature"`
	SignatureName string   `xml:"signatureName"`
	SignatureDate string   `xml:"signatureDate"`
}
