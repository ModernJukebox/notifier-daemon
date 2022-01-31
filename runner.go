package main

import (
	"context"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func Run(ctx context.Context, config *Configuration, stdout io.Writer) error {
	log.SetOutput(os.Stdout)

	err := config.LoadConfig(os.Args)

	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(time.Duration(config.Tick)):

			data, err := ExecuteCommand(config)

			if err != nil {
				return err
			}

			data = strings.TrimSpace(data)

			log.Printf("data: %s", data)

			err = (*config.Transport).Send(data)

			if err != nil {
				return err
			}

		}
	}
}
