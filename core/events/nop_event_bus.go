package events

import (
	"context"

	tmPubSub "github.com/tendermint/tendermint/libs/pubsub"
)

type nopEventBus struct{}

func NewNopeEventBus() EventBus {
	return &nopEventBus{}
}

func (nopEventBus) Start() error {
	return nil
}
func (nopEventBus) Stop() error {
	return nil
}

func (nopEventBus) Subscribe(ctx context.Context, subscriber string, query tmPubSub.Query) (*tmPubSub.Subscription, error) {
	return nil, nil
}

func (nopEventBus) Unsubscribe(ctx context.Context, subscriber string, query tmPubSub.Query) error {
	return nil
}

func (nopEventBus) UnsubscribeAll(ctx context.Context, subscriber string) error {
	return nil
}

func (nopEventBus) Publish(msg interface{}, tags map[string]string) error {
	return nil
}
