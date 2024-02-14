package blockchain

import (
	"crypto/sha256"
	// "encoding/json"
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

	return fullData.Chains[pollingUnit]
}

func UpdatePUBlockchain(data Blockchains, pollingUnit string, block Block) []Block {
	fullData.Chains[pollingUnit] = append(fullData.Chains[pollingUnit], block)
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
	newBlock.mine(_difficulty)
	mutex.Lock()
	fullData.Chains[pollingUnit] = append(fullData.Chains[pollingUnit], newBlock)
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