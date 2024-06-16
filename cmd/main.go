// Package main is the entry point for the application.
package main

import (
	"context"
	"log"
	"os"

	"github.com/atropos112/atrka/consumers"
	"github.com/atropos112/atrka/producers"
	"github.com/atropos112/atrka/triggers"
	"github.com/atropos112/atrka/types"
	prettyconsole "github.com/thessem/zap-prettyconsole"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger := prettyconsole.NewLogger(zap.DebugLevel)
	ctx := context.Background()
	plugins := []types.Plugin{
		{
			Producer:       producers.LoggingProducer{Logger: logger},
			Consumer:       consumers.NoOpConsumer{},
			TriggerHandler: triggers.WebhookHandler(8080, "/webhook"),
			Context:        ctx,
		},
		{
			Producer:       producers.LoggingProducer{Logger: logger},
			Consumer:       consumers.NoOpConsumer{},
			TriggerHandler: triggers.CronHandler("*/1 * * * * *"),
			Context:        ctx,
		},
		{
			Producer:       producers.InMemoryProducer{},
			Consumer:       consumers.InMemoryConsumer{},
			TriggerHandler: triggers.CronHandler("*/1 * * * * *"),
			Context:        ctx,
		},
	}

	eg := errgroup.Group{}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "dev",
			},
			&cli.StringFlag{
				Name:    "service",
				Aliases: []string{"s"},
				Usage:   "service type to run (producer, consumer or combined)",
			},
			&cli.StringFlag{
				Name:    "broker(s)",
				Aliases: []string{"t"},
				Value:   "rp0:9094, rp1:9094, rp2:9094",
				Usage:   "broker(s) to use, separated by commas",
			},
		},
		Action: func(cCtx *cli.Context) error {
			var logger *zap.Logger

			// Starting logger (different depending on dev flag)
			if cCtx.Bool("dev") {
				// logger, _ = zap.NewDevelopment()
				logger = prettyconsole.NewLogger(zap.DebugLevel)
			} else {
				logger, _ = zap.NewProduction()
			}
			defer logger.Sync() // nolint: errcheck
			l := logger.Sugar()

			// Topic(s) to use
			// brokers := strings.Split(cCtx.String("topic(s)"), ",")

			if cCtx.String("service") == "producer" {
				eg.Go(func() error {
					return types.ErrGoPluginsProduce(plugins)
				})
			} else if cCtx.String("service") == "consumer" {
				eg.Go(func() error {
					return types.ErrGoPluginsConsume(plugins)
				})
			} else if cCtx.String("service") == "combined" {

				eg.Go(func() error {
					return types.ErrGoPluginsProduce(plugins)
				})
				eg.Go(func() error {
					return types.ErrGoPluginsConsume(plugins)
				})
			} else {
				l.Panic("Invalid service type.")
			}

			return eg.Wait()
		},
	}

	if err := eg.Wait(); err != nil {
		panic(err)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
