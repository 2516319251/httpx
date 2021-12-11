package httpx

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

// Get 发起 get 请求
func Get(url string) *Request {
	return NewRequest(http.MethodGet, url)
}

// Post 发起 post 请求
func Post(url string) *Request {
	return NewRequest(http.MethodPost, url)
}

// Put 发起 put 请求
func Put(url string) *Request {
	return NewRequest(http.MethodPut, url)
}

// Delete 发起 delete 请求
func Delete(url string) *Request {
	return NewRequest(http.MethodDelete, url)
}

// Header 设置 header
func (r *Request) Header(header Header) *Request {
	r.header = header
	return r
}

// ContentType 设置 Content-Type，默认 application/json; charset=utf-8
func (r *Request) ContentType(contentype ContentType) *Request {
	r.contentype = contentype
	return r
}

// encode 进行 urlencode 编码
func (r *Request) encode(data Any) string {
	s := url.Values{}
	for k, v := range data {
		s.Set(k, fmt.Sprintf("%v", v))
	}
	return s.Encode()
}

// Query 设置 url 中的参数（to url encode）
func (r *Request) Query(query Any) *Request {
	r.url = fmt.Sprintf("%s?%s", r.url, r.encode(query))
	return r
}

// Body 设置请求 body
func (r *Request) Body(body Any) *Request {
	buf := new(bytes.Buffer)

	switch r.contentype {
	case JSON:
		e := json.NewEncoder(buf)
		e.SetEscapeHTML(false)
		if err := e.Encode(body); err != nil {
			r.err = err
		}
	case Form:
		if _, err := buf.WriteString(r.encode(body)); err != nil {
			r.err = err
		}
	case XML:
		if err := xml.NewEncoder(buf).Encode(body); err != nil {
			r.err = err
		}
	}

	r.body = buf
	return r
}

// Send 发起请求
func (r *Request) Send() *Response {
	// 初始化 response
	resp := &Response{}

	// 如果发起请求的过程中存在错误
	if r.err != nil {
		resp.err = r.err
		return resp
	}

	// 生成请求信息
	req, nerr := http.NewRequest(r.method, r.url, r.body)
	if nerr != nil {
		resp.err = nerr
		return resp
	}

	// 设置 header
	req.Header.Set("Content-Type", r.contentype.String())
	for k, v := range r.header {
		req.Header.Set(k, v)
	}

	// 发起请求
	res, derr := r.client.Do(req)
	if derr != nil {
		resp.err = derr
		return resp
	}

	// 返回
	resp.Response = res
	return resp
}
