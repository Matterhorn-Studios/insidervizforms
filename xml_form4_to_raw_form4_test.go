package insidervizforms_test

import (
	"encoding/xml"
	"testing"

	insidervizforms "github.com/Matterhorn-Studios/insidervizforms"
	"github.com/Matterhorn-Studios/insidervizforms/iv_models"
)

func Test1(t *testing.T) {
	actualForm, err := insidervizforms.XmlForm4ToRawForm4(
		"data/testdata_form4_1.xml",
		"0001366868-22-000065",
		"https://www.sec.gov/Archives/edgar/data/1026935/000136686822000065/wf-form4_166277274218582.xml",
	)

	if err != nil {
		t.Errorf("parse failed: %s", err.Error())
	}

	if actualForm.PeriodOfReport != "2022-09-09" {
		t.Errorf("period of report mismatch: %s", actualForm.PeriodOfReport)
	}

	if actualForm.AccessionNumber != "0001366868-22-000065" {
		t.Errorf("accession number mismatch: %s", actualForm.AccessionNumber)
	}

	expectedIssuer := iv_models.Issuer{
		XMLName:             xml.Name{Local: "issuer"},
		IssuerCIK:           "0001366868",
		IssuerName:          "Globalstar, Inc.",
		IssuerTradingSymbol: "GSAT",
	}
	if actualForm.Issuer != expectedIssuer {
		t.Errorf("issuer mismatch: \n%s\n%s", actualForm.Issuer, expectedIssuer)
	}

	expectedReporter := iv_models.ReportingOwner{
		XMLName: xml.Name{Local: "reportingOwner"},
		ReportingOwnerId: iv_models.ReportingOwnerId{
			XMLName:      xml.Name{Local: "reportingOwnerId"},
			RptOwnerCik:  "0001026935",
			RptOwnerCcc:  "",
			RptOwnerName: "COWAN KEITH O",
		},
		ReportingOwnerAddress: iv_models.ReportingOwnerAddress{
			XMLName:                  xml.Name{Local: "reportingOwnerAddress"},
			RptOwnerStreet1:          "1351 HOLIDAY SQUARE BLVD",
			RptOwnerStreet2:          "",
			RptOwnerCity:             "COVINGTON",
			RptOwnerState:            "LA",
			RptOwnerZipCode:          "70433",
			RptOwnerStateDescription: "",
		},
		ReportingOwnerRelationship: iv_models.ReportingOwnerRelationship{
			XMLName:           xml.Name{Local: "reportingOwnerRelationship"},
			IsDirector:        true,
			IsOfficer:         false,
			IsTenPercentOwner: false,
			IsOther:           false,
			OfficerTitle:      "",
			OtherText:         "",
		},
	}
	if expectedReporter != actualForm.ReportingOwners[0] {
		t.Errorf("reporting owner mismatch")
	}

	actualTransaction := actualForm.NonDerivativeTable.NonDerivativeTransactions[0]
	if actualTransaction.SecurityTitle.Value != "Voting Common Stock" {
		t.Errorf("security title mismatch: %s", actualTransaction.SecurityTitle.Value)
	}
	if actualTransaction.TransactionDate.Value != "2022-09-09" {
		t.Errorf("transaction date mismatch: %s", actualTransaction.TransactionDate.Value)
	}
	if actualTransaction.TransactionCoding.TransactionCode != "P" {
		t.Errorf("transaction code mismatch: %s", actualTransaction.TransactionCoding.TransactionCode)
	}
	if actualTransaction.TransactionAmounts.TransactionShares.Value != 88000 {
		t.Errorf("transaction shares mismatch: %f", actualTransaction.TransactionAmounts.TransactionShares.Value)
	}
	if actualTransaction.TransactionAmounts.TransactionPricePerShare.Value != 1.689 {
		t.Errorf("transaction shares mismatch: %f", actualTransaction.TransactionAmounts.TransactionShares.Value)
	}
	if actualTransaction.TransactionAmounts.TransactionAcquiredDisposedCode.Value != "A" {
		t.Errorf("transaction acquired disposed code mismatch: %s", actualTransaction.TransactionAmounts.TransactionAcquiredDisposedCode.Value)
	}

}
