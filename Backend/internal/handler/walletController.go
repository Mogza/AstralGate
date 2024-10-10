package handler

import (
	"fmt"
	"github.com/Mogza/AstralGate/internal/models"
	"github.com/Mogza/AstralGate/internal/utils"
)

func (h Handler) UpdateBalance() {
	var wallets []models.Wallet
	err := h.DB.Where("currency = ?", "POL").Find(&wallets).Error
	if err != nil {
		fmt.Println("[UpdateBalance] Error while retrieving all wallets: ", err)
		return
	}

	if len(wallets) == 0 {
		fmt.Println("[UpdateBalance] No wallets found")
		return
	}

	for _, wallet := range wallets {
		wallet.Balance = utils.GetMaticBalance(wallet.CryptoAddress)

		err = h.DB.Save(wallet).Error
		if err != nil {
			fmt.Println("[UpdateBalance] Error while saving wallet ( ", wallet.CryptoAddress, "): ", err)
		}
	}
}
