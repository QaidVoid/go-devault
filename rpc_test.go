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

func TestRawTransaction(t *testing.T) {
	inputs := &Inputs{
		{
			TXID: "01e4014fc8466bab448c01ff277c4b43f0393bf407aa09f87a8274e15c222f43",
			VOut: 0,
		},
	}
	var outputs Outputs
	output := make(map[string]float64)
	output["dvtest:qrpfpackqd730cgp7qutltsndhpugc9qay4peyk2e4"] = 150

	outputs = append(outputs, output)

	raw, err := rpc.CreateRawTransaction(inputs, &outputs)
	if err != nil {
		t.Error(err)
		return
	}

	tx, err := rpc.SignRawTransaction(raw, []string{"testpriv:zq0hn5m9h7yem3wrezyutv8tnphg70aaelzj5el35nejvypuzudvskt0tj0yw"})
	if err != nil {
		t.Error(err)
		return
	}

	if _, err = rpc.SendRawTransaction(tx, false); err != nil {
		t.Error(err)
	}
}

func TestCombineRawTransaction(t *testing.T) {
	tx, err := rpc.CombineRawTransaction("0200000001efec297945eff8a81602e8f9fdfccdcc9d2583c4dff3655da570dec95049dc30000000004847304402201c638f4c7191930e42dbbb00bf698f36781042f0ab43742f04e2cf19e4ac1c3d0220062b9f121ef41deb2ba06f832801c07ab0727d5e82e382008dfdcee2c83435b741ffffffff0100f2052a010000001976a914c290f716037d17e101f038bfae136dc3c460a0e988ac00000000", "0200000001dc36a297b2628c2e2917fc161e999b3cdd8e268fd2475bb143bf47de9d058872000000004847304402204ec4c208f050ad66dc8f9ca36480714083e81200300d95b083baf74af174ea0d02206d0e5c0c1370c00ec6ae2323992a70666af0a445941c0dda96bcb4675d98a86541ffffffff0100f2052a010000001976a914c290f716037d17e101f038bfae136dc3c460a0e988ac00000000")
	if err != nil {
		t.Error(err)
	}
	if _, err := rpc.SendRawTransaction(tx, false); err != nil {
		t.Error(err)
	}
}
