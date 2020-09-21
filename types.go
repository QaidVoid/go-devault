package devault

// Info ...
type Info struct {
	Result struct {
		Version         int     `json:"version"`
		ProtocolVersion int     `json:"protocolversion"`
		WalletVersion   int     `json:"walletversion"`
		Balance         float64 `json:"balance"`
		Blocks          int     `json:"blocks"`
		TimeOffset      int     `json:"timeoffset"`
		Connections     int     `json:"connections"`
		Proxy           string  `json:"proxy"`
		Difficulty      float64 `json:"difficulty"`
		TestNet         bool    `json:"testnet"`
		KeyPoolOldest   int     `json:"keypoololdest"`
		KeyPoolSize     int     `json:"keypoolsize"`
		UnlockedUntil   int     `json:"unlocked_until"`
		PayTxFee        float64 `json:"paytxfee"`
		RelayFee        float64 `json:"relayfee"`
		Errors          string  `json:"errors"`
	} `json:"result"`
	Error *string     `json:"error"`
	ID    interface{} `json:"id"`
}

// BlockHash ...
type BlockHash string
