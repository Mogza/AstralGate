package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateWallet(password string) string {
	keyStorePath := "/app/keystores"
	key := keystore.NewKeyStore(keyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	a, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

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
