package devault

import (
	"encoding/json"
	"strconv"
)

// GetAddressesByLabels returns addresses from wallet as key-value pair of label and address
func (rpc *RPC) GetAddressesByLabels() (map[string]string, error) {
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
func (rpc *RPC) GetWalletBalance(minconf int) (float64, error) {
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
func (rpc *RPC) GetBalance(address string) (float64, error) {
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

// SendToAddress sends specified amount to the provided address
//
// Wallet must be unlocked using 'UnlockWallet' method before sending transaction
func (rpc *RPC) SendToAddress(toAddress string, amount float64) (string, error) {
	res, err := rpc.Call("sendtoaddress", toAddress, amount)
	if err != nil {
		return "", err
	}
	tx, err := strconv.Unquote(string(res))
	if err != nil {
		return "", err
	}
	return tx, nil
}

// UnlockWallet unlocks wallet for certain period of time
func (rpc *RPC) UnlockWallet(passphrase string, time int64) error {
	_, err := rpc.Call("walletpassphrase", passphrase, time)
	if err != nil {
		return err
	}
	return nil
}

// GetTransaction returns detailed information about in-wallet transaction <txid>
func (rpc *RPC) GetTransaction(txid string) (*Transaction, error) {
	res, err := rpc.Call("gettransaction", txid)
	if err != nil {
		return nil, err
	}
	tx := new(Transaction)
	if err := json.Unmarshal(res, tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// GetNewAddress returns a new Devault address for receiving payments
// If 'label' is specified, it is added to the address book
// so payments received with the address will be associated with 'label'.
func (rpc *RPC) GetNewAddress(label *string) (string, error) {
	res, err := rpc.Call("getnewaddress", label)
	if err != nil {
		return "", err
	}
	tx, err := strconv.Unquote(string(res))
	if err != nil {
		return "", err
	}
	return tx, nil
}
