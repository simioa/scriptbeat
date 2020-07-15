// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type ScriptConfig struct {
	Script string        `config:"script"`
	Args   []string      `config:"args"`
	Env    []string      `config:"env"`
	Period time.Duration `config:"period"`
	Alias  string        `config:"alias"`
}

type Config struct {
	Scripts []ScriptConfig `config:"scripts"`
}
