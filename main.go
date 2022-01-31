package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)

	config := &Configuration{}

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGHUP:
					err := config.LoadConfig(os.Args)

					if err != nil {
						log.Printf("Error loading config: %s", err)
						cancel()
						os.Exit(1)
					}
				case os.Interrupt:
					cancel()
					os.Exit(1)
				}
			case <-ctx.Done():
				log.Printf("Done.")
				os.Exit(1)
			}
		}
	}()

	if err := Run(ctx, config, os.Stdout); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%s\n", err)

		if err != nil {
			log.Fatal(err)
		}

		os.Exit(1)
	}
}
