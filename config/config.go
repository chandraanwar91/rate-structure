package config

import "github.com/jinzhu/configor"

var Config = struct {
	App struct {
		ENV      string
		HttpAddr string
		HttpPort string
	}

	DB struct {
		Driver   string
		Host     string
		Port     string `default:"3306"`
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
	}

	DBHotel struct {
		Driver   string
		Host     string
		Port     string `default:"3306"`
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
	}

	SENTRY struct {
		SentryDSN string
	}
}{}

func init() {
	configor.Load(&Config, "config.yaml")
}
