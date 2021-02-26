package goutils

import (
	"testing"
)

// 虚拟手机号码
const virtualPhoneNumber = "18888888888"

func TestPhoneValidate(t *testing.T) {
	if !MobilePhoneValidate(virtualPhoneNumber) {
		t.Fail()
	}
}
