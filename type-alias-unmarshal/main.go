package main

import (
	"encoding/json"
	"strconv"
	"time"
)

type Ali struct {
	I64  int64
	T    time.Time
	AI64 AliasInt64
	AT1  AliasTime
	AT2  AliasTime
}

type AliasInt64 int64

type AliasTime time.Time

func (ai *AliasInt64) UnmarshalJSON(data []byte) error {
	parsed, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*ai = AliasInt64(parsed)
	return nil
}

func (at *AliasTime) UnmarshalJSON(data []byte) error {
	// 数値型であればUnixTimeとしてロード
	parsed, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		//	time.Timeとしてパース
		var tt time.Time
		if err := json.Unmarshal(data, &tt); err != nil {
			return err
		}
		*at = AliasTime(tt)
		return nil
	}
	tt := time.Unix(parsed, 0).In(time.UTC)
	*at = AliasTime(tt)
	return nil
}
