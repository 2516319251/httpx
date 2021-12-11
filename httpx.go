package httpx

import (
	"io"
	"net/http"
	"strings"
	"time"
)

type (
	// Request 请求结构体
	Request struct {
		method     string
		url        string
		header     Header
		contentype ContentType
		body       io.Reader
		client     *http.Client
		err        error
	}

	// Header header 类型定义
	Header map[string]string

	// Any 请求数据类型定义
	Any map[string]interface{}
)

// NewRequest 实例化请求结构体
func NewRequest(method, url string) *Request {
	request := &Request{
		method:     strings.ToUpper(method),
		url:        url,
		contentype: JSON,
		header:     Header{},
		body:       nil,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	return request
}
