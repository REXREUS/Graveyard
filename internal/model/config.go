package model

type Config struct {
	APIKey    string
	VTAPIKey  string
}

func DefaultConfig() *Config {
	return &Config{
		APIKey:   "",
		VTAPIKey: "",
	}
}
