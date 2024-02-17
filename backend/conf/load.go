package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
)

// Configuration the load Object method

// LoadFromFile
// github.com/BurntSushi/toml go toml parse library
// https://github.com/BurntSushi/toml library
// object <---> toml configurations file
func LoadFromFile(filepath string) error {
	c := DefaultConfig()
	if _, err := toml.DecodeFile(filepath, c); err != nil {
		return err
	}
	config = c
	return nil
}

// LoadFromEnv
// "github.com/caarlos0/env/v6" read ENV variable
// env ---> object
func LoadFromEnv() error {
	c := DefaultConfig()
	// env Tag
	if err := env.Parse(c); err != nil {
		return err
	}
	config = c
	// c.MySQL.Host = os.Getenv("DATASOURCE_HOST")
	return nil
}
