// Package types contains all the types that are used in the application.
package types

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Plugin is a struct that contains the producer, consumer, trigger handler and context. Forming a single use-case like "If X happens then do Y".
type Plugin struct {
	Producer       Producer
	Consumer       Consumer
	TriggerHandler func(ctx context.Context, ch chan<- Message) error
	Context        context.Context
}

// Produce is a method that starts the producer within the plugin.
func (p *Plugin) Produce() error {
	trigger := NewTrigger(p.TriggerHandler, p.Producer)
	return trigger.Start(p.Context)
}

// Consume is a method that starts the consumer within the plugin.
func (p *Plugin) Consume() error {
	return p.Consumer.Init()
}

// ErrGoPluginsProduce is a function that starts all the producers in a separate goroutine.
func ErrGoPluginsProduce(plugins []Plugin) error {
	eg := errgroup.Group{}
	for _, plugin := range plugins {
		plugin := plugin
		eg.Go(func() error {
			return plugin.Produce()
		})
	}

	return eg.Wait()
}

// ErrGoPluginsConsume is a function that starts all the consumers in a separate goroutine.
func ErrGoPluginsConsume(plugins []Plugin) error {
	eg := errgroup.Group{}
	for _, plugin := range plugins {
		plugin := plugin
		eg.Go(func() error {
			return plugin.Consume()
		})
	}

	return eg.Wait()
}
