package main

import (
	"context"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type DataStorage struct {
	data string
}

func Run(ctx context.Context, config *Configuration, stdout io.Writer) error {
	log.SetOutput(stdout)

	err := config.LoadConfig(os.Args)

	if err != nil {
		return err
	}

	storage := &DataStorage{}

	if err := execute(config, storage); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(time.Duration(config.Tick)):
			if err := execute(config, storage); err != nil {
				return err
			}
		}
	}
}

func execute(config *Configuration, storage *DataStorage) error {
	data, err := ExecuteCommand(config)

	if err != nil {
		return err
	}

	data = strings.TrimSpace(data)

	// If the data is the same, don't do anything
	if storage.data == data {
		return nil
	}

	err = (*config.Transport).Send(data)

	if err != nil {
		return err
	}

	storage.data = data

	return nil
}
