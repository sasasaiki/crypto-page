package main

import (
	"net/http"

	"bytes"
	"log"

	"github.com/gorilla/rpc/v2/json2"
	"github.com/sasasaiki/crypto-page/src/go/router"
)

func main() {
	r := router.CreateRoute(router.NewProdRoutingHandlers())

	resp, err := http.Post("http://localhost:8545/rpc", "application/json",
		bytes.NewBufferString(`{
			"jsonrpc": "2.0",
			"method": "eth_getBalance",
			"params": ["0x2f7f14f554632786d74320161a6152adf3b6e94d","latest"],
			"id":1
		}`))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	var ret interface{}
	err = json2.DecodeClientResponse(resp.Body, &ret)
	if err != nil {
		log.Println("aaa")
		log.Fatalln(err)
	}

	log.Println("call")
	if err != nil {
		log.Fatalf("call %v", err)
	}
	log.Printf("result %#v", ret)

	http.ListenAndServe(":8080", r)
}
