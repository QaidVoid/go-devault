package devault

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// CreateRawTransaction creates a transaction spending the given inputs and creating new outputs
//
// Outputs can be addresses or data.
//
// Returns hex-encoded raw transaction.
//
// Note that the transaction's inputs are not signed, and it is not stored in the wallet or transmitted to the network
func (rpc *rpcStruct) CreateRawTransaction(inputs *Inputs, outputs *Outputs) (string, error) {
	res, err := rpc.Call("createrawtransaction", inputs, outputs)
	if err != nil {
		return "", err
	}
	rawTx, err := strconv.Unquote(string(res))
	if err != nil {
		return "", err
	}
	return rawTx, nil
}

// SignRawTransaction signs inputs for raw transaction (serialized, hex-encoded)
//
// Private keys is an array of Bech32 encoded private keys
func (rpc *rpcStruct) SignRawTransaction(rawTx string, privateKeys []string) (string, error) {
	res, err := rpc.Call("signrawtransactionwithkey", rawTx, privateKeys)
	if err != nil {
		return "", err
	}
	signedTx := new(SignedRawTx)
	if err := json.Unmarshal(res, signedTx); err != nil {
		return "", err
	}
	if len(signedTx.Errors) > 0 {
		sb := strings.Builder{}
		sb.WriteString("TXSign Error")
		for i, v := range signedTx.Errors {
			sb.WriteString(fmt.Sprintf("\n%d: %s, TXID: %s", i+1, v.Error, v.TXID))
		}
		return "", errors.New(sb.String())
	}
	return signedTx.Hex, nil
}

// SendRawTransaction submits raw transaction (serialized, hex-encoded) to local node and network
//
// Returns transaction hash
func (rpc *rpcStruct) SendRawTransaction(signedTx string, allowHighFees bool) (string, error) {
	res, err := rpc.Call("sendrawtransaction", signedTx, allowHighFees)
	if err != nil {
		return "", err
	}
	tx, err := strconv.Unquote(string(res))
	if err != nil {
		return "", err
	}
	return tx, nil
}

// CombineRawTransaction combines multiple partially signed transactions into one transaction.
//
// The combined transaction may be another partially signed transaction or a fully signed transaction.
func (rpc *rpcStruct) CombineRawTransaction(txs ...string) (string, error) {
	res, err := rpc.Call("combinerawtransaction", txs)
	if err != nil {
		return "", err
	}
	tx, err := strconv.Unquote(string(res))
	if err != nil {
		return "", err
	}
	return tx, nil
}
