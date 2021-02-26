package goutils

import (
	"fmt"
	"math/rand"
)

/*
	使用random包内的方法时，请确保在项目初始化时设置随机数种子.
	如：rand.Seed(time.Now().UnixNano()).
*/

var IpLong = [][]int{
	{607649792, 608174079},
	{975044608, 977272831},
	{999751680, 999784447},
	{1019346944, 1019478015},
	{1038614528, 1039007743},
	{1783627776, 1784676351},
	{1947009024, 1947074559},
	{1987051520, 1988034559},
	{2035023872, 2035154943},
	{2078801920, 2079064063},
	{-1950089216, -1948778497},
	{-1425539072, -1425014785},
	{-1236271104, -1235419137},
	{-770113536, -768606209},
	{-569376768, -564133889},
}

// GetRandomIp生成国内随机Ip.
func GetRandomIp() string {
	r := rand.Intn(len(IpLong))
	return long2ip(uint32(RandInt64(int64(IpLong[r][0]), int64(IpLong[r][1]))))
}

func long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// GetRandomString随机生成大写字母和数字的组合.
// n是生成字符串的长度.
func GetRandomString(n int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

