package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"time"
)

const defaultTick = Duration(500 * time.Millisecond)

type Configuration struct {
	DSN            string                      `json:"dsn"`
	Authentication AuthenticationConfiguration `json:"authentication"`
	Tick           Duration                    `json:"tick,omitempty"`
	Command        string                      `json:"command"`

	authenticationStrategy *HttpStrategy
	Transport              *Transport
}

func (config *Configuration) LoadConfig(args []string) error {
	if len(args) <= 1 {
		return fmt.Errorf("no config file specified (example: ./notifier-daemon config.json)")
	}

	if len(args) > 2 {
		return fmt.Errorf("too many arguments (example: ./notifier-daemon config.json)")
	}

	configFile := args[1]

	jsonFile, err := os.Open(configFile)

	if err != nil {
		return err
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()

		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &config)

	if err != nil {
		return err
	}

	if config.Tick == 0 {
		config.Tick = defaultTick
	}

	if config.Tick < Duration(time.Millisecond*20) {
		return fmt.Errorf("tick must be greater than 20 milliseconds")
	}

	dsn, err := url.Parse(config.DSN)

	if err != nil {
		return err
	}

	strategy, err := NewStrategy(&config.Authentication)

	if err != nil {
		return err
	}

	config.authenticationStrategy = &strategy

	transport, err := NewTransport(dsn, config.authenticationStrategy)

	if err != nil {
		return err
	}

	config.Transport = &transport

	return nil
}
