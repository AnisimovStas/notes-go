package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEVELOPMENT" env-default:"false"`
	Listen        struct {
		Type       string `env:"LISTEN_TYPE" env-default:"port" env-description:"PORT or SOCK. if 'sock' then env SOCKET_FILE is required"`
		BindIP     string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port       string `env:"PORT" env-default:"10000"`
		SocketFile string `env:SOCKET_FILE env-default:"app.sock"`
	}
	AppConfig struct {
		LogLevel  string `env:"LOG_LEVEL"`
		AdminUser struct {
			email    string `env:"ADMIN_EMAIL" env-default:"admin"`
			password string `env:"ADMIN_PASSWORD" env-default:"admin"`
		}
	}
}

var instanse *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {

		log.Print("gather config")

		instanse = &Config{}
		if err := cleanenv.ReadEnv(instanse); err != nil {
			helpText := "Notes project"
			help, _ := cleanenv.GetDescription(instanse, &helpText)
			log.Print(help)
			log.Fatal(help)
		}
	})
	return instanse
}
