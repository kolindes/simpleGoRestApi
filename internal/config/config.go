package config

import (
	"encoding/json"
	"os"
)

type MainConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DBConfig struct {
	Engine       string `json:"engine"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type JWTConfig struct {
	SecretKey      string `json:"secret"`
	ExpiresInHours int64  `json:"expires_in_hours"`
}

type LoggingConfig struct {
	Colorful bool
	LogLevel int
}

type Config struct {
	Main          MainConfig    `json:"main"`
	DB            DBConfig      `json:"db"`
	JWT           JWTConfig     `json:"jwt"`
	LoggingConfig LoggingConfig `json:"logging"`
}

func Load() (*Config, error) {
	configFile, err := os.Open("config.json")
	if err != nil {
		// If the config file does not exist, create a default one with default values
		if os.IsNotExist(err) {
			defaultConfig := &Config{
				Main: MainConfig{
					Host: "0.0.0.0",
					Port: "3000",
				},
				DB: DBConfig{
					Engine:       "sqlite3",
					MaxOpenConns: 100,
					MaxIdleConns: 1000,
				},
				JWT: JWTConfig{
					SecretKey:      "secret",
					ExpiresInHours: 1,
				},
				LoggingConfig: LoggingConfig{
					Colorful: false,
					LogLevel: 1,
				},
			}
			err := Write(defaultConfig)
			if err != nil {
				return nil, err
			}

			return defaultConfig, nil
		}

		// If there's an error other than the file not existing, return the error
		return nil, err
	}
	defer configFile.Close()

	config := new(Config)
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func Write(c *Config) error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(c)
	if err != nil {
		return err
	}

	return nil
}
