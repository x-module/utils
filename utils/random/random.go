// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package random implements some basic functions to generate random int and string.
package random

import (
	crand "crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"strings"
	"time"
)

const (
	NUMERAL       = "0123456789"
	LOWER_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPER_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LETTERS       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandInt generate random int between min and max, maybe min,  not be max.
// Play: https://go.dev/play/p/pXyyAAI5YxD
func RandInt(min, max int) int {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(max-min) + min
}

// RandBytes generate random byte slice.
// Play: https://go.dev/play/p/EkiLESeXf8d
func RandBytes(length int) []byte {
	if length < 1 {
		return []byte{}
	}
	b := make([]byte, length)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return nil
	}

	return b
}

// RandString generate random string of specified length.
// Play: https://go.dev/play/p/W2xvRUXA7Mi
func RandString(length int) string {
	return random(LETTERS, length)
}

// RandUpper generate a random upper case string.
// Play: https://go.dev/play/p/29QfOh0DVuh
func RandUpper(length int) string {
	return random(UPPER_LETTERS, length)
}

// RandLower generate a random lower case string.
// Play: https://go.dev/play/p/XJtZ471cmtI
func RandLower(length int) string {
	return random(LOWER_LETTERS, length)
}

// RandNumeral generate a random numeral string of specified length.
// Play: https://go.dev/play/p/g4JWVpHsJcf
func RandNumeral(length int) string {
	return random(NUMERAL, length)
}

// RandNumeralOrLetter generate a random numeral or letter string.
// Play: https://go.dev/play/p/19CEQvpx2jD
func RandNumeralOrLetter(length int) string {
	return random(NUMERAL+LETTERS, length)
}

// random generate a random string based on given string range.
func random(s string, length int) string {
	b := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range b {
		b[i] = s[r.Int63()%int64(len(s))]
	}

	return string(b)
}

// UUIdV4 generate a random UUID of version 4 according to RFC 4122.
// Play: https://go.dev/play/p/_Z9SFmr28ft
func UUIdV4() (string, error) {
	uuid := make([]byte, 16)

	n, err := io.ReadFull(crand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// GetUUIdByTime 根据时间生成UUID true去除“-”，false不去除
func GetUUIdByTime(flag bool) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	if flag {
		return strings.Replace(id.String(), "-", "", -1), nil
	}
	return id.String(), err
}

// IdUUIdByRand V4 基于随机数 true去除“-”，false不去除
func IdUUIdByRand(flag bool) string {
	u4 := uuid.New()
	if flag {
		return strings.Replace(u4.String(), "-", "", -1)
	}
	return u4.String()
}
