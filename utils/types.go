package utils

type Block struct {
	VoterId					string
	ElectionType 		string
	PollingUnit			string
	Selection				string
	Hash        		string
	PreviousHash		string
	Timestamp    		string
	Pow          		int
	Difficulty	    int
}

type Blockchains struct {
	Chains map[string][]Block
}

type VoteData struct {
	VoterId string
	ElectionType string
	Selection string
	PollingUnit string
}

type PU struct {
	PollingUnit string
}