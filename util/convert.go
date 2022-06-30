package util

import (
	"reflect"
	"strings"
	"unsafe"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

// String2Bytes convert string to bytes.
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// Bytes2String convert bytes to string.
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ConvertString2slice convert string to string slice according to sep.
func ConvertString2slice(in string, sep string) []string {
	return strings.Split(in, sep)
}

func ConvertHtml2Md(filePath, html string) error {
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(html)
	if err != nil {
		return err
	}

	return WriteToFile(filePath, markdown)
}
