package tcpx

import "net"

type ConnEx struct {
	net.Conn
	data interface{}
}

func NewConnEx(conn net.Conn) *ConnEx {
	return &ConnEx{
		Conn: conn,
	}
}

func (c *ConnEx) GetData() interface{} {
	return c.data
}

func (c *ConnEx) SetData(data interface{}) {
	c.data = data
}
