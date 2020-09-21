package devault

import (
	"testing"
)

var rpc = NewRPC("127.0.0.1", 13339, "user", "pass")

func TestGetBestBlockHash(t *testing.T) {
	if _, err := rpc.GetBestBlockHash(); err != nil {
		t.Error(err)
	}
}

func TestGetBlock(t *testing.T) {
	if _, err := rpc.GetBlock("00606632412e771796ff5e809913bc223d5c7927fec361158a21e9653cc32591"); err != nil {
		t.Error(err)
	}
}

func TestGetBlockChainInfo(t *testing.T) {
	if _, err := rpc.GetBlockChainInfo(); err != nil {
		t.Error(err)
	}
}
func TestGetBlockHeader(t *testing.T) {
	if _, err := rpc.GetBlockHeader("0031f614b80186345fe2b4ab2d676d63f74ccb296592fb0d4bf5fa5a8bc81f36"); err != nil {
		t.Error(err)
	}
}

func TestGetBlockStats(t *testing.T) {
	if _, err := rpc.GetBlockStats(9000); err != nil {
		t.Error(err)
	}
}
func TestGetChainTips(t *testing.T) {
	if _, err := rpc.GetChainTips(); err != nil {
		t.Error(err)
	}
}

func TestGetChainTxStats(t *testing.T) {
	if _, err := rpc.GetChainTxStats(2019, "00606632412e771796ff5e809913bc223d5c7927fec361158a21e9653cc32591"); err != nil {
		t.Error(err)
	}
}

func TestGetDifficulty(t *testing.T) {
	if _, err := rpc.GetDifficulty(); err != nil {
		t.Error(err)
	}
}

func TestGetMemPoolInfo(t *testing.T) {
	if _, err := rpc.GetMemPoolInfo(); err != nil {
		t.Error(err)
	}
}
