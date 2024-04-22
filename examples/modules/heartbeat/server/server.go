package main

import (
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

	//srv.RewriteHeartBeatHandler(1300, func(c *tcpx.Context) {
	//	fmt.Println("rewrite heartbeat handler")
	//	c.RecvHeartBeat()
	//})

	tcpx.SetLogMode(tcpx.DEBUG)

	srv.ListenAndServe("tcp", ":8101")
}
