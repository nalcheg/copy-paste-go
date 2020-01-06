package main

import "testing"

func TestDoRequest(t *testing.T) {
	r, err := DoRequest("https://yobit.net/api/2/ltc_btc/ticker")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(r)
}
