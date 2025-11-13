package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/yourusername/process-monitor-cli/internal/model"
)

type ConfigService struct {
	configPath string
	config     *model.Config
}

func NewConfigService() *ConfigService {
	return &ConfigService{
		configPath: ".env",
		config:     model.DefaultConfig(),
	}
}

func (c *ConfigService) LoadConfig() error {
	if err := godotenv.Load(c.configPath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	c.config.APIKey = os.Getenv("GEMINI_API_KEY")
	c.config.VTAPIKey = os.Getenv("VIRUSTOTAL_API_KEY")
	return nil
}

func (c *ConfigService) SaveAPIKey(key string) error {
	vtKey := c.GetVTAPIKey()
	content := fmt.Sprintf("GEMINI_API_KEY=%s\n", key)
	if vtKey != "" {
		content += fmt.Sprintf("VIRUSTOTAL_API_KEY=%s\n", vtKey)
	}
	return os.WriteFile(c.configPath, []byte(content), 0600)
}

func (c *ConfigService) SaveVTAPIKey(key string) error {
	geminiKey := c.GetAPIKey()
	content := ""
	if geminiKey != "" {
		content += fmt.Sprintf("GEMINI_API_KEY=%s\n", geminiKey)
	}
	content += fmt.Sprintf("VIRUSTOTAL_API_KEY=%s\n", key)
	return os.WriteFile(c.configPath, []byte(content), 0600)
}

func (c *ConfigService) GetVTAPIKey() string {
	return c.config.VTAPIKey
}

func (c *ConfigService) GetAPIKey() string {
	return c.config.APIKey
}

func (c *ConfigService) DeleteAPIKey() error {
	return os.Remove(c.configPath)
}

func (c *ConfigService) ValidateAPIKey(key string) bool {
	return len(strings.TrimSpace(key)) > 0
}
