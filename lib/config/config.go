package config

import (
	"fmt"
	"os"
)

// Config содержит всю конфигурацию приложения
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Logging  LoggingConfig  `yaml:"logging"`
}

// ServerConfig содержит настройки HTTP-сервера
type ServerConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}

// DatabaseConfig содержит настройки базы данных
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// LoggingConfig содержит настройки логирования
type LoggingConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

// LoadConfig загружает конфигурацию из YAML-файла
func LoadConfig(configPath string) (*Config, error) {
	// Проверяем, существует ли файл
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("конфигурационный файл не найден: %s", configPath)
	}

	// Читаем файл
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурационного файла: %v", err)
	}

	// Парсим YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("ошибка парсинга YAML: %v", err)
	}

	return &config, nil
}
