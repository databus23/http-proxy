package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {

	var port = flag.Int("port", 7070, "port to bind to")
	var verbose = flag.Bool("verbose", false, "verbose")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		if *verbose {
			fmt.Println("Accepting CONNECT to", host)
		}
		return goproxy.OkConnect, host
	})
	proxy.Verbose = *verbose
	bindAddr := fmt.Sprintf("127.0.0.1:%d", *port)
	log.Printf("Starting proxy server on %s", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, proxy))
}
