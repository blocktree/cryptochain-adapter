/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"testing"

	"github.com/blocktree/openwallet/v2/openw"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "123456", contract, nil)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract) ([]*openwallet.RawTransactionWithError, error) {

	rawTxArray, err := tm.CreateSummaryRawTransactionWithError(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract, nil)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WJBd5NWkYpWvb4Mhga5SKUHevDSRktNQrJ"
	accountID := "8BZiYkQFPjbeqasYiJXnjpZxBxrYjLNyytGuqFs3SaR8"
	to := "cro14ywgrs9eqkavxkr7d8fhyvanpqluxe98qrl0vj"
	

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "5", "", nil)
	if err != nil {
		return
	}

	log.Std.Info("rawTx: %+v", rawTx)

	_, err = testSignTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testVerifyTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testSubmitTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	//rawTx, err = testCreateTransactionStep(tm, walletID, accountID, to, "0.02", "", nil)
	//if err != nil {
	//	return
	//}
	//
	//log.Std.Info("rawTx: %+v", rawTx)
	//
	//_, err = testSignTransactionStep(tm, rawTx)
	//if err != nil {
	//	return
	//}
	//
	//_, err = testVerifyTransactionStep(tm, rawTx)
	//if err != nil {
	//	return
	//}
	//
	//_, err = testSubmitTransactionStep(tm, rawTx)
	//if err != nil {
	//	return
	//}
	//
	//rawTx, err = testCreateTransactionStep(tm, walletID, accountID, to, "0.06", "", nil)
	//if err != nil {
	//	return
	//}
	//
	//log.Std.Info("rawTx: %+v", rawTx)
	//
	//_, err = testSignTransactionStep(tm, rawTx)
	//if err != nil {
	//	return
	//}
	//
	//_, err = testVerifyTransactionStep(tm, rawTx)
	//if err != nil {
	//	return
	//}
	//
	//_, err = testSubmitTransactionStep(tm, rawTx)
	//if err != nil {
	//	return
	//}
	//
	//rawTx, err = testCreateTransactionStep(tm, walletID, accountID, to, "0.06", "", nil)
	//if err != nil {
	//	return
	//}
	//
	//log.Std.Info("rawTx: %+v", rawTx)
	//
	//_, err = testSignTransactionStep(tm, rawTx)
	//if err != nil {
	//	return
	//}
	//
	//_, err = testVerifyTransactionStep(tm, rawTx)
	//if err != nil {
	//	return
	//}
	//
	//_, err = testSubmitTransactionStep(tm, rawTx)
	//if err != nil {
	//	return
	//}
	////	getdata("cosmos1z9k73l7trgshqpgg7m6hk9ehe4gphea5ch9dyh")

}

func TestSummary(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WCPa2UFFUt5HAXPHbzwWty5x3NJpEeBRHY"
	accountID := "8WivbG5nnxGEn9kruuWa7NKQWWSyzjTuUhZMJaExC5Pq"
	summaryAddress := "cosmos1xv66sa5tlplm68j4fec6stdzszg3pcvswag06j"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//??????????????????
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}
