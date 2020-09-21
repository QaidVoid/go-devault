package devault

import (
	"bytes"
	"encoding/json"
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
	Call(string, []string) ([]byte, error)
	GetInfo() (*Info, error)
	GetBestBlockHash() (*BlockHash, error)
}

// Client ...
type Client string

// New RPC Client
func New(host string, port int, username string, password string) RPC {
	var rpc RPC = &rpcStruct{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
	return rpc
}

// Call ...
func (rpc *rpcStruct) Call(method string, params []string) ([]byte, error) {
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
	return body, nil
}

// GetInfo ...
func (rpc *rpcStruct) GetInfo() (*Info, error) {
	res, err := rpc.Call("getinfo", nil)
	if err != nil {
		return nil, err
	}
	info := new(Info)
	json.Unmarshal(res, info)
	return info, err
}
