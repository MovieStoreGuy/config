package config

import "io"

// Config defines a basic interface that allows for wrapping
// behaviour for a nice config object
type Config interface {
	// Err returns any internal errors that happened while processing input
	Err() error

	// From allows for reading from a writer and applying to the config
	From(data io.Reader) Config

	// ParseEnv reads from the environment and will apply the found values
	// to the fields configured.
	ParseEnv() Config

	// ParseFlags will take an array of input and parse them
	// applying to the internal config.
	ParseFlags(args []string) Config
}
