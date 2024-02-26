package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/obiewalker/block-vote/utils"
)

var TemporaryBroadcastBlock []utils.Block

var fullData = utils.Blockchains{
	Chains: make(map[string][]utils.Block),
}

func CreatePUBlockchain(pollingUnit string) []utils.Block {
	genesisBlock := utils.Block{}
	genesisBlock = utils.Block{
		PollingUnit: pollingUnit,
		Hash: utils.CalculateHash(genesisBlock),		
		Timestamp: time.Now().String(),
		Difficulty: utils.Difficulty,
}
	newChain := []utils.Block{genesisBlock}
	fullData.Chains[pollingUnit] = newChain
	TemporaryBroadcastBlock = append(TemporaryBroadcastBlock, genesisBlock)

	return fullData.Chains[pollingUnit]
}

func AddBlock(pollingUnit string, voterId string, electionType string, selection string) []utils.Block{
	lastBlock := fullData.Chains[pollingUnit][len(fullData.Chains[pollingUnit])-1]
	b := sha256.Sum256([]byte(voterId))
	hashedVoterId := hex.EncodeToString(b[:])
	newBlock := utils.Block{
		VoterId:			hashedVoterId,
		ElectionType: electionType,
		PollingUnit:  pollingUnit,
		Selection: 		selection,
		PreviousHash: lastBlock.Hash,
		Timestamp:    time.Now().String(),
	}

	TemporaryBroadcastBlock = append(TemporaryBroadcastBlock, newBlock)
	return fullData.Chains[pollingUnit]
}

func IsBlockChainValid(pollingUnit string) bool {
	blockchain := fullData.Chains[pollingUnit]
	for i := range blockchain[1:] {
		previousBlock := blockchain[i]
		currentBlock := blockchain[i+1]
		if currentBlock.Hash != utils.CalculateHash(currentBlock) || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func GetPUBlockChain(pollingUnit string) []utils.Block {
	blockchain := fullData.Chains[pollingUnit]
	return blockchain
}

func GetAllBlockChain() utils.Blockchains{
	return fullData
}

// This ensures the same voter does not vote twice at the same polling unit
func CheckVoterIntegrity(voterID string, pollingUnit string) (bool, error) {
	if _, ok := fullData.Chains[pollingUnit]; !ok {
		return false, &utils.PollingUnitNotExistError{}
	}

	b := sha256.Sum256([]byte(voterID))
	hashedVoterId := hex.EncodeToString(b[:])
	
	for _, val := range fullData.Chains[pollingUnit][1:] {
		if val.VoterId == hashedVoterId {
			return false, &utils.VoterAlreadyVotedError{}
		}
	}
	return true, nil
}

func CalculatePUVotes(pollingUnit string) (map[string]int, string){
	if _, ok := fullData.Chains[pollingUnit]; !ok {
		errMessage :=fmt.Sprintf("Polling Unit %v doesn't exists.", pollingUnit)
		return nil, errMessage
	}
	frequency := make(map[string]int)
	for _, item := range fullData.Chains[pollingUnit][1:] {
		frequency[item.Selection]++
	}
	return frequency, ""
}

func CalculateVotes() map[string]int {
	frequency := make(map[string]int)
	for _, val := range fullData.Chains {
		for _, item := range val[1:] {
			frequency[item.Selection]++
		}
	}
	return frequency
}

func CalculateVotesByPU() map[string]map[string]int {
	puVotes := make(map[string]map[string]int )
	for _, val := range fullData.Chains {
		votes, err := CalculatePUVotes(val[0].PollingUnit)

		if err != "" {
			fmt.Println("Error getting votes")
		}
		puVotes[val[0].PollingUnit] = votes
	}
	return puVotes
}