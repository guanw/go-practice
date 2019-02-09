package main

import (
	"go.uber.org/dig"
	"encoding/json"
	"log"
	"os"
	)

type Config struct {
	Prefix string
}

type Middleware struct {
	name string
}

func main() {
	c := dig.New()

	// Provide a Config object. This can fail to decode.
	err := c.Provide(func() (*Config, error) {
		// In a real program, the configuration will probably be read from a
		// file.
		var cfg Config
		err := json.Unmarshal([]byte(`{"prefix": "[foo] "}`), &cfg)
		return &cfg, err
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func(cfg *Config) (Middleware, error) {
		mw := Middleware{
			name: cfg.Prefix,
		}
		return mw, nil
	})
	if err != nil {
		panic(err)
	}

	// Provide a way to build the logger based on the configuration.
	err = c.Provide(func(mw Middleware) *log.Logger {
		return log.New(os.Stdout, mw.name, 0)
	})
	if err != nil {
		panic(err)
	}

	// Invoke a function that requires the logger, which in turn builds the
	// Config first.
	err = c.Invoke(func(l *log.Logger) {
		l.Print("You've been invoked")
	})
	if err != nil {
		panic(err)
	}
}

