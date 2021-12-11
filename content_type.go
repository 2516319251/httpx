package httpx

type ContentType uint

const (
	JSON ContentType = iota
	Form
	XML
)

var contentype = map[ContentType]string{
	JSON: "application/json; charset=utf-8",
	Form: "application/x-www-form-urlencoded",
	XML:  "application/xml",
}

// ContentTypeInterface 接口定义
type ContentTypeInterface interface {
	String() string
}

// String 获取 Content-Type 字符串
func (c ContentType) String() string {
	t, ok := contentype[c]
	if !ok {
		t = contentype[JSON]
	}
	return t
}
