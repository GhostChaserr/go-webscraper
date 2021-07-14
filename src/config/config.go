package config

import (
	"os"
)

type Config struct {
	DB_URL string
	PORT   string
}

func GetConfigs() Config {
	Config := new(Config)
	Config.DB_URL = "mongodb+srv://tridioUser:tridioUser@cluster0.spzhp.mongodb.net/tridio_development?retryWrites=true&w=majority"
	Config.PORT = os.Getenv("PORT")

	return *Config
}
