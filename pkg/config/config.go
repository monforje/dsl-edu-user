package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ConfigDatabase struct {
	MongoURI string `yaml:"mongo_uri" env:"MONGO_URI" env-default:"mongodb://localhost:27017"`
	DBname   string `yaml:"db_name" env:"DB_NAME" env-default:"dsl_users"`
	COLname  string `yaml:"collection_name" env:"COL_NAME" env-default:"users"`
}

func New(source string) *ConfigDatabase {
	var cfg ConfigDatabase

	switch source {
	case "env":
		if err := godotenv.Load(); err != nil {
			log.Println(".env файл не найден, используются значения по умолчанию")
		}

		if err := cleanenv.ReadEnv(&cfg); err != nil {
			log.Fatalf("ошибка чтения ENV переменных: %v", err)
		}

	case "yaml":
		if err := cleanenv.ReadConfig("config.yaml", &cfg); err != nil {
			log.Fatalf("ошибка чтения config.yaml: %v", err)
		}

	default:
		log.Fatalf("неизвестный источник конфигурации: %s (используй 'env' или 'yaml')", source)
	}

	return &cfg
}
