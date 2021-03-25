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

package cryptochain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
	owcrypt "github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/v2/openwallet"
)


var (
	CRO_mainnetAddress = addressEncoder.AddressType{"bech32", addressEncoder.ATOMBech32Alphabet, "cro", "h160", 20, nil, nil}
	CRO_testnetAddress = addressEncoder.AddressType{"bech32", addressEncoder.ATOMBech32Alphabet, "tcro", "h160", 20, nil, nil}
)

type AddressDecoderV2 struct {
	openwallet.AddressDecoderV2Base
	wm *WalletManager //钱包管理者
}
//type addressDecoder struct {
//	wm *WalletManager //钱包管理者
//}

//NewAddressDecoder 地址解析器
func NewAddressDecoderV2(wm *WalletManager) *AddressDecoderV2 {
	decoder := AddressDecoderV2{
		AddressDecoderV2Base: openwallet.AddressDecoderV2Base{},
		wm:                   wm,
	}
	return &decoder
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := CRO_mainnetAddress
	if dec.wm.Config.IsTestNet {
		cfg = CRO_testnetAddress
	}

	pkHash := owcrypt.Hash(hash, 32, owcrypt.HASH_ALG_HASH160)

	address := addressEncoder.AddressEncode(pkHash, cfg)
	return address, nil
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {
	cfg := addressEncoder.AddressType{}
	if strings.Index(addr, CRO_mainnetAddress.ChecksumType) == 0 {
		cfg = CRO_mainnetAddress
	} else if strings.Index(addr, CRO_testnetAddress.ChecksumType) == 0 {
		cfg = CRO_testnetAddress
	} else {
		return nil, errors.New("invalid bech32 prefix")
	}

	decodeHash, err := addressEncoder.AddressDecode(addr, cfg)
	if err != nil {
		return nil, err
	}
	return decodeHash, nil
}

// AddressVerify 地址校验
func (dec *AddressDecoderV2) AddressVerify(address string, opts ...interface{}) bool {
	cfg := addressEncoder.AddressType{}
	if strings.Index(address, CRO_mainnetAddress.ChecksumType) == 0 {
		cfg = CRO_mainnetAddress
	} else if strings.Index(address, CRO_testnetAddress.ChecksumType) == 0 {
		cfg = CRO_testnetAddress
	} else {
		return false
	}

	_, err := addressEncoder.AddressDecode(address, cfg)
	if err != nil {
		return false
	}
	return true
}

//PrivateKeyToWIF 私钥转WIF
func (dec *AddressDecoderV2) PrivateKeyToWIF(priv []byte, isTestnet bool) (string, error) {
	return "", fmt.Errorf("PrivateKeyToWIF not implement")
}

//PublicKeyToAddress 公钥转地址
func (dec *AddressDecoderV2) PublicKeyToAddress(pub []byte, isTestnet bool) (string, error) {

	cfg := CRO_mainnetAddress
	if isTestnet {
		cfg = CRO_testnetAddress
	}

	pkHash := owcrypt.Hash(pub, 32, owcrypt.HASH_ALG_HASH160)

	address := addressEncoder.AddressEncode(pkHash, cfg)
	return address, nil
}

//WIFToPrivateKey WIF转私钥
func (dec *AddressDecoderV2) WIFToPrivateKey(wif string, isTestnet bool) ([]byte, error) {
	return nil, fmt.Errorf("WIFToPrivateKey not implement")
}

//RedeemScriptToAddress 多重签名赎回脚本转地址
func (dec *AddressDecoderV2) RedeemScriptToAddress(pubs [][]byte, required uint64, isTestnet bool) (string, error) {
	return "", fmt.Errorf("RedeemScriptToAddress not implement")
}

// CustomCreateAddress 创建账户地址
func (dec *AddressDecoderV2) CustomCreateAddress(account *openwallet.AssetsAccount, newIndex uint64) (*openwallet.Address, error) {
	return nil, fmt.Errorf("CreateAddressByAccount not implement")
}

// SupportCustomCreateAddressFunction 支持创建地址实现
func (dec *AddressDecoderV2) SupportCustomCreateAddressFunction() bool {
	return false
}
