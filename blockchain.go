package devault

import (
	"encoding/json"
	"errors"
	"strconv"
)

// GetBestBlockHash returns the hash of the best (tip) block in the longest blockchain
func (rpc *rpcStruct) GetBestBlockHash() (string, error) {
	res, err := rpc.Call("getbestblockhash")
	if err != nil {
		return "", err
	}
	bh, err := strconv.Unquote(string(res))
	if err != nil {
		return "", err
	}
	return bh, nil
}

// GetBlock returns an Object with information about block 'hash'
func (rpc *rpcStruct) GetBlock(hash string) (*Block, error) {
	res, err := rpc.Call("getblock", hash)
	if err != nil {
		return nil, err
	}
	block := new(Block)
	json.Unmarshal(res, block)
	return block, err
}

// GetBlockChainInfo returns an object containing various state info regarding blockchain processing
func (rpc *rpcStruct) GetBlockChainInfo() (*BlockChain, error) {
	res, err := rpc.Call("getblockchaininfo")
	if err != nil {
		return nil, err
	}
	blockchain := new(BlockChain)
	if err := json.Unmarshal(res, blockchain); err != nil {
		return nil, err
	}
	return blockchain, nil
}

// GetBlockCount returns the number of blocks in the longest blockchain
func (rpc *rpcStruct) GetBlockCount() (int64, error) {
	res, err := rpc.Call("getblockcount")
	if err != nil {
		return 0, err
	}
	bc, err := strconv.ParseInt(string(res), 10, 64)
	if err != nil {
		return 0, err
	}
	return bc, nil
}

// GetBlockHash returns hash of block in best-block-chain at height provided
func (rpc *rpcStruct) GetBlockHash(height int) (string, error) {
	res, err := rpc.Call("getblockhash", height)
	if err != nil {
		return "", err
	}
	bh, err := strconv.Unquote(string(res))
	if err != nil {
		return "", err
	}
	return bh, nil
}

// GetBlockHeader returns an Object with information about blockheader 'hash'
func (rpc *rpcStruct) GetBlockHeader(blockhash string) (*BlockHeader, error) {
	res, err := rpc.Call("getblockheader", blockhash)
	if err != nil {
		return nil, err
	}
	blockheader := new(BlockHeader)
	if err := json.Unmarshal(res, blockheader); err != nil {
		return nil, err
	}
	return blockheader, nil
}

// GetBlockStats
//
// Compute per block statistics for a given window. All amounts are in satoshis.
//
// It won’t work for some heights with pruning.
//
// It won’t work without -txindex for utxo_size_inc, *fee or *feerate stats.
func (rpc *rpcStruct) GetBlockStats(hashOrHeight interface{}) (*BlockStats, error) {
	switch hashOrHeight.(type) {
	case int, string:
		res, err := rpc.Call("getblockstats", hashOrHeight)
		if err != nil {
			return nil, err
		}
		blockstat := new(BlockStats)
		if err := json.Unmarshal(res, blockstat); err != nil {
			return nil, err
		}
		return blockstat, nil
	default:
		return nil, errors.New("Invalid argument type. Expected string or int")
	}
}

// GetChainTips return information about all known tips in the block tree, including the main chain as well as orphaned branches
func (rpc *rpcStruct) GetChainTips() (*ChainTips, error) {
	res, err := rpc.Call("getchaintips")
	if err != nil {
		return nil, err
	}
	chaintips := new(ChainTips)
	if err := json.Unmarshal(res, chaintips); err != nil {
		return nil, err
	}
	return chaintips, nil
}

// GetChainTxStats compute statistics about the total number and rate of transactions in the chain
//
// Argument #1 - nblocks
//
// Type: numeric, optional, default=one month
//
// Size of the window in number of blocks
//
// Argument #2 - blockhash
//
// Type: string, optional, default=chain tip
//
// The hash of the block that ends the window.
func (rpc *rpcStruct) GetChainTxStats(args ...interface{}) (*ChainTxStats, error) {
	res, err := rpc.Call("getchaintxstats", args...)
	if err != nil {
		return nil, err
	}
	chaintx := new(ChainTxStats)
	if err := json.Unmarshal(res, chaintx); err != nil {
		return nil, err
	}
	return chaintx, nil
}

// GetDifficulty returns the proof-of-work difficulty as a multiple of the minimum difficulty
func (rpc *rpcStruct) GetDifficulty() (float64, error) {
	res, err := rpc.Call("getdifficulty")
	if err != nil {
		return 0, err
	}
	diff, err := strconv.ParseFloat(string(res), 64)
	if err != nil {
		return 0, err
	}
	return diff, nil
}

// GetMemPoolInfo returns details on the active state of the TX memory pool
func (rpc *rpcStruct) GetMemPoolInfo() (*MemPoolInfo, error) {
	res, err := rpc.Call("getmempoolinfo")
	if err != nil {
		return nil, err
	}
	mempool := new(MemPoolInfo)
	if err := json.Unmarshal(res, mempool); err != nil {
		return nil, err
	}
	return mempool, nil
}
