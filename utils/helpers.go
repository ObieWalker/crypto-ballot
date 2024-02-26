package utils

import (
"fmt"
"crypto/sha256"
"encoding/json"
"strconv"
"strings"
"sync"
)

var mutex = &sync.Mutex{}
var Difficulty = 5


func UnmarshalVaryingStructs (pu []byte, newBlock *Block, fullData *Blockchains, isNewBlock *bool){

	if err := json.Unmarshal(pu, &newBlock); err == nil {
		chainLength := len(fullData.Chains[newBlock.PollingUnit])
		go Mine(newBlock, Difficulty)

		// this avoids a race condition on the chain
		if chainLength == len(fullData.Chains[newBlock.PollingUnit]) {
			*isNewBlock = true
			mutex.Lock()
			fullData.Chains[newBlock.PollingUnit] = append(fullData.Chains[newBlock.PollingUnit], *newBlock)
			mutex.Unlock()
		}
		return
	} else {
		var chainKey string
		var newChain map[string][]Block
		if err := json.Unmarshal(pu, &newChain); err == nil {
			for key := range newChain {
				chainKey = key
			}
			if len(newChain[chainKey]) > len(fullData.Chains[chainKey]) {
				fullData.Chains[chainKey] = newChain[chainKey]
			}
			return
		} else {
			fmt.Println("Failed to unmarshal as either a struct or an array of structs")
			panic(err)
		}
	}
}

func CalculateHash(b Block) string {
	voterData := fmt.Sprint(b.VoterId) + b.ElectionType + b.PollingUnit + b.Selection
	blockData := b.PreviousHash + voterData + b.Timestamp + strconv.Itoa(b.Pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func Mine(b *Block, difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Pow++
		b.Hash = CalculateHash(*b)
	}
}