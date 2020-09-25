package devault

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RPC connection details
type RPC struct {
	Username string
	Password string
	Host     string
	Port     int
}

// RPCClient interface
type RPCClient interface {
	Call(string, ...interface{}) ([]byte, error)

	GetBestBlockHash() (string, error)
	GetBlockChainInfo() (*BlockChain, error)
	GetBlockCount() (int64, error)

	GetBlockHash(height int64) (string, error)
	GetBlock(hash string) (*Block, error)
	GetBlockHeader(blockhash string) (*BlockHeader, error)
	GetBlockStats(hashOrHeight interface{}) (*BlockStats, error)

	GetChainTips() (*ChainTips, error)
	GetChainTxStats(args ...interface{}) (*ChainTxStats, error)
	GetDifficulty() (float64, error)
	GetMemPoolInfo() (*MemPoolInfo, error)

	CombineRawTransaction(txs ...string) (string, error)
	CreateRawTransaction(inputs *Inputs, outputs *Outputs) (string, error)
	DecodeRawTransaction(signedTx string) (*DecodedTx, error)
	SignRawTransaction(rawTx string, privateKeys []string) (string, error)
	SendRawTransaction(signedTx string, allowHighFees bool) (string, error)

	GetAddressesByLabels() (map[string]string, error)
	GetWalletBalance(minconf int) (float64, error)
	GetBalance(address string) (float64, error)
	SendToAddress(toAddress string, amount float64) (string, error)
	UnlockWallet(passphrase string, time int64) error
	GetTransaction(txid string) (*Transaction, error)

	SubscribeNewBlock(c chan int64)
	SubscribeNewTransaction(c chan string)
}

// NewRPC creates new RPC interface
func NewRPC(host string, port int, username string, password string) RPCClient {
	var rpc RPCClient = &RPC{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
	return rpc
}

// Call specified RPC method
func (rpc *RPC) Call(method string, params ...interface{}) ([]byte, error) {
	u := fmt.Sprintf("http://%s:%s@%s:%d", rpc.Username, rpc.Password, rpc.Host, rpc.Port)
	reqBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "1.0",
		"method":  method,
		"params":  params,
	})
	client := http.Client{}
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(body, resp); err != nil {
		return nil, errors.New("Invalid Response Received")
	}

	if resp.Error != nil {
		return nil, err
	}

	m, err := json.Marshal(resp.Result)
	if err != nil {
		return nil, err
	}
	return m, nil
}
