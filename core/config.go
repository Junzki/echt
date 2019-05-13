package core

import (
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/pkg/errors"
	"log"
	"strings"
)

type DatabaseArgs struct {
	Type     string `env:"DATABASE_TYPE"`

	Host     string `env:"DATABASE_HOST"`
	Port     string `env:"DATABASE_PORT"`

	Name     string `env:"DATABASE_NAME"`

	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
}


type Config struct {
	Bind      string `env:"BIND"`
	Port      int    `env:"PORT"`
	Database  DatabaseArgs

	Extras	  env.EnvSet
}

// LoadEnv sets up Config struct from environment variables.
func (this *Config) LoadEnv() {
	extras, err := env.UnmarshalFromEnviron(this)
	if nil != err {
		log.Fatal("Cannot load config from env: ", err)
	}

	this.Extras = extras
}


// DatabaseArgs detects and returns database type and connection arguments.
func (a *DatabaseArgs) Args() (string, string, error) {
	if "" == a.Type {
		return "", "", errors.New("Database configuration not specified.")
	}

	dbType := strings.ToLower(a.Type)
	if strings.HasPrefix(dbType, "sqlite") {
		dbType = "sqlite3"
	}

	switch dbType {
	case "sqlite3":
		return dbType, a.Name, nil
	case "postgres":
		return dbType, a.getUrl(), nil
	default:
		return "", "", errors.New("Cannot detect database type.")
	}
}

func (a *DatabaseArgs) getUrl() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", a.Type, a.User, a.Password, a.Host, a.Port, a.Name)
}

var settings = &Config{
	Bind: "127.0.0.1",
	Port: 8000,
	Database: DatabaseArgs{},
}

func GetSettings() *Config {
	return settings
}
