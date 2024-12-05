package service

import "github.com/Mogza/AstralGate/internal/models"

type RevenueResponse struct {
	Revenue float64 `json:"revenue"`
}

func GetTotalRevenue(wallet models.Wallet) RevenueResponse {
	return RevenueResponse{
		Revenue: wallet.Balance,
	}
}
