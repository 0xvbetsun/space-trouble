// package configs uses for configuring different parts of our application
package configs

import "github.com/spf13/viper"

// LoadConfig loads configurations from received path and env variables
func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.AddConfigPath("deployments")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	return viper.MergeInConfig()
}
