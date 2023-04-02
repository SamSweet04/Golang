package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string
	DB   struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

func loadConfig() *Config {
	viper.SetDefault("port", "8080")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "Haker15987")
	viper.SetDefault("db.name", "book-store")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/bookstore")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
}

return config
}
