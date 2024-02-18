package p2p

import (
	"context"
	"flag"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	"github.com/obiewalker/block-vote/pubsubpkg"
)

var (
	blockFlag = flag.String("blocks", "block", "Write and Read blocks")
)

func Main(host bool) {
	flag.Parse()
	ctx := context.Background()

	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}
	go pubsubpkg.DiscoverPeers(ctx, h, blockFlag)

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}
	topic, err := ps.Join(*blockFlag)
	if err != nil {
		panic(err)
	}
	go pubsubpkg.PublishData(ctx, topic)

	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}
	pubsubpkg.ReadIncoming(ctx, sub, topic, host)
}
