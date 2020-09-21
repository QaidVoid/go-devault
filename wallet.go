package devault

import (
	"encoding/json"
	"strconv"
)

// AbandonTransaction Mark in-wallet transaction <txid> as abandoned This will mark this transaction and all its in-wallet descendants as abandoned which will allow for their inputs to be respent. It can be used to replace “stuck” or evicted transactions.
//
// It only works on transactions which are not included in a block and are not currently in the mempool.
//
// It has no effect on transactions which are already abandoned.
func (rpc *rpcStruct) AbandonTransaction(txid string) {
	// TODO
}

func (rpc *rpcStruct) GetAddressesByLabels() (map[string]string, error) {
	res, err := rpc.Call("getaddressesbylabels")
	if err != nil {
		return nil, err
	}
	var addresses = make(map[string]string)
	json.Unmarshal(res, &addresses)
	return addresses, nil
}

// GetWalletBalance returns the total available balance
//
// The available balance is what the wallet considers currently spendable, and is thus affected by options which limit spendability such as -spendzeroconfchange.
func (rpc *rpcStruct) GetWalletBalance(minconf int) (float64, error) {
	res, err := rpc.Call("getbalance", "*", minconf)
	if err != nil {
		return 0, err
	}
	bal, err := strconv.ParseFloat(string(res), 64)
	if err != nil {
		return 0, err
	}
	return bal, nil
}

// GetBalance returns the available balance on specified address
func (rpc *rpcStruct) GetBalance(address string) (float64, error) {
	res, err := rpc.Call("getaddressbalance", address)
	if err != nil {
		return 0, err
	}
	var bal struct {
		Balance  float64
		Received float64
	}

	if err := json.Unmarshal(res, &bal); err != nil {
		return 0, err
	}
	return bal.Balance, nil
}
