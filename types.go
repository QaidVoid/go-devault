package devault

// Response ...
type Response struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	ID     interface{} `json:"id"`
}

// Block ...
type Block struct {
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

// BlockChain ...
type BlockChain struct {
	Chain                string  `json:"chain"`
	BestBlockHash        string  `json:"bestblockhash"`
	Blocks               int     `json:"blocks"`
	Headers              int     `json:"headers"`
	MedianTime           int     `json:"mediantime"`
	Difficulty           float64 `json:"difficulty"`
	CoinSupply           float64 `json:"coinsupply"`
	VerificationProgress float64 `json:"verificationprogress"`
	InitialBlockDownload bool    `json:"initialblockdownload"`
	Pruned               bool    `json:"pruned"`
	ChainWork            string  `json:"chainwork"`
	Warnings             string  `json:"warnings"`
}

// BlockHeader ...
type BlockHeader struct {
	Confirmations     int     `json:"confirmations"`
	Height            int     `json:"height"`
	Version           int     `json:"version"`
	Time              int     `json:"time"`
	MedianTime        int     `json:"mediantime"`
	Nonce             int     `json:"nonce"`
	Hash              string  `json:"hash"`
	VersionHex        string  `json:"versionHex"`
	MerkleRoot        string  `json:"merkleroot"`
	Bits              string  `json:"bits"`
	ChainWork         string  `json:"chainwork"`
	PreviousBlockHash string  `json:"previousblockhash"`
	NextBlockHash     string  `json:"nextblockhash"`
	Difficulty        float64 `json:"difficulty"`
}

// BlockStats ...
type BlockStats struct {
	BlockHash       string  `json:"blockhash"`
	AvgFee          float64 `json:"avgfee"`
	AvgFeeRate      float64 `json:"avgfeerate"`
	MaxFee          float64 `json:"maxfee"`
	MaxFeeRate      float64 `json:"maxfeerate"`
	MedianFee       float64 `json:"medianfee"`
	MedianFeeRate   float64 `json:"medianfeerate"`
	MinFee          float64 `json:"minfee"`
	MinFeeRate      float64 `json:"minfeerate"`
	Subsidy         float64 `json:"subsidy"`
	TotalOut        float64 `json:"total_out"`
	TotalFee        float64 `json:"totalfee"`
	AvgTxSize       int     `json:"avgtxsize"`
	Height          int     `json:"height"`
	NumberOfInputs  int     `json:"ins"`
	MaxTxSize       int     `json:"maxtxsize"`
	MedianTime      int     `json:"mediantime"`
	MedianTxSize    int     `json:"mediantxsize"`
	MinTxSize       int     `json:"mintxsize"`
	NumberOfOutputs int     `json:"outs"`
	Time            int     `json:"time"`
	TotalSize       int     `json:"total_size"`
	NumberOfTxs     int     `json:"txs"`
	UTXOIncrease    int     `json:"utxo_increase"`
	UTXOSizeInc     int     `json:"utxo_size_inc"`
}

// ChainTip ...
type ChainTip struct {
	Height    int    `json:"height"`
	BranchLen int    `json:"branchlen"`
	Hash      string `json:"hash"`
	Status    string `json:"status"`
}

// ChainTips ...
type ChainTips []ChainTip

// ChainTxStats ...
type ChainTxStats struct {
	Time                 int     `json:"time"`
	TxCount              int     `json:"txcount"`
	WindowBlockCount     int     `json:"window_block_count"`
	WindowTxCount        int     `json:"window_tx_count"`
	WindowInterval       int     `json:"window_interval"`
	WindowFinalBlockHash string  `json:"window_final_block_hash"`
	TxRate               float64 `json:"txrate"`
}

// MemPoolInfo ...
type MemPoolInfo struct {
	Loaded        bool    `json:"loaded"`
	Size          int     `json:"size"`
	MaxMemPool    int     `json:"maxmempool"`
	Bytes         float64 `json:"bytes"`
	Usage         float64 `json:"usage"`
	MemPoolMinFee float64 `json:"mempoolminfee"`
	MinRelayTxFee float64 `json:"minrelaytxfee"`
}

// Inputs ...
type Inputs []Input

// Input ...
type Input struct {
	TXID     string  `json:"txid"`
	VOut     int     `json:"vout"`
	Sequence float64 `json:"sequence,omitempty"`
}

// Outputs ...
type Outputs []map[string]float64

// SignedRawTx ...
type SignedRawTx struct {
	Hex      string `json:"hex,omitempty"`
	Complete bool   `json:"complete,omitempty"`
	Errors   []struct {
		Error string `json:"error,omitempty"`
		TXID  string `json:"txid,omitempty"`
	} `json:"errors,omitempty"`
}
