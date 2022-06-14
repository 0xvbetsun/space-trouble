// package configs uses for configuring different parts of our application
package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string `yaml:"port"`
	DB   struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"db"`
	API struct {
		URL string `yaml:"url"`
	} `yaml:"api"`
}

// LoadConfig loads configurations from received path and env variables
func LoadConfig(path string) (*Config, error) {
	conf := &Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return conf, err
	}
	if err := viper.Unmarshal(conf); err != nil {
		return conf, err
	}
	conf.Port = viper.GetString("PORT")
	conf.DB.Host = viper.GetString("POSTGRES_HOST")
	conf.DB.Password = viper.GetString("POSTGRES_PASSWORD")

	return conf, nil
}
