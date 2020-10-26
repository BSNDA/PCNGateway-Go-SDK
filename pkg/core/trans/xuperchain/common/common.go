// Copyright (c) 2019. Baidu Inc. All Rights Reserved.

// package common is related to common variables and utils funcs
package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	walletRand "github.com/xuperchain/crypto/core/hdwallet/rand"
)

const (
	// TxVersion tx version number
	TxVersion = 1
)

// GetNonce get nonce value
func GetNonce() string {
	return fmt.Sprintf("%d%8d", time.Now().Unix(), rand.Intn(100000000))
}

// SetSeed set seed
func SetSeed() error {
	// 生成加强版的随机熵
	seedByte, err := walletRand.GenerateSeedWithStrengthAndKeyLen(walletRand.KeyStrengthHard, walletRand.KeyLengthInt64)
	if err != nil {
		return err
	}

	// 生成大随机数
	bytesBuffer := bytes.NewBuffer(seedByte)
	var seed int64
	binary.Read(bytesBuffer, binary.BigEndian, &seed)

	// 设置随机数生成器的随机源
	rand.Seed(seed)

	return nil
}

// PathExistsAndMkdir judge whether path is existant or not
func PathExistsAndMkdir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// IsValidAmount judge whether the number is legal
func IsValidAmount(amount string) (string, bool) {
	if amount == "" {
		amount = "0"
		return amount, true
	}

	amountInt64, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Printf("Transfer amount to int64 err: %v", err)
		return "", false
	}

	if amountInt64 < 0 {
		log.Printf("Transfer amount is negative")
		return "", false
	}

	return amount, true
}
