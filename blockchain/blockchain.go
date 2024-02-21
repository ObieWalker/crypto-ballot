package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
	"sync"

	"github.com/obiewalker/block-vote/utils"
)

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

var mutex = &sync.Mutex{}

var TemporaryBroadcastBlock []Block

type Blockchains struct {
	Chains map[string][]Block
}

var fullData = Blockchains{
	Chains: make(map[string][]Block),
}

var _difficulty = 5

func CreatePUBlockchain(pollingUnit string) []Block {
	genesisBlock := Block{}
	genesisBlock = Block{
		PollingUnit: pollingUnit,
		Hash: genesisBlock.calculateHash(),		
		Timestamp: time.Now().String(),
		Difficulty: _difficulty,
}
	newChain := []Block{genesisBlock}
	fullData.Chains[pollingUnit] = newChain
	TemporaryBroadcastBlock = append(TemporaryBroadcastBlock, genesisBlock)

	return fullData.Chains[pollingUnit]
}

func AddBlock(pollingUnit string, voterId string, electionType string, selection string) []Block{
	lastBlock := fullData.Chains[pollingUnit][len(fullData.Chains[pollingUnit])-1]
	b := sha256.Sum256([]byte(voterId))
	hashedVoterId := hex.EncodeToString(b[:])
	newBlock := Block{
		VoterId:			hashedVoterId,
		ElectionType: electionType,
		PollingUnit:  pollingUnit,
		Selection: 		selection,
		PreviousHash: lastBlock.Hash,
		Timestamp:    time.Now().String(),
	}

	// In the real world application, the next lines of code would be
	// handled by the connected nodes after the block is broadcast
	newBlock.mine(_difficulty)
	// I am doing this to avoid race conditions
	mutex.Lock()
	fullData.Chains[pollingUnit] = append(fullData.Chains[pollingUnit], newBlock)
	TemporaryBroadcastBlock = append(TemporaryBroadcastBlock, newBlock)
	mutex.Unlock()
	return fullData.Chains[pollingUnit]
}

func (b Block) calculateHash() string {
	voterData := fmt.Sprint(b.VoterId) + b.ElectionType + b.PollingUnit + b.Selection
	blockData := b.PreviousHash + voterData + b.Timestamp + strconv.Itoa(b.Pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Pow++
		b.Hash = b.calculateHash()
	}
}

func IsBlockChainValid(pollingUnit string) bool {
	blockchain := fullData.Chains[pollingUnit]
	for i := range blockchain[1:] {
		previousBlock := blockchain[i]
		currentBlock := blockchain[i+1]
		if currentBlock.Hash != currentBlock.calculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func GetPUBlockChain(pollingUnit string) []Block {
	blockchain := fullData.Chains[pollingUnit]
	return blockchain
}

func GetAllBlockChain() Blockchains{
	return fullData
}

// This ensures the same voter does not vote twice at the same polling unit
func CheckVoterIntegrity(voterID string, pollingUnit string) (bool, error) {
	if _, ok := fullData.Chains[pollingUnit]; !ok {
		return false, &errors.PollingUnitNotExistError{}
	}

	b := sha256.Sum256([]byte(voterID))
	hashedVoterId := hex.EncodeToString(b[:])
	
	for _, val := range fullData.Chains[pollingUnit][1:] {
		if val.VoterId == hashedVoterId {
			return false, &errors.VoterAlreadyVotedError{}
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