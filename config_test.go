package config_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MovieStoreGuy/config"
)

func TestDecodingConfig(t *testing.T) {
	t.Parallel()

	conf := &struct {
		Nevermore string `json:"nevermore"`
	}{}

	w := config.Default(conf).From(strings.NewReader(`{"nevermore":"more"}`))

	assert.NoError(t, w.Err(), `Must have been able to decode json string`)
	assert.Equal(t, `more`, conf.Nevermore)
}

func TestParsingFlags(t *testing.T) {
	t.Parallel()

	v := &struct {
		Index int `short:"i"`
	}{}

	conf := config.Default(v).ParseFlags([]string{
		"-i", "10",
	})

	assert.NoError(t, conf.Err(), `Must not have thrown an error`)
	assert.Equal(t, 10, v.Index, `Must be equal`)
}
