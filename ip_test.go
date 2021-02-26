package goutils

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRemoteIp(t *testing.T) {
	go func() {
		http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("客户端的ip:", RemoteIp(r))
		})
		_ = http.ListenAndServe("0.0.0.0:7777", nil)
	}()
	resp, err := http.Get("http://localhost:7777/test")
	if err != nil {
		t.Log(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fail()
	}
	return
}
