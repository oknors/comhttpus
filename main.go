package main

import (
	"fmt"
	"github.com/oknors/comhttpus/hnd"
	"log"
	"net/http"
	"time"
)

func main() {
	port := "4477"
	srv := &http.Server{
		Handler:      hnd.Handlers(),
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on port: ", port)
	log.Fatal(srv.ListenAndServe())
}
