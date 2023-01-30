package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	database DataBaseSettings
}

type DataBaseSettings struct {
	username string
	password string
	port string
	host string
	database_name string
	require_ssl string
}
