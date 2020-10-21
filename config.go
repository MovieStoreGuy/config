package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/jessevdk/go-flags"
	"github.com/kelseyhightower/envconfig"
)

type wrapper struct {
	err error

	AppName    string
	NewDecoder func(r io.Reader) Decoder

	Values interface{}
}

var _ Config = (*wrapper)(nil)

// Default wraps the values and provides support for the configuration builder
// It will set the default decoder as json
func Default(values interface{}) Config {
	return &wrapper{
		NewDecoder: decoders["json"],
		AppName:    path.Base(os.Args[0]),
		Values:     values,
	}
}

// WithDecoding allows for a custom decoder to be used when running `From`
// if there is no valid decode, an error is returned.
//
// Valid decodings are yaml, json, or toml.
func WithDecoding(value interface{}, decode string) (Config, error) {
	dec, exist := decoders[decode]
	if !exist {
		return nil, fmt.Errorf(`invalid decoding %s provided`, decode)
	}
	return &wrapper{
		AppName:    path.Base(os.Args[0]),
		NewDecoder: dec,
		Values:     value,
	}, nil
}

func (w *wrapper) Err() error {
	return w.err
}

func (w *wrapper) From(data io.Reader) Config {
	if w.err != nil {
		return w
	}
	var enc Decoder = json.NewDecoder(data)
	if w.NewDecoder != nil {
		enc = w.NewDecoder(data)
	}
	w.err = enc.Decode(w.Values)
	return w
}

func (w *wrapper) ParseEnv() Config {
	if w.err != nil {
		return w
	}
	w.err = envconfig.Process(w.AppName, w.Values)
	return w
}

func (w *wrapper) ParseFlags(args []string) Config {
	if w.err != nil {
		return w
	}
	_, w.err = flags.NewParser(w.Values, flags.Default).ParseArgs(args)
	return w
}
