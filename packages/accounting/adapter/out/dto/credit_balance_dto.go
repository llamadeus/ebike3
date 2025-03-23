package dto

type CreditBalanceDTO struct {
	CustomerID    string `json:"customerId"`
	CreditBalance int32  `json:"creditBalance"`
}

func CreditBalanceToDTO(customerID uint64, creditBalance int32) *CreditBalanceDTO {
	return &CreditBalanceDTO{
		CustomerID:    IDToDTO(customerID),
		CreditBalance: creditBalance,
	}
}
