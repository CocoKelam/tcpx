package main

import (
	"fmt"
	"log"

	"github.com/CocoKelam/tcpx"
	//"tcpx"
	"time"
)

func main() {
	srv := tcpx.NewTcpX(nil)

	// send client heartbeat
	//srv.HeartBeatModeDetail(true, 5*time.Second, true, 5*time.Second, false, tcpx.DEFAULT_HEARTBEAT_MESSAGEID)

	// donot send client heartbeat
	srv.ClientHBModeDetail(true, 5*time.Second, true, tcpx.DEFAULT_HEARTBEAT_MESSAGEID)
	srv.HBToClientMode(true, 5*time.Second, func(c *tcpx.Context) {
		if err := c.JSON(10088, nil, nil); err != nil {
			log.Println(fmt.Sprintf("send heartbeat to client '%s'", c.ClientIP()))
		}
	})

	//srv.RewriteHeartBeatHandler(1300, func(c *tcpx.Context) {
	//	fmt.Println("rewrite heartbeat handler")
	//	c.RecvHeartBeat()
	//})

	tcpx.SetLogMode(tcpx.DEBUG)

	srv.ListenAndServe("tcp", ":8101")
}
