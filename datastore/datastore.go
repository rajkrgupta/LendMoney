package datastore

import (
	"ledgerCo-loans/model"
)

// IDataStore is an interface
type IDataStore interface {
	StoreLoanDetails(loanDetail model.LoanInfo) bool
	GetLoanInfo(bankName, borrowerName string) model.LoanInfo
	SavePaymentRecord(bankName, borrowerName string, payment model.Payment) bool
}

// InMemoryDataStore is a struct
type InMemoryDataStore struct {
	LoanRecords map[model.LoanId]model.LoanInfo
}

var InMem InMemoryDataStore

func init() {
	InMem.LoanRecords = make(map[model.LoanId]model.LoanInfo)
}

// InMemoryDataStore method StoreLoanDetails
func (memStore InMemoryDataStore) StoreLoanDetails(loanDetail model.LoanInfo) bool {
	id := model.LoanId{
		BankName:     loanDetail.BankName,
		BorrowerName: loanDetail.BorrowerName,
	}
	memStore.LoanRecords[id] = loanDetail
	return true
}

// InMemoryDataStore method SavePaymentRecord
func (memStore InMemoryDataStore) SavePaymentRecord(bankName, borrowerName string, payment model.Payment) bool {
	id := model.LoanId{
		BankName:     bankName,
		BorrowerName: borrowerName,
	}
	loanRec, valid := memStore.LoanRecords[id]

	if valid {
		loanRec.Payments = append(loanRec.Payments, payment)
		memStore.LoanRecords[id] = loanRec
		return true
	}
	return false
}

// InMemoryDataStore method GetLoanInfo
func (memStore InMemoryDataStore) GetLoanInfo(bankName, borrowerName string) (model.LoanInfo, bool) {
	id := model.LoanId{
		BankName:     bankName,
		BorrowerName: borrowerName,
	}

	loanRec, valid := memStore.LoanRecords[id]
	return loanRec, valid
}
