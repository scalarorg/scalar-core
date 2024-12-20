package config

import (
	"reflect"

	"github.com/mitchellh/mapstructure"

	"github.com/scalarorg/scalar-core/vald/evm/rpc"
)

func stringToEnumType(
	f reflect.Type,
	t reflect.Type,
	data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String {
		return data, nil
	}

	if t == reflect.TypeOf(rpc.FinalityOverride(0)) {
		return rpc.ParseFinalityOverride(data.(string))
	}

	return data, nil
}

// AddDecodeHooks adds decode hooks to the given config to correctly translate string into FinalityOverride
func AddDecodeHooks(cfg *mapstructure.DecoderConfig) {
	hooks := []mapstructure.DecodeHookFunc{
		stringToEnumType,
	}
	if cfg.DecodeHook != nil {
		hooks = append(hooks, cfg.DecodeHook)
	}

	cfg.DecodeHook = mapstructure.ComposeDecodeHookFunc(hooks...)
}
