package conf

import (
	"encoding/json"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Instead of directly exposing variables, a better way is to use functions
var config *Config

func C() *Config {
	// sync.Lock
	if config == nil {
		// define default value
		config = DefaultConfig()
	}
	return config
}

func DefaultConfig() *Config {
	return &Config{
		Application: &Application{
			Domain: "127.0.0.1",
		},
		MySQL: &MySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "yyblog",
			Username: "root",
			Password: "root",
			Debug:    true,
		},
	}
}

// Program configuration, the configuration will be read at startup and provide the required global variables for the program
// toml
// Make the configuration object a global variable (single mode)
/*
[mysql]
host="127.0.0.1"
port=3306
...
*/

type Config struct {
	Application *Application `json:"app" yaml:"app" toml:"app"`
	MySQL       *MySQL       `json:"mysql" yaml:"mysql" toml:"mysql"`
}

type Application struct {
	Domain string `json:"domain" yaml:"domain" toml:"domain" env:"APP_DOMAIN"`
}

// if Customize Object fmt.PrintXXX()
// String() string
// &{0x1400007c600} ---> JSON {}
func (c *Config) String() string {
	jd, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("%p", c)
	}
	return string(jd)
}

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

// DSN dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m *MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}

// GetDB via configurations to get a DB Object
func (m *MySQL) GetDB() *gorm.DB {
	m.l.Lock()
	defer m.l.Unlock()

	if m.db == nil {
		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.db = db

		// Add Debug Config
		if m.Debug {
			m.db = db.Debug()
		}
	}

	return m.db
}

// DB Configuration object to provide the global single mode configuration
func (c *Config) DB() *gorm.DB {
	return c.MySQL.GetDB()
}
