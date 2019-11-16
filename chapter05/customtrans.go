package main

import (
	"io"
	"net/http"
	"os"
)

type OurCustomTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *OurCustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 处理一些事情 ...
	// 发起HTTP请求，并接受传入的*Request请求值作为参数并返回*Response响应值
	// 不应该试图解析http响应的信息
	// 添加一些域到req.Header中
	return t.transport().RoundTrip(req)
}

func (t *OurCustomTransport) Client() *http.Client {
	return &http.Client{
		Transport: t,
	}
}

func main() {
	/*
		resp, err := http.Get("http://example.com/")
		if err != nil {
			// 处理错误...
			return
		}
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	*/

	t := &OurCustomTransport{}
	c := t.Client()
	resp, err := c.Get("http://example.com/")

	http.ListenAndServe()
}
