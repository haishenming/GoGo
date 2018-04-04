package ttStruct

import "net/textproto"

type SuperInt int

func (a SuperInt) Less(b SuperInt) bool {
	return a < b
}

func (a *SuperInt) Add(b SuperInt) {
	*a += b
}

type Header map[string][] string

// Add()方法用于添加一个键值对到HTTP头部
func (h Header) Add(key, value string) {
	textproto.MIMEHeader(h).Add(key, value)
}

// Set()方法用于设置某个键对应的值,如果该键已存在,则替换已存在的值
func (h Header) Set(key, value string) {
	textproto.MIMEHeader(h).Set(key, value)
}