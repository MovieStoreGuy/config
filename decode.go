package config

import (
	"encoding/json"
	"io"

	"github.com/influxdata/toml"
	"gopkg.in/yaml.v3"
)

// Decoder allows for dynamic decodings for configuration.
type Decoder interface {
	Decode(value interface{}) error
}

var decoders = map[string]func(r io.Reader) Decoder{
	"json": func(r io.Reader) Decoder {
		return json.NewDecoder(r)
	},
	"yaml": func(r io.Reader) Decoder {
		return yaml.NewDecoder(r)
	},
	"toml": func(r io.Reader) Decoder {
		return toml.NewDecoder(r)
	},
}
