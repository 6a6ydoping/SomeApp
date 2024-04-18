package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
	"time"
)

type Config struct {
	HTTP ServerConfig `yaml:"http"`
	DB   DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"prt"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"DB_PASSWORD"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = cleanenv.ReadEnv(&cfg.DB)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
