package controller

import (
	datastore "ledgerCo-loans/DataStore"
	"ledgerCo-loans/entity"
	"math"
)

// BalanceCtlr is a struct
type BalanceCtlr struct {
}

// Do is a struct BalanceCtlr method
func (ctlr BalanceCtlr) Do(object interface{}) (entity.BalanceOutput, bool) {
	cmdData, ok := object.(entity.BalanceCmdData)
	if !ok {
		return entity.BalanceOutput{}, false
	}
	loanRec, exist := datastore.InMem.GetLoanInfo(cmdData.BankName, cmdData.BorrowerName)
	if !exist {
		// error loan does not exist
		return entity.BalanceOutput{}, false
	}
	totalLumpSumPaidTillEmiNum := loanRec.TotalLumpSumPaidTillEmiNum(cmdData.EmiNumber)
	emiMonthly := loanRec.MonthlyEmiAmount()
	totalEmiPaid := float64(cmdData.EmiNumber) * emiMonthly
	totalRepaymentsTillEmiNum := totalEmiPaid + totalLumpSumPaidTillEmiNum

	amountDue := loanRec.TotalRepayments() - totalRepaymentsTillEmiNum
	numEmiLeft := math.Ceil(amountDue / emiMonthly)

	return entity.BalanceOutput{
		BankName:     cmdData.BankName,
		BorrowerName: cmdData.BorrowerName,
		AmountPaid:   totalRepaymentsTillEmiNum,
		NumEmiLeft:   numEmiLeft,
	}, true
}
