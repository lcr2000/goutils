# goutils

`golang`一些常用工具类封装

### emoji

```go
func EmojiFilter(content string) (newContent string)
```

函数`EmojiFilter`过滤掉字符串中的`Emoji` 。入参`content`为`string`类型；返回一个新的字符串，类型为`string`。



### ip

```go
func RemoteIp(req *http.Request) string
```

函数`RemoteIp`获取远程客户端的`IP`。入参`req`为`http`包中`http.Request`。



```go
func Ip2long(ipStr string) uint32
```

函数`Ip2long`将`IPv4`字符串形式转为`uint32`。



### money

```go
func FenToYuanFloat(money int) float64
```

函数`FenToYuanFloat`将金额分转换成元。入参`money`分值为分，类型为`int`；返回参数是元，类型为`float64`。



```go
func FenToYuanString(money int) string 
```

函数`FenToYuanString`将金额分转换成元。入参`money`为金额分，类型为`int`；返回参数是元，类型为`string`。



```go
func YuanToFen(money float64) int64
```

函数`YuanToFen`将金额元转换为分。入参`money`为金额元，为`float64`类型；返回参数是分，类型为`int64`。



```go
func YuanToFenRounding(money float64) int64
```

函数`YuanToFenRounding`将金额元转换为分。它与`YuanToFen`不同的是，它满足当乘以`100`后，仍然有小数位，**四舍五入法**后，再取整数部分。



```go
func DiscountCalculation(realMoney, tagPrice float64) string
```

函数`DiscountCalculation`计算两个价格的折扣。入参`realMoney`为实际价格，`tagPrice`为吊牌价格，均是`float64`类型；返回参数是入参两者对应的折扣，返回值类型为`string`（保留一位小数）。



```go
func DiscountCalculationInt(realMoney, tagPrice int) string
```

函数`DiscountCalculation`计算两个价格的折扣。与`DiscountCalculation`不同的是传入的参数类型。



### random

**注意：**使用random包内的方法时，请确保项目初始化时设置随机数种子，如：`rand.Seed(time.Now().UnixNano())`。



```go
func GetRandomIp() string
```

函数`GetRandomIp`获取国内随机`Ip`。返回参数`Ip`为`string`类型。



```go
func GetRandomString(n int) string
```

函数`GetRandomString`随机生成大写字母和数字的组合。入参`n`是生成字符串的长度，类型为`int`；返回参数是随机数字，类型为`string`。



### time

```go
func TimeFormatter(i interface{}, format string) string
```

函数`TimeFormatter`为时间格式化函数，将对应时间进行格式化。
入参`i`为时间，请传入`time`、`int`、`int32`、`int64`类型；
入参`format`为需要格式化的日期格式，如：`2006-01-02 15:04:05`、`2006.01.02`等等；返回参数为格式化后的时间，类型为`string`。



```go
func GetTodayStartTime(formatLayout string) time.Time
```

函数`GetTodayStartTime`获取本日起始时间。
入参`formatLayout`为格式化格式，类型为`string`；返回参数是今日起始时间，类型为`time.Time`。



```go
func IsYesterday(t int64) bool
```

函数`IsYesterday`判断传入的时间`t`是否为昨天。返回参数为`bool`类型。



```go
func IsTomorrow(t int64) bool
```

函数`IsTomorrow`判断传入的时间`t`是否为明天。返回参数为`bool`类型。



```go
func IsToday(t int64) bool
```

函数`IsToday`判断传入的时间`t`是否为今天。返回参数为`bool`类型。



### verify

```go
func MobilePhoneValidate(phoneNumber string) bool
```

函数`MobilePhoneValidate`校验传入的手机号是否正确。入参`phoneNumber`为手机号，类型为`string`；返回参数为`bool`类型。



```go
func IdCardValidate(idCardNumber string) bool
```

函数`IdCardValidate`校验传入的身份证号码是否正确。入参`idCardNumber`为身份证号码，类型为`string`；返回参数为`bool`类型。