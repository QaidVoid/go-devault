package devault

import (
	"encoding/json"
	"fmt"
)

// GetBestBlockHash ...
func (rpc *rpcStruct) GetBestBlockHash() (string, error) {
	res, err := rpc.Call("getbestblockhash")
	if err != nil {
		return "", err
	}

	var data struct {
		Result string      `json:"result"`
		Error  interface{} `json:"error"`
		ID     interface{} `json:"id"`
	}
	json.Unmarshal(res, &data)

	if data.Error != nil {
		return "", fmt.Errorf("%v", data.Error)
	}

	return data.Result, nil
}

// GetBlock ...
func (rpc *rpcStruct) GetBlock(hash string) (*Block, error) {
	res, err := rpc.Call("getblock", hash)
	if err != nil {
		return nil, err
	}
	block := new(Block)
	json.Unmarshal(res, block)

	if block.Error != nil {
		return nil, fmt.Errorf("%v", block.Error)
	}

	return block, err
}

// GetBlockChainInfo ...
func (rpc *rpcStruct) GetBlockChainInfo() (*BlockChainInfo, error) {
	res, err := rpc.Call("getblockchaininfo")
	if err != nil {
		return nil, err
	}
	blockchain := new(BlockChainInfo)
	if err := json.Unmarshal(res, blockchain); err != nil {
		return nil, err
	}
	if blockchain.Error != nil {
		return nil, fmt.Errorf("%v", blockchain.Error)
	}
	return blockchain, nil
}
