package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Echo synchronously.
	out = frame
	return

	/*
	   // Echo asynchronously.
	   data := append([]byte{}, frame...)
	   go func() {
	       time.Sleep(time.Second)
	       c.AsyncWrite(data)
	   }()
	   return
	*/
}

func main() {
	var port int
	var multicore, reuseport bool

	flag.IntVar(&port, "port", 25565, "--port 25565")
	flag.BoolVar(&multicore, "multicore", true, "--multicore true")
	flag.BoolVar(&reuseport, "reuseport", true, "--reuseport true")
	flag.Parse()

	echo := new(echoServer)
	log.Fatal(gnet.Serve(echo, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore), gnet.WithReusePort(reuseport)))
}
