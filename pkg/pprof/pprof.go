package pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func Pprof() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	log.Println("starting server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	log.Println(http.ListenAndServe(":80", nil))
	log.Println("server stopped")
	select {}
}
