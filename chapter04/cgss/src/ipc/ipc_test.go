package ipc

import "testing"

type EchoServer struct {
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{
		Code: "OK",
		Body: "ECHO: " + method + " ~ " + params,
	}
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("foo", "From Client1")
	resp2, _ := client2.Call("foo", "From Client2")

	if resp1.Body != "ECHO: foo ~ From Client1" || resp2.Body != "ECHO: foo ~ From Client2" {
		t.Error("IpcClient.Call falied. resp1:", resp1, "resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
