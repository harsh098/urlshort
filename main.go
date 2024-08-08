package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/harsh098/urlshort/internal"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Healthy")
	})
	return mux
}

func main() {
	var handler http.Handler
	var mux *http.ServeMux = defaultMux()
	socketAddress, err := internal.GetSocketAddress()
	if err != nil {
		log.Fatalf("Cannot Determine Socket Address: \n\t%v", err.Error())
		os.Exit(-1)
	}
	handler, err = internal.YAMLHandler(mux)
	if err != nil {
		log.Fatalf("Cannot Start Server \n\t%v", err.Error())
		os.Exit(-1)
	}
	
	if err:=http.ListenAndServe(socketAddress, handler); err != nil {
		log.Fatalf("Server Failed to start due to following errors:\n%v", err.Error())
	}
}