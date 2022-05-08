package entity

type ProcessorType int

const (
	FileProcessor ProcessorType = iota
	// more processor to add ...
)

type DataStoreType int

const (
	InMemoryStore DataStoreType = iota
	// more data store to add ...
)

// LoanCmdData a struct
type LoanCmdData struct {
	BankName, BorrowerName string
	PrincipalAmount        float64
	InterestRate           float64
	TenureYears            uint64
}

// PaymentCmdData a struct
type PaymentCmdData struct {
	BankName, BorrowerName string
	EmiNumber              uint64
	AmountPaid             float64
}

// BalanceCmdData a struct
type BalanceCmdData struct {
	BankName, BorrowerName string
	EmiNumber              uint64
}

// BalanceOutput a struct
type BalanceOutput struct {
	BankName, BorrowerName string
	AmountPaid             float64
	NumEmiLeft             float64
}
