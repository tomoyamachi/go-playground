package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestInt64(t *testing.T) {
	tests := []struct {
		input []byte
		want  Ali
	}{
		{
			input: []byte(`{"I64":10,"T":"2022-10-18T00:00:00Z","AI64":10,"AT1":1666051200,"AT2":"2022-10-18T00:00:00Z"}`),
			want: Ali{
				I64:  10,
				T:    time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC),
				AI64: 10,
				AT1:  AliasTime(time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)),
				AT2:  AliasTime(time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)),
			},
		},
	}
	for _, tt := range tests {
		var ali Ali
		if err := json.Unmarshal(tt.input, &ali); err != nil {
			t.Fatal("unmarshal error", err)
		}
		if ali.I64 != tt.want.I64 {
			t.Error("I64 not match")
		}
		if ali.T != tt.want.T {
			t.Error("T not match")
		}

		if ali.AI64 != tt.want.AI64 {
			t.Error("AI64 not match")
		}
		if ali.AT1 != tt.want.AT1 {
			t.Error("AT1 not match", ali.AT1, tt.want.AT1)
		}
		if ali.AT2 != tt.want.AT2 {
			t.Error("AT2 not match")
		}
	}

}
