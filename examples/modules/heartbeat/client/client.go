package main

import (
	"fmt"
	"net"
	"os"
	"sync"

	"time"

	"github.com/CocoKelam/tcpx"
)

func main() {

	var locker sync.Mutex

	conn, e := net.Dial("tcp", "localhost:8101")
	if e != nil {
		panic(e)
	}
	var heartBeat []byte
	heartBeat, e = tcpx.PackWithMarshaller(tcpx.Message{
		MessageID: tcpx.DEFAULT_HEARTBEAT_MESSAGEID,
		Header:    nil,
		Body:      nil,
	}, nil)

	// send heartbeat

	// sendHeartbeatTicker := time.NewTicker(4*time.Second + 500*time.Millisecond)
	// svrHeartbeatTimeTicker :=

	// send heartbeat
	go func() {
		for {
			_, e = conn.Write(heartBeat)
			if e != nil {
				fmt.Println(e.Error())
				break
			}
			time.Sleep(4*time.Second + 500*time.Millisecond)
		}
	}()

	svrActively := time.Now()
	go func() {
		var buf = make([]byte, 500)
		for {
			buf, e = tcpx.FirstBlockOf(conn)
			if e != nil {
				fmt.Println(e.Error())
				os.Exit(0)
				break
			}
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "-", buf)

			locker.Lock()
			svrActively = time.Now()
			locker.Unlock()
		}
	}()

	//check server heartbeat
	sendHeartbeatTicker := time.NewTicker(1 * time.Second)
	for _ = range sendHeartbeatTicker.C {
		locker.Lock()
		if time.Since(svrActively) > 5*time.Second {
			conn.Close()
			return
		}
		locker.Unlock()
	}
}
