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
	log.SetOutput(stdout)

	err := config.LoadConfig(os.Args)

	if err != nil {
		return err
	}

	if err := execute(config); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(time.Duration(config.Tick)):
			if err := execute(config); err != nil {
				return err
			}
		}
	}
}

func execute(config *Configuration) error {
	data, err := ExecuteCommand(config)

	if err != nil {
		return err
	}

	data = strings.TrimSpace(data)

	err = (*config.Transport).Send(data)

	if err != nil {
		return err
	}

	return nil
}
