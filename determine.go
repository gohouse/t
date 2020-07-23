package t

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"regexp"
	"time"
)

// iDetermine 判断
type iDetermine interface {
	IsNumeric() bool                        // 是否数字
	IsInteger() bool                        // 是否整数
	IsFloat() bool                          // 是否浮点数
	IsZero() bool                           // 是否零值
	IsChineseCharacters() bool              // 是否汉字
	IsChineseName() bool                    // 是否中文名字
	IsHost() bool                           // 是否域名
	IsUrl() bool                            // 是否互联网url地址
	IsEmail() bool                          // 是否邮箱地址
	IsChineseMobile() bool                  // 是否中国手机号
	IsDate() bool                           // 是否常用的日期格式
	IsDateTime() bool                       // 是否常用的日期时间格式
	IsIpV4() bool                           // 是否ipv4地址
	IsIpV6() bool                           // 是否ipv6地址
	IsIp() bool                             // 是否ip地址
	IsChineseID() bool                      // 是否中国大陆身份证号码
	IsXml() bool                            // 是否xml
	IsJson() bool                           // 是否json
	IsJsonMap() bool                        // 是否是json对象
	IsJsonSlice() bool                      // 是否是json数组
	IsBetween(min, max interface{}) bool    // 是否在两数之间
	IsBetweenFloat64(min, max float64) bool // 是否在两个浮点数之间
	IsBetweenAlpha(min, max string) bool    // 是否在两个字符之间
}

// IsInteger 是否为整数
func (t Type) IsInteger() bool {
	return t.String() == New(t.Int64()).String()
}

// IsNumeric 是否为数字,包含整数和小数
func (t Type) IsNumeric() bool {
	return t.String() == New(t.Float64()).String()
}

// IsFloat 是否为float
func (t Type) IsFloat() bool {
	switch t.val.(type) {
	case float64, float32:
		return true
	default:
		return t.IsNumeric()
	}
}

// IsJsonSlice 是否json数组
func (t Type) IsJsonSlice() bool {
	var js []interface{}
	return json.Unmarshal(t.Bytes(), &js) == nil
}

// IsJsonMap 是否json对象
func (t Type) IsJsonMap() bool {
	var js = map[string]interface{}{}
	return json.Unmarshal(t.Bytes(), &js) == nil
}

func (t Type) IsZero() bool {
	return reflect.ValueOf(t.val).IsZero()
}

// IsChineseCharacters 是否汉字
func (t Type) IsChineseCharacters() bool {
	exp := regexp.MustCompile("^[\u4e00-\u9fa5]*$")
	return exp.Match(t.Bytes())
}

// IsChineseCharacters 是否中国的名字
func (t Type) IsChineseName() bool {
	exp := regexp.MustCompile("^[\u4e00-\u9fa5]+[[·•●][\u4e00-\u9fa5]+]*$")
	return exp.Match(t.Bytes())
}

// IsEmail 是否邮箱地址
func (t Type) IsEmail() bool {
	exp := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	return exp.Match(t.Bytes())
}

// IsEmail 是否中国手机号
func (t Type) IsChineseMobile() bool {
	exp := regexp.MustCompile(`^1[3456789]\d{9}$`)
	return exp.Match(t.Bytes())
}

// IsHost 是否域名
func (t Type) IsHost() bool {
	exp := regexp.MustCompile(`[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(/.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+/.?`)
	return exp.Match(t.Bytes())
}

// IsUrl 是否域名
func (t Type) IsUrl() bool {
	var strRegex = `(ht|f)tp(s?)\:\/\/[0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*(:(0-9)*)*(\/?)([a-zA-Z0-9\-\.\?\,\'\/\\\+&amp;%\$#_]*)?` //请求参数结尾- 英文或数字和[]内的各种字符
	exp := regexp.MustCompile(strRegex)
	return exp.Match(t.Bytes())
}

// IsChineseID 是否中国大陆身份证号码
func (t Type) IsChineseID() bool {
	exp := regexp.MustCompile(`^([0-9]){7,18}(x|X)?$`)
	return exp.Match(t.Bytes())
}

// IsDate 是否日期 yyyy-mm-dd 格式
func (t Type) IsDate() bool {
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("2006-01-02", t.String(), loc)
	return t1.Format("2006-01-02") == t.String()
}

// IsDateTime 是否日期时间 yyyy-mm-dd HH:mm:ss 格式
func (t Type) IsDateTime() bool {
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", t.String(), loc)
	return t1.Format("2006-01-02 15:04:05") == t.String()
}

// IsIpV4 是否ipv4
func (t Type) IsIpV4() bool {
	exp := regexp.MustCompile(`\\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\b`)
	return exp.Match(t.Bytes())
}

// IsIpV6 是否ipv6
func (t Type) IsIpV6() bool {
	exp := regexp.MustCompile(`(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`)
	return exp.Match(t.Bytes())
}

// IsIp 是否ip
func (t Type) IsIp() bool {
	return t.IsIpV4() || t.IsIpV6()
}

// IsXml 是否xml
func (t Type) IsXml() bool {
	return xml.Unmarshal(t.Bytes(), new(interface{})) == nil
}

// IsJson 是否json
func (t Type) IsJson() bool {
	var x struct{}
	return json.Unmarshal(t.Bytes(), &x) == nil
}

// Between 是否在给定的两个整数之间
func (t Type) IsBetween(min, max interface{}) bool {
	var mi = New(min)
	var ma = New(max)
	if New(min).IsNumeric() {
		return t.IsBetweenFloat64(mi.Float64(), ma.Float64())
	}
	return t.IsBetweenAlpha(mi.String(), ma.String())
}

// BetweenFloat64 是否在给定的两个浮点数字之间
func (t Type) IsBetweenFloat64(min, max float64) bool {
	return min <= t.Float64() && t.Float64() <= max
}

// BetweenFloat64 是否在给定的两个字母之间
func (t Type) IsBetweenAlpha(min, max string) bool {
	return min <= t.String() && t.String() <= max
}
