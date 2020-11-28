package config_test

import (
	"testing"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	config.New()
}

func TestErrorGetMySQLCli(t *testing.T) {
	assert.Panics(t, func() {
		config.GetMySQLCli("-")
	})
}
