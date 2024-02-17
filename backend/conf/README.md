### Program Configurations Object

#### Define Configurations
```
// MySQL db object (single mode)
type MySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`

	// Judge this private property to determine whether to return an existing object
	db *gorm.DB
	l  sync.Mutex
}
```

### load configurations
#### File: TOML

- JSON/YAML/TOML: have a format of the parsing library
```
# {"mysql": {"host":...}}
[mysql]
  host = "127.0.0.1"
  port = 3306
  database = ""
  username = ""
  password = ""
  debug = false
```

```go
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
```