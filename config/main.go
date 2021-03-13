package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type (
	// Type exposes methods to read and write configurations.
	Type struct {
		Path string
		v    *viper.Viper
		raw  map[string]interface{}
	}
)

//
// Get returns a key value,
// * contextualized for a node (by default the local node, customized by the
//   impersonate option)
// * dereferenced
// * evaluated
//
func (t *Type) Get(key string) interface{} {
	val := t.v.Get(key)
	log.Debug().Msgf("config %s get %s => %s", t.Path, key, val)
	return val
}
