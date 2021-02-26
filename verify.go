package goutils

import (
	"regexp"
)

// 校验正则表达式
const (
	phone  = "^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[01235678]|18[0-9]|19[189])\\d{8}$"
	idCard = "(^[1-9]\\d{5}(18|19|([23]\\d))\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$)|(^[1-9]\\d{5}\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{2}[0-9Xx]$)"
)

// MobilePhoneValidate校验传入的手机号是否正确.
func MobilePhoneValidate(phoneNumber string) bool {
	reg := regexp.MustCompile(phone)
	return reg.MatchString(phoneNumber)
}

// IdCardValidate校验传入的身份证号码是否正确.
func IdCardValidate(idCardNumber string) bool {
	reg := regexp.MustCompile(idCard)
	return reg.MatchString(idCardNumber)
}
