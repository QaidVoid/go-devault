package devault

// Response structure for RPC call response
type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
	ID     interface{} `json:"id,omitempty"`
}

// Block stores for blocks information
type Block struct {
	Confirmations     int      `json:"confirmations,omitempty"`
	Size              int      `json:"size,omitempty"`
	Height            int      `json:"height,omitempty"`
	Version           int      `json:"version,omitempty"`
	Time              int      `json:"time,omitempty"`
	MedianTime        int      `json:"mediantime,omitempty"`
	Nonce             int      `json:"nonce,omitempty"`
	Hash              string   `json:"hash,omitempty"`
	VersionHex        string   `json:"versionHex,omitempty"`
	MerkleRoot        string   `json:"merkleroot,omitempty"`
	Bits              string   `json:"bits,omitempty"`
	ChainWork         string   `json:"chainwork,omitempty"`
	PreviousBlockHash string   `json:"previousblockhash,omitempty"`
	Transactions      []string `json:"tx,omitempty"`
	Difficulty        float64  `json:"difficulty,omitempty"`
}

// BlockChain stores blockchain information
type BlockChain struct {
	Chain                string  `json:"chain,omitempty"`
	BestBlockHash        string  `json:"bestblockhash,omitempty"`
	Blocks               int     `json:"blocks,omitempty"`
	Headers              int     `json:"headers,omitempty"`
	MedianTime           int     `json:"mediantime,omitempty"`
	Difficulty           float64 `json:"difficulty,omitempty"`
	CoinSupply           float64 `json:"coinsupply,omitempty"`
	VerificationProgress float64 `json:"verificationprogress,omitempty"`
	InitialBlockDownload bool    `json:"initialblockdownload,omitempty"`
	Pruned               bool    `json:"pruned,omitempty"`
	ChainWork            string  `json:"chainwork,omitempty"`
	Warnings             string  `json:"warnings,omitempty"`
}

// BlockHeader stores block's header information for provided block hash
type BlockHeader struct {
	Confirmations     int     `json:"confirmations,omitempty"`
	Height            int     `json:"height,omitempty"`
	Version           int     `json:"version,omitempty"`
	Time              int     `json:"time,omitempty"`
	MedianTime        int     `json:"mediantime,omitempty"`
	Nonce             int     `json:"nonce,omitempty"`
	Hash              string  `json:"hash,omitempty"`
	VersionHex        string  `json:"versionHex,omitempty"`
	MerkleRoot        string  `json:"merkleroot,omitempty"`
	Bits              string  `json:"bits,omitempty"`
	ChainWork         string  `json:"chainwork,omitempty"`
	PreviousBlockHash string  `json:"previousblockhash,omitempty"`
	NextBlockHash     string  `json:"nextblockhash,omitempty"`
	Difficulty        float64 `json:"difficulty,omitempty"`
}

// BlockStats stores block stats on specified block hash or height
type BlockStats struct {
	BlockHash       string  `json:"blockhash,omitempty"`
	AvgFee          float64 `json:"avgfee,omitempty"`
	AvgFeeRate      float64 `json:"avgfeerate,omitempty"`
	MaxFee          float64 `json:"maxfee,omitempty"`
	MaxFeeRate      float64 `json:"maxfeerate,omitempty"`
	MedianFee       float64 `json:"medianfee,omitempty"`
	MedianFeeRate   float64 `json:"medianfeerate,omitempty"`
	MinFee          float64 `json:"minfee,omitempty"`
	MinFeeRate      float64 `json:"minfeerate,omitempty"`
	Subsidy         float64 `json:"subsidy,omitempty"`
	TotalOut        float64 `json:"total_out,omitempty"`
	TotalFee        float64 `json:"totalfee,omitempty"`
	AvgTxSize       int     `json:"avgtxsize,omitempty"`
	Height          int     `json:"height,omitempty"`
	NumberOfInputs  int     `json:"ins,omitempty"`
	MaxTxSize       int     `json:"maxtxsize,omitempty"`
	MedianTime      int     `json:"mediantime,omitempty"`
	MedianTxSize    int     `json:"mediantxsize,omitempty"`
	MinTxSize       int     `json:"mintxsize,omitempty"`
	NumberOfOutputs int     `json:"outs,omitempty"`
	Time            int     `json:"time,omitempty"`
	TotalSize       int     `json:"total_size,omitempty"`
	NumberOfTxs     int     `json:"txs,omitempty"`
	UTXOIncrease    int     `json:"utxo_increase,omitempty"`
	UTXOSizeInc     int     `json:"utxo_size_inc,omitempty"`
}

// ChainTip stores information of chain tip
type ChainTip struct {
	Height    int    `json:"height,omitempty"`
	BranchLen int    `json:"branchlen,omitempty"`
	Hash      string `json:"hash,omitempty"`
	Status    string `json:"status,omitempty"`
}

// ChainTips store array of ChainTip
type ChainTips []ChainTip

// ChainTxStats stores information about the total number and rate of transactions in the chain
type ChainTxStats struct {
	Time                 int     `json:"time,omitempty"`
	TxCount              int     `json:"txcount,omitempty"`
	WindowBlockCount     int     `json:"window_block_count,omitempty"`
	WindowTxCount        int     `json:"window_tx_count,omitempty"`
	WindowInterval       int     `json:"window_interval,omitempty"`
	WindowFinalBlockHash string  `json:"window_final_block_hash,omitempty"`
	TxRate               float64 `json:"txrate,omitempty"`
}

// MemPoolInfo stores details on the active state of the TX memory pool
type MemPoolInfo struct {
	Loaded        bool    `json:"loaded,omitempty"`
	Size          int     `json:"size,omitempty"`
	MaxMemPool    int     `json:"maxmempool,omitempty"`
	Bytes         float64 `json:"bytes,omitempty"`
	Usage         float64 `json:"usage,omitempty"`
	MemPoolMinFee float64 `json:"mempoolminfee,omitempty"`
	MinRelayTxFee float64 `json:"minrelaytxfee,omitempty"`
}

// Inputs holds array of input data for transaction
type Inputs []Input

// Input data
type Input struct {
	TXID     string  `json:"txid"`
	VOut     int     `json:"vout"`
	Sequence float64 `json:"sequence,omitempty"`
}

// Outputs holds array of transaction output data
type Outputs []map[string]float64

// SignedRawTx holds signed raw transaction info
type SignedRawTx struct {
	Hex      string `json:"hex,omitempty"`
	Complete bool   `json:"complete,omitempty"`
	Errors   []struct {
		Error string `json:"error,omitempty"`
		TXID  string `json:"txid,omitempty"`
	} `json:"errors,omitempty"`
}

// DecodedTx holds decoded raw transaction info
type DecodedTx struct {
	TXID     string `json:"txid,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Version  int    `json:"version,omitempty"`
	Size     int    `json:"size,omitempty"`
	LockTime int    `json:"locktime,omitempty"`
	VIn      []struct {
		TXID      string `json:"txid,omitempty"`
		Vout      int    `json:"vout,omitempty"`
		ScriptSig struct {
			Asm string `json:"asm,omitempty"`
			Hex string `json:"hex,omitempty"`
		} `json:"scriptSig,omitempty"`
		Coinbase    string   `json:"coinbase,omitempty"`
		TXInWitness []string `json:"txinwitness,omitempty"`
		Sequence    int      `json:"sequence,omitempty"`
	} `json:"vin,omitempty"`
	Vout []struct {
		Value        float64 `json:"value,omitempty"`
		N            int     `json:"n,omitempty"`
		ScriptPubKey struct {
			Asm       string   `json:"asm,omitempty"`
			Hex       string   `json:"hex,omitempty"`
			ReqSigs   int      `json:"reqSigs,omitempty"`
			Type      string   `json:"type,omitempty"`
			Addresses []string `json:"addresses,omitempty"`
		} `json:"scriptPubKey,omitempty"`
	} `json:"vout,omitempty"`
}
