package cryptochain

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_transaction(t *testing.T) {
	cosmosTx := CryptoTx{
		AddrPrefix: "tcro",
		From:      "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
		To:        "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
		Denom:     "basetcro",
		FeeDenom:  "basetcro",
		Memo:      "Hello Test Memo",
		ChainID:   "testnet-croeseid-2",
		PublicKey: "03c3d281a28592adce81bee3094f00eae26932cbc682fba239b90f47dac9fe7036",
		Amount:    1210,
		Fee:       6500,
		AccNum:    41,
		AccSeq:    13,
		GasLimit:  280000,
		Timeout:   341910,
	}

	unsignedTrans, hash, err := cosmosTx.getUnsignedTxAndHash()

	if err != nil {
		t.Error("create failed")
		return
	}

	fmt.Println("tx : ", unsignedTrans)
	fmt.Println("hash : ", hash)

	private_key, _ := hex.DecodeString("d10928ec4afc23578f39eda4249c403bf704c99ebd7d287a24bd2700903408e3")
	signature, err := signTransactionHash(hash, private_key)
	if err != nil {
		t.Error("sign failed")
		return
	}

	// signature = hex.EncodeToString([]byte{169,213,22,69,126,153,158,86,46,52,137,108,112,198,224,171,82,18,230,38,133,30,179,81,31,245,74,123,106,248,57,172,95,199,41,201,188,125,51,35,152,52,112,59,149,92,235,23,217,162,165,83,76,118,72,22,31,24,222,231,10,100,65,227})
	fmt.Println("signature : ", signature)

	broadcastBytes, err := getBroadcastBytes(unsignedTrans, signature)
	if err != nil {
		t.Error("combine failed")
		return
	}

	fmt.Println("broadcast : ", broadcastBytes)
}
