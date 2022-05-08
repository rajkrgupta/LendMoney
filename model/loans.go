package model

import (
	"math"
	"time"
)

// Payment a struct
type Payment struct {
	EmiNumber  uint64
	AmountPaid float64
}

// LoanId a struct
type LoanId struct {
	BankName     string
	BorrowerName string
}

// LoanInfo a struct
type LoanInfo struct {
	BankName, BorrowerName string
	PrincipalAmount        float64
	InterestRate           float64
	TenureYears            uint64
	CreateAt               time.Time

	Payments []Payment
}

func (loan LoanInfo) TotalRepayments() float64 {
	simpleInterest := (loan.PrincipalAmount * (float64(loan.TenureYears)) * loan.InterestRate) / 100
	return (loan.PrincipalAmount + simpleInterest)
}

func (loan LoanInfo) MonthlyEmiAmount() float64 {
	emi := loan.TotalRepayments() / (float64(loan.TenureYears) * 12)
	return math.Ceil(emi)
}

func (loan LoanInfo) TotalLumpSumPaidTillEmiNum(emiNum uint64) float64 {
	var lumpSumPay float64 = 0
	for _, payment := range loan.Payments {
		if payment.EmiNumber <= emiNum {
			lumpSumPay += payment.AmountPaid
		}
	}
	return lumpSumPay
}
