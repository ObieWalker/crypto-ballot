package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/obiewalker/block-vote/handlers"
)


func main() {
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

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handlers.HandleTotalVotes).Methods("GET")
	muxRouter.HandleFunc("/", handlers.HandleWriteBlock).Methods("POST")
	muxRouter.HandleFunc("/{pu}", handlers.HandleGetPUBlockchain).Methods("GET")
	muxRouter.HandleFunc("/pollingunit", handlers.HandleCreatePU).Methods("POST")
	muxRouter.HandleFunc("/validate/{pu}", handlers.HandleValidateChain).Methods("POST")
	muxRouter.HandleFunc("/puvotes/{pu}", handlers.HandleCountPUVotes).Methods("GET")
	muxRouter.HandleFunc("/", handlers.HandleTotalVotes).Methods("GET")
	return muxRouter
}