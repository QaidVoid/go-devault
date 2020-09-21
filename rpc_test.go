package devault

import (
	"testing"
)

func TestGetInfo(t *testing.T) {
	rpc := New("127.0.0.1", 13339, "user", "pass")
	_, err := rpc.GetInfo()
	if err != nil {
		t.Error(err)
	}
}

func TestGetBestBlockHash(t *testing.T) {
	rpc := New("127.0.0.1", 13339, "user", "pass")
	_, err := rpc.GetBestBlockHash()
	if err != nil {
		t.Error(err)
	}
}

func TestGetBlock(t *testing.T) {
	rpc := New("127.0.0.1", 13339, "user", "pass")
	_, err := rpc.GetBlock("00606632412e771796ff5e809913bc223d5c7927fec361158a21e9653cc32591")
	if err != nil {
		t.Error(err)
	}
}

func TestGetBlockChainInfo(t *testing.T) {
	rpc := New("127.0.0.1", 13339, "user", "pass")
	_, err := rpc.GetBlockChainInfo()
	if err != nil {
		t.Error(err)
	}
}
