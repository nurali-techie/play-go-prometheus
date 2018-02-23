package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	log.Println("Staring http server ..")

	http.Handle("/metrics", prometheus.Handler())
	http.Handle("/myapp", prometheus.InstrumentHandler("myapp", http.HandlerFunc(myAppHander)))

	log.Println("Server is up at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

var helloCnt = 0

func myAppHander(res http.ResponseWriter, req *http.Request) {
	helloCnt = helloCnt + 1
	res.Write([]byte("Hello World:" + strconv.Itoa(helloCnt)))
}
