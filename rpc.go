package devault

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// rpcStruct ...
type rpcStruct struct {
	Username string
	Password string
	Host     string
	Port     int
}

// RPC ...
type RPC interface {
	Call(string, ...interface{}) ([]byte, error)

	GetBestBlockHash() (string, error)
	GetBlockChainInfo() (*BlockChain, error)
	GetBlockCount() (int64, error)

	GetBlockHash(height int) (string, error)
	GetBlock(hash string) (*Block, error)
	GetBlockHeader(blockhash string) (*BlockHeader, error)
	GetBlockStats(hashOrHeight interface{}) (*BlockStats, error)

	GetChainTips() (*ChainTips, error)
	GetChainTxStats(args ...interface{}) (*ChainTxStats, error)
	GetDifficulty() (float64, error)
	GetMemPoolInfo() (*MemPoolInfo, error)

	CombineRawTransaction(txs ...string) (string, error)
	CreateRawTransaction(inputs *Inputs, outputs *Outputs) (string, error)
	SignRawTransaction(rawTx string, privateKeys []string) (string, error)
	SendRawTransaction(signedTx string, allowHighFees bool) (string, error)
}

// NewRPC RPC Client
func NewRPC(host string, port int, username string, password string) RPC {
	var rpc RPC = &rpcStruct{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
	return rpc
}

// Call ...
func (rpc *rpcStruct) Call(method string, params ...interface{}) ([]byte, error) {
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
