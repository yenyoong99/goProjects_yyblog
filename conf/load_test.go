package conf_test

import (
	"testing"

	"goProjects/yyblog/conf"
)

func TestLoadFromFile(t *testing.T) {
	err := conf.LoadFromFile("etc/application.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

func TestLoadFromEnv(t *testing.T) {
	// os.Setenv("DATASOURCE_HOST", "127.0.0.1")
	err := conf.LoadFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

