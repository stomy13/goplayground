package main

import (
	"expvar"
	"net"
	"net/http"
)

func main() {
	// curl http://localhost:8081/ -> "name": "stomy"
	Name.Set("stomy")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	http.Serve(listener, expvar.Handler())
}

var Name *expvar.String = expvar.NewString("name")
