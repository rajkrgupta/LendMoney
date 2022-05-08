package controller

import (
	datastore "ledgerCo-loans/DataStore"
	"ledgerCo-loans/entity"
	"ledgerCo-loans/model"
)

// PaymentCtlr is a struct
type PaymentCtlr struct {
}

// Do is a struct PaymentCtlr method
func (ctlr PaymentCtlr) Do(object interface{}) bool {
	cmdData, ok := object.(entity.PaymentCmdData)
	if !ok {
		return false
	}
	loanRec, exist := datastore.InMem.GetLoanInfo(cmdData.BankName, cmdData.BorrowerName)
	if !exist {
		// error loan does not exist
		return false
	}
	var maxEmis = loanRec.TenureYears * 12
	if cmdData.EmiNumber > maxEmis {
		// error emi number invalid
		return false
	}
	payment := model.Payment{
		EmiNumber:  cmdData.EmiNumber,
		AmountPaid: cmdData.AmountPaid,
	}
	return datastore.InMem.SavePaymentRecord(cmdData.BankName, cmdData.BorrowerName, payment)
}
