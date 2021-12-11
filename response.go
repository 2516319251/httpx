package httpx

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Response 响应处理结构体
type Response struct {
	*http.Response
	err error
}

// Error 获取请求错误信息
func (resp *Response) Error() error {
	return resp.err
}

// GetBytes 获取响应的 body 内容
func (resp *Response) GetBytes() ([]byte, error) {
	if resp.err != nil {
		return nil, resp.err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// BindJson 绑定响应到 tag 为 json 的结构体
func (resp *Response) BindJson(v interface{}) error {
	if resp.err != nil {
		return resp.err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}

// BindXml 绑定响应到 tag 为 xml 的结构体
func (resp *Response) BindXml(v interface{}) error {
	if resp.err != nil {
		return resp.err
	}
	defer resp.Body.Close()

	return xml.NewDecoder(resp.Body).Decode(v)
}
