package events

import (
	"crypto/rand"
	"fmt"

	tmQuery "github.com/tendermint/tendermint/libs/pubsub/query"
	hex "github.com/tmthrgd/go-hex"
)

func QueryForTx(txHash []byte) *tmQuery.Query {
	return tmQuery.MustParse(fmt.Sprintf("gallactic.events.tx.hash='%X'", txHash))
}

func TagsForTx(txHash []byte) map[string]string {
	return map[string]string{"gallactic.events.tx.hash": fmt.Sprintf("%X", txHash)}
}

func GenSubID() string {
	bs := make([]byte, 32)
	rand.Read(bs)
	return hex.EncodeUpperToString(bs)
}
