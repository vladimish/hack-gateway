package cfg

import "os"

type Config struct {
	AuthKey string `json:"auth_key"`
}

var cfg Config

func GetConfig() Config {
	return cfg
}

func init() {
	var res bool
	cfg.AuthKey, res = os.LookupEnv("AUTH_KEY")
	if !res {
		panic("Variable TG_KEY is not defined.")
	}
}
