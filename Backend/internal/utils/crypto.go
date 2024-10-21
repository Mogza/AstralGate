package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

func CreateWallet(password string) string {
	// Create keystore
	keyStorePath := "/app/keystores"
	key := keystore.NewKeyStore(keyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	a, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	// Change default keystore name to new one
	filenameAddress := strings.ToLower(a.Address.String()[2:])
	newFilename := fmt.Sprintf("%s.keystore", a.Address.Hex())

	var originalFilePath string
	err = filepath.Walk(keyStorePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(info.Name(), filenameAddress) {
			originalFilePath = path
		}
		return nil
	})
	if err != nil {
		LogFatal(err, "Error while walking the filepath")
	}

	if originalFilePath == "" {
		LogFatal(err, "No Keystore file found")
	}

	newFilePath := filepath.Join(keyStorePath, newFilename)
	err = os.Rename(originalFilePath, newFilePath)
	if err != nil {
		log.Fatal(err)
	}

	return a.Address.Hex()
}

func GetMaticBalance(address string) float64 {
	// Connect to polygon rpc
	polygonRpc := os.Getenv("POLYGON_AMOY_RPC")
	client, err := rpc.Dial(polygonRpc)
	if err != nil {
		log.Fatalf("Failed to connect to the Polygon network: %v", err)
	}
	defer client.Close()

	// Retrieve eth_balance
	var balanceHex string
	err = client.CallContext(context.Background(), &balanceHex, "eth_getBalance", address, "latest")
	if err != nil {
		log.Fatalf("Failed to fetch account balance: %v", err)
	}

	// Convert balance Hex to float
	balance, _ := new(big.Int).SetString(balanceHex[2:], 16)
	maticBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	maticFloat64, _ := maticBalance.Float64()

	return maticFloat64
}
