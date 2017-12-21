package main

import (
	"net/http"

	"github.com/sasasaiki/eth-page/src/go/router"
)

func main() {
	r := router.CreateRoute(router.NewProdRoutingHandlers())
	http.ListenAndServe(":8080", r)
	
}
