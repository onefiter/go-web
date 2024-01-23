package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

type Context struct {
	Req        *http.Request
	Resp       http.ResponseWriter
	PathParams map[string]string

	queryValues url.Values

	// cookieSameSite http.SameSite
}

// 用户每次都得自己检测是不是 500，然后调这个方法
func (c *Context) ErrPage() {

}

func (c *Context) SetCookie(ck *http.Cookie) {
	// 不推荐
	// ck.SameSite = c.cookieSameSite
	http.SetCookie(c.Resp, ck)
}

func (c *Context) RespJSONOK(val any) error {
	return c.RespJSON(http.StatusOK, val)
}

func (c *Context) RespJSON(status int, val any) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}

	c.Resp.WriteHeader(status)

	n, err := c.Resp.Write(data)

	if n != len(data) {
		return errors.New("web: 未写入全部数据")
	}
	return err

}

// 解决大多数人的需求
func (c *Context) BindJSON(val any) error {

	if c.Req.Body == nil {
		return errors.New("web: body 为 nil")
	}
	// bs, _:= io.ReadAll(c.Req.Body)
	// json.Unmarshal(bs, val)
	decoder := json.NewDecoder(c.Req.Body)
	// useNumber => 数字就是用 Number 来表示
	// 否则默认是 float64
	// if jsonUseNumber {
	// 	decoder.UseNumber()
	// }

	// 如果要是有一个未知的字段，就会报错
	// 比如说你 User 只有 Name 和 Email 两个字段
	// JSON 里面额外多了一个 Age 字段，那么就会报错
	// decoder.DisallowUnknownFields()
	return decoder.Decode(val)
}

// FormValue(key1)
// FormValue(key2)
func (c *Context) FormValue(key string) (string, error) {
	err := c.Req.ParseForm()
	if err != nil {
		return "", err
	}
	return c.Req.FormValue(key), nil
}

// Query 和表单比起来，它没有缓存
func (c *Context) QueryValue(key string) (string, error) {

	if c.queryValues == nil {
		c.queryValues = c.Req.URL.Query()
	}

	vals, ok := c.queryValues[key]
	if !ok {
		return "", errors.New("web: key 不存在")
	}
	return vals[0], nil

	// 用户区别不出来是真的有值，但是值恰好是空字符串
	// 还是没有值
	// return c.queryValues.Get(key), nil
}

func (c *Context) QueryValueV1(key string) StringValue {

	if c.queryValues == nil {
		c.queryValues = c.Req.URL.Query()
	}

	vals, ok := c.queryValues[key]
	if !ok {
		return StringValue{
			err: errors.New("web: key 不存在"),
		}
	}
	return StringValue{
		val: vals[0],
	}

	// 用户区别不出来是真的有值，但是值恰好是空字符串
	// 还是没有值
	// return c.queryValues.Get(key), nil
}

func (c *Context) PathValueV1(key string) StringValue {
	val, ok := c.PathParams[key]
	if !ok {
		return StringValue{
			err: errors.New("web: key 不存在"),
		}
	}
	return StringValue{
		val: val,
	}
}

func (c *Context) PathValue(key string) (string, error) {
	val, ok := c.PathParams[key]
	if !ok {
		return "", errors.New("web: key 不存在")
	}
	return val, nil
}

// 这种泛型不行，因为在创建的时候我们不知道用户需要什么作为 T
// type StringValue[T any] struct {
// 	val string
// 	err error
// }

type StringValue struct {
	val string
	err error
}

func (s StringValue) AsInt64() (int64, error) {
	if s.err != nil {
		return 0, s.err
	}
	return strconv.ParseInt(s.val, 10, 64)
}
