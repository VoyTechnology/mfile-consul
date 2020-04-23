// Package mfileconsul registers the consul provider for
// github.com/voytechnology/mfile
package mfileconsul // import "github.com/voytechnology/mfile-consul"

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/voytechnology/mfile"
)

func init() {
	// We use the non-pooled config to avoid using too many connections.
	// TODO: Allow the use of configurable client
	c, err := api.NewClient(api.DefaultNonPooledConfig())
	if err != nil {
		panic(fmt.Sprintf("mfile(consul): unable to connect to consul: %v", err))
	}

	if err := mfile.Register("consul", handler{c.KV()}); err != nil {
		panic(fmt.Sprintf("mfile(consul): unable to register: %v", err))
	}
}

type handler struct {
	kv *api.KV
}

// ReadFile gets a single KV from consul.
func (h handler) ReadFile(path string) ([]byte, error) {
	res, _, err := h.kv.Get(path, nil)
	if err != nil {
		return nil, fmt.Errorf("mfile(consul): unable to get %s: %w", path, err)
	}
	return res.Value, nil
}
