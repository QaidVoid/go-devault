# Devault JSON-RPC Client Library for Golang

[![PkgGoDev](https://pkg.go.dev/badge/github.com/QaidVoid/go-devault)](https://pkg.go.dev/github.com/QaidVoid/go-devault)

# Example
No errors are handled in the example. But you should always handle errors appropriately..

```go
package main

import (
	"log"

	"github.com/QaidVoid/go-devault"
)

func main() {
	// Initialise devault RPC Client interface
	testnet := devault.NewRPC("127.0.0.1", 13339, "user", "pass")
	mainnet := devault.NewRPC("127.0.0.1", 3339, "user", "pass")

  // Get new address from wallet
	addr, _ := testnet.GetNewAddress(nil)
	log.Printf("New testnet address: %s", addr)

  // Get balance from wallet with atleast 6 confs
	balance, _ := testnet.GetWalletBalance(6)
	log.Printf("Wallet Balance: %.3f", balance)

	// Unlock wallet and send from wallet
	testnet.UnlockWallet("Passphrase", 60)
	tx, _ := testnet.SendToAddress("dvtest:qrxxxxxxxxxxxxxxxxxxxx", 500)

	// Inputs for raw transaction
	inputs := devault.Inputs{
		{
			TXID: "txid",
			VOut: 0,
		},
	}

	// Outputs for raw transaction
	var outputs devault.Outputs
	output := make(map[string]float64)
	output["devault:qrxxxxxxxxxxxxxxxxxxxxxxxx"] = 100
	outputs = append(outputs, output)

	// Private keys to sign raw transaction with
	keys := []string{
		"priv:zqxxxxxxxxxxxxxxxxxxxxxxxx",
		"priv:zqxxxxxxxxxxxxxxxxxxxxxxx2",
	}

	raw, _ := mainnet.CreateRawTransaction(&inputs, &outputs)
	tx, _ = mainnet.SignRawTransaction(raw, keys)
	mainnet.SendRawTransaction(tx, false)
	log.Fatal("Transaction success!")
}
```
