package devault

import "encoding/json"

// GetBestBlockHash ...
func (rpc *rpcStruct) GetBestBlockHash() (*BlockHash, error) {
	res, err := rpc.Call("getbestblockhash", []string{})
	if err != nil {
		return nil, err
	}

	var data struct {
		Result BlockHash   `json:"result"`
		Error  interface{} `json:"error"`
		ID     interface{} `json:"id"`
	}
	json.Unmarshal(res, &data)
	return &data.Result, nil
}
