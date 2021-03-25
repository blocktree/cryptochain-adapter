package cryptochain

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

var testUrl = "https://testnet-croeseid.crypto.com:1317"

func Test_getBlockHeight(t *testing.T) {
	c := NewClient(testUrl, false)

	r, err := c.getBlockHeight()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_getBlockByHash(t *testing.T) {
	hash := "523724B328F9082A0627E73BE1911178E709F535BA372DB03D2F221A1D3A9E57"

	c := NewClient("http://localhost:9922/", false)

	r, err := c.Call("blocks/signature/"+hash, nil, "GET")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_getBlockHash(t *testing.T) {
	c := NewClient(testUrl, false)

	height := uint64(945033)

	r, err := c.getBlockHash(height)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
func Test_tmp(t *testing.T) {
	test, err := time.Parse(time.RFC3339Nano, "2019-05-08T02:13:41.937681458Z")
	fmt.Println(err)
	fmt.Println(test.Unix())
}
func Test_getBalance(t *testing.T) {
	c := NewClient(testUrl, true)

	address := "tcro1yr0p3t3uz9fxjzgle3m69u52gkmqmg6np3ffgr"

	r, err := c.getBalance(address, "basetcro")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getTransaction(t *testing.T) {
	c := NewClient(testUrl, false)
	txid := "564E47FE4A5C23244238A824C0D75FD91EB7131ABF7E9A8B3A89116592E3A152"

	path := "/cosmos/tx/v1beta1/txs/" + txid
	r, err := c.Call(path, nil, "GET")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	trx := NewTransaction(r, "cosmos-sdk/StdTx", "/cosmos.bank.v1beta1.MsgSend", "basetcro")

	fmt.Println(trx)
}

func Test_convert(t *testing.T) {

	amount := uint64(5000000001)

	amountStr := fmt.Sprintf("%d", amount)

	fmt.Println(amountStr)

	d, _ := decimal.NewFromString(amountStr)

	w, _ := decimal.NewFromString("100000000")

	d = d.Div(w)

	fmt.Println(d.String())

	d = d.Mul(w)

	fmt.Println(d.String())

	r, _ := strconv.ParseInt(d.String(), 10, 64)

	fmt.Println(r)

	fmt.Println(time.Now().UnixNano())
}

func Test_getTransactionByAddresses(t *testing.T) {
	addrs := "ARAA8AnUYa4kWwWkiZTTyztG5C6S9MFTx11"

	c := NewClient("http://localhost:9922/", false)
	result, err := c.getMultiAddrTransactions("auth/StdTx", "cosmos-sdk/MsgSend", "uatom", 0, -1, addrs)

	if err != nil {
		t.Error("get transactions failed!")
	} else {
		for _, tx := range result {
			fmt.Println(tx.TxID)
		}
	}
}

func Test_getBlockByHeight(t *testing.T) {
	height := uint64(164)
	c := NewClient(testUrl, true)
	result, err := c.getBlockByHeight(height)
	if err != nil {
		t.Error("get block failed!")
	} else {
		fmt.Println(result)
	}
}
