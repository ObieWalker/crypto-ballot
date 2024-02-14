package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"github.com/obiewalker/block-vote/blockchain"
)

var mutex = &sync.Mutex{}

type VoteData struct {
	VoterId string
	ElectionType string
	Selection string
	PollingUnit string
}

type PU struct {
	PollingUnit string
}

func HandleGetPUBlockchain(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
  pollingUnit := params["pu"]
	chains := blockchain.GetPUBlockChain(pollingUnit)
	bytes, err := json.MarshalIndent(chains, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("HTTP 500: Internal Server Error"))
					return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d VoteData

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&d); err != nil {
					respondWithJSON(w, r, http.StatusBadRequest, r.Body)
					return
	}   
	defer r.Body.Close()
	_, err := blockchain.CheckVoterIntegrity(d.VoterId, d.PollingUnit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("HTTP 400: Voter already voted."))
		return
	}

	newBlock := blockchain.AddBlock(d.PollingUnit, d.VoterId, d.ElectionType, d.Selection)
	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func HandleCreatePU(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p PU

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
					respondWithJSON(w, r, http.StatusBadRequest, r.Body)
					return
	}   
	defer r.Body.Close()
	newBlock := blockchain.CreatePUBlockchain(p.PollingUnit)

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func HandleValidateChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
  pollingUnit := params["pu"]

	defer r.Body.Close()
	isValid := blockchain.IsBlockChainValid(pollingUnit)

	respondWithJSON(w, r, http.StatusOK, isValid)
}

func HandleCountPUVotes(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
  pollingUnit := params["pu"]
	frequency, err := blockchain.CalculatePUVotes(pollingUnit)
	if err != "" {
		respondWithJSON(w, r, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, r, http.StatusOK, frequency)
	return
}

func HandleTotalVotes(w http.ResponseWriter, r *http.Request){
	frequency := blockchain.CalculateVotes()

	respondWithJSON(w, r, http.StatusOK, frequency)
	return
}