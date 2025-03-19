package dto

type CreditBalanceDTO struct {
	CustomerID    string `json:"customerId"`
	CreditBalance int    `json:"creditBalance"`
}

func CreditBalanceToDTO(customerID uint64, creditBalance int) *CreditBalanceDTO {
	return &CreditBalanceDTO{
		CustomerID:    IDToDTO(customerID),
		CreditBalance: creditBalance,
	}
}
