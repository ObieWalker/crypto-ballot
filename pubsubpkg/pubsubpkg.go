package pubsubpkg

import (
	// "bufio"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"

	"github.com/obiewalker/block-vote/blockchain"
)

var mutex = &sync.Mutex{}

func initDHT(ctx context.Context, h host.Host) *dht.IpfsDHT {
	kademliaDHT, err := dht.New(ctx, h)
	if err != nil {
		panic(err)
	}
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, peerAddr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Connect(ctx, *peerinfo); err != nil {
				fmt.Println("Bootstrap warning:", err)
			}
		}()
	}
	wg.Wait()

	return kademliaDHT
}

func DiscoverPeers(ctx context.Context, h host.Host, topicNameFlag *string) {
	kademliaDHT := initDHT(ctx, h)
	routingDiscovery := drouting.NewRoutingDiscovery(kademliaDHT)
	dutil.Advertise(ctx, routingDiscovery, *topicNameFlag)

	// Look for others who have announced and attempt to connect to them
	anyConnected := false
	for !anyConnected {
		fmt.Println("Searching for peers...")
		peerChan, err := routingDiscovery.FindPeers(ctx, *topicNameFlag)
		if err != nil {
			panic(err)
		}
		for peer := range peerChan {
			if peer.ID == h.ID() {
				continue // No self connection
			}
			err := h.Connect(ctx, peer)
			if err != nil {
				fmt.Printf("Failed connecting to %s, error: %s\n", peer.ID, err)
			} else {
				fmt.Println("Connected to:", peer.ID)
				anyConnected = true
			}
		}
	}
	fmt.Println("Peer discovery complete")
}

func PublishData(ctx context.Context, topic *pubsub.Topic) {
	for {
		if len(blockchain.TemporaryBroadcastBlock) > 0 {
			bytes, err := json.Marshal(blockchain.TemporaryBroadcastBlock[0])
			if err != nil {
				fmt.Println("Error Marshalling", err)
			}
			if err := topic.Publish(ctx, []byte(bytes)); err != nil {
				fmt.Println("### Publish error:", err)
			}
			blockchain.TemporaryBroadcastBlock = blockchain.TemporaryBroadcastBlock[1:]
		}
	}
}

func ReadIncoming(ctx context.Context, sub *pubsub.Subscription, topic *pubsub.Topic, host bool) {
	for {
		if !host {
			var newBlock blockchain.Block
			var fullData = blockchain.Blockchains{
				Chains: make(map[string][]blockchain.Block),
			}
			m, err := sub.Next(ctx)
			if err != nil {
				panic(err)
			}
			pu := m.Message.Data
			if err := json.Unmarshal(pu, &newBlock); err != nil {
        panic(err)
    	}
			mutex.Lock()
			fullData.Chains[newBlock.PollingUnit] = append(fullData.Chains[newBlock.PollingUnit], newBlock)
			mutex.Unlock()

			bytes, err := json.Marshal(fullData.Chains[newBlock.PollingUnit])
			if err != nil {
				fmt.Println("Error Marshalling", err)
			}
			if err := topic.Publish(ctx, []byte(bytes)); err != nil {
				fmt.Println("### Publish error:", err)
			}
		}
	}
}
