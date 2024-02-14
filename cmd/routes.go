package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/obiewalker/block-vote/handlers"
)


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