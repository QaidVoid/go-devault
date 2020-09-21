package devault

// RPCInfo ...
type RPCInfo struct {
	Result struct {
		Version         int     `json:"version"`
		ProtocolVersion int     `json:"protocolversion"`
		WalletVersion   int     `json:"walletversion"`
		Blocks          int     `json:"blocks"`
		TimeOffset      int     `json:"timeoffset"`
		Connections     int     `json:"connections"`
		KeyPoolOldest   int     `json:"keypoololdest"`
		KeyPoolSize     int     `json:"keypoolsize"`
		UnlockedUntil   int     `json:"unlocked_until"`
		Balance         float64 `json:"balance"`
		PayTxFee        float64 `json:"paytxfee"`
		RelayFee        float64 `json:"relayfee"`
		Difficulty      float64 `json:"difficulty"`
		TestNet         bool    `json:"testnet"`
		Proxy           string  `json:"proxy"`
		Errors          string  `json:"errors"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    interface{} `json:"id"`
}

// Block ...
type Block struct {
	Result struct {
		Confirmations     int      `json:"confirmations"`
		Size              int      `json:"size"`
		Height            int      `json:"height"`
		Version           int      `json:"version"`
		Time              int      `json:"time"`
		MedianTime        int      `json:"mediantime"`
		Nonce             int      `json:"nonce"`
		Hash              string   `json:"hash"`
		VersionHex        string   `json:"versionHex"`
		MerkleRoot        string   `json:"merkleroot"`
		Bits              string   `json:"bits"`
		ChainWork         string   `json:"chainwork"`
		PreviousBlockHash string   `json:"previousblockhash"`
		Transactions      []string `json:"tx"`
		Difficulty        float64  `json:"difficulty"`
	}
	Error interface{} `json:"error"`
	ID    interface{} `json:"id"`
}

// BlockChainInfo ...
type BlockChainInfo struct {
	Result struct {
		Chain                string  `json:"chain"`
		BestBlockHash        string  `json:"bestblockhash"`
		Blocks               int     `json:"blocks"`
		Headers              int     `json:"headers"`
		Difficulty           float64 `json:"difficulty"`
		MedianTime           int     `json:"mediantime"`
		VerificationProgress float64 `json:"verificationprogress"`
		InitialBlockDownload bool    `json:"initialblockdownload"`
		ChainWork            string  `json:"chainwork"`
		CoinSupply           float64 `json:"coinsupply"`
		Pruned               bool    `json:"pruned"`
		Warnings             string  `json:"warnings"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    interface{} `json:"id"`
}
