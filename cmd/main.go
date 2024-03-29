package main

import (
	"log"
	"net/http"
	"time"

	"github.com/obiewalker/block-vote/p2p"
)

func main() {
	go p2p.Main(true)
	log.Fatal(run())
}

func run() error {
	mux := makeMuxRouter()
	log.Println("Listening on ", 8081)
	s := &http.Server{
					Addr:           ":8081",
					Handler:        mux,
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second,
					MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
