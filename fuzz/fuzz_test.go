//go:build go1.23
// +build go1.23

package fuzz

import (
	"net/url"
	"reflect"
	"testing"
)

func FuzzParseQuery(f *testing.F) {
	f.Add("x=1&y=2")
	f.Fuzz(func(t *testing.T, queryStr string) {
		query, err := url.ParseQuery(queryStr)
		if err != nil {
			t.Skip()
		}
		queryStr2 := query.Encode()
		query2, err := url.ParseQuery(queryStr2)
		if err != nil {
			t.Fatalf("ParseQuery failed to decode a valid encoded query %s: %v", queryStr2, err)
		}
		if !reflect.DeepEqual(query, query2) {
			t.Errorf("ParseQuery gave different query after being encoded\nbefore: %v\nafter: %v", query, query2)
		}
	})
}

// func FuzzArgsMismatch(f *testing.F) {
// 	f.Add("1") // ファズターゲットの引数と型や個数が異なる場合はエラー
// 	f.Fuzz(func(t *testing.T, s string) {
// 	})
// }
