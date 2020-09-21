package devault_test

import (
	"devault"
	"testing"
)

func TestGetInfo(t *testing.T) {
	rpc := devault.New("127.0.0.1", 13339, "user", "pass")
	_, err := rpc.GetInfo()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBestBlockHash(t *testing.T) {
	rpc := devault.New("127.0.0.1", 13339, "user", "pass")
	_, err := rpc.GetBestBlockHash()
	if err != nil {
		t.Fatal(err)
	}
}
