package config

import (
	"fmt"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	database DataBaseSettings
}

type DataBaseSettings struct {
	username      string
	password      string
	port          string
	host          string
	database_name string
	require_ssl   string
}

func (d *DataBaseSettings) WithoutDb() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s", d.username, d.password, d.host, d.port)
}

func (d *DataBaseSettings) WithDb() string {
	return fmt.Sprintf("%s/%s", d.WithoutDb(), d.database_name)
}
