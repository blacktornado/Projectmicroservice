package main

import (
	gateway "apigateway"
	m "middleware"
	"net/http"
)

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/top-tracks-with-lyrics", gateway.GetTopTrackL)
	w := m.NewLoggerMiddleware(mux)
	http.ListenAndServe(":8082", w)
}
