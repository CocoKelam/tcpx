package main

import (
	"fmt"
	"net"

	"github.com/CocoKelam/tcpx"
	"github.com/CocoKelam/tcpx/examples/modules/my-marshal/marshaller"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:7011")

	if e != nil {
		panic(e)
	}

	var payload = []byte(`hello`)
	buf, e := tcpx.PackWithMarshaller(tcpx.Message{
		MessageID: 22,
		Header:    nil,
		Body:      payload,
	}, marshaller.ByteMarshaller{})

	_, e = conn.Write(buf)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	buf, e = tcpx.PackWithMarshaller(tcpx.Message{
		MessageID: 23,
		Header:    nil,
		Body:      payload,
	}, marshaller.ByteMarshaller{})

	_, e = conn.Write(buf)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	buf, e = tcpx.PackWithMarshaller(tcpx.Message{
		MessageID: 24,
		Header:    nil,
		Body:      payload,
	}, marshaller.ByteMarshaller{})

	_, e = conn.Write(buf)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	select {}
}
