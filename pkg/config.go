package pkg

import (
	"github.com/joho/godotenv"
	"github.com/peterszarvas94/configloader"
)

type AppConfig struct {
	CONTENT_DIR string
	PUBLIC_DIR  string
	STATIC_DIR  string
}

var Config AppConfig

func LoadEnvs() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	err = configloader.Load(&Config)
	return err
}
