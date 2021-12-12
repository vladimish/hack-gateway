package cfg

import "os"

type Config struct {
	AuthKey string
	DBUser  string
	DBPass  string
	DBAddr  string
}

var cfg Config

func GetConfig() *Config {
	return &cfg
}

func init() {
	var res bool
	cfg.AuthKey, res = os.LookupEnv("AUTH_KEY")
	if !res {
		panic("Variable AUTH_KEY is not defined.")
	}

	cfg.DBUser, res = os.LookupEnv("DB_USER")
	if !res {
		panic("Variable DB_USER is not defined.")
	}

	cfg.DBPass, res = os.LookupEnv("DB_PASS")
	if !res {
		panic("Variable DB_PASS is not defined.")
	}

	cfg.DBAddr, res = os.LookupEnv("DB_ADDR")
	if !res {
		panic("Variable DB_ADDR is not defined.")
	}
}
