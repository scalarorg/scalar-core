package utils

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	tmos "github.com/tendermint/tendermint/libs/os"
)

// OverwriteFlagDefaults overwrites the default values for already defined flags for the given command and all its children.
// Changes the currently set value if the flag is set.
func OverwriteFlagDefaults(c *cobra.Command, defaults map[string]string, updateVal bool) {
	set := func(s *pflag.FlagSet, key, val string) {
		if f := s.Lookup(key); f != nil {
			f.DefValue = val
			if updateVal || !f.Changed {
				_ = f.Value.Set(val)
			}

			if updateVal {
				f.Changed = true
			}
		}
	}
	for key, val := range defaults {
		set(c.Flags(), key, val)
		set(c.PersistentFlags(), key, val)
	}
	for _, c := range c.Commands() {
		OverwriteFlagDefaults(c, defaults, updateVal)
	}
}

func WriteFile(name string, dir string, contents []byte) error {
	file := filepath.Join(dir, name)

	err := tmos.EnsureDir(dir, 0o755)
	if err != nil {
		return err
	}

	return tmos.WriteFile(file, contents, 0o644)
}
