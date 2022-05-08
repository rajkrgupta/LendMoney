package controller

import (
	datastore "ledgerCo-loans/DataStore"
	"ledgerCo-loans/entity"
	"ledgerCo-loans/model"
	"time"
)

type LoanCtlr struct {
}

// Do is a struct LoanCtlr method
func (ctlr LoanCtlr) Do(object interface{}) bool {
	cmdData, ok := object.(entity.LoanCmdData)
	if !ok {
		return false
	}
	_, exist := datastore.InMem.GetLoanInfo(cmdData.BankName, cmdData.BorrowerName)
	if exist {
		// error loan already created
		return false
	}

	loanDetail := model.LoanInfo{
		BankName:        cmdData.BankName,
		BorrowerName:    cmdData.BorrowerName,
		PrincipalAmount: cmdData.PrincipalAmount,
		InterestRate:    cmdData.InterestRate,
		TenureYears:     cmdData.TenureYears,
		CreateAt:        time.Now().UTC(),
		Payments:        []model.Payment{},
	}
	return datastore.InMem.StoreLoanDetails(loanDetail)
}
