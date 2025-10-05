/*
TeaBrewSim 設定管理
設定ファイルの読み込みと管理
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config アプリケーション設定
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Logging  LoggingConfig  `json:"logging"`
	Tea      TeaConfig      `json:"tea"`
}

// ServerConfig サーバー設定
type ServerConfig struct {
	Port         string `json:"port"`
	Host         string `json:"host"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

// DatabaseConfig データベース設定
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

// LoggingConfig ログ設定
type LoggingConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"`
	Output string `json:"output"`
}

// TeaConfig お茶関連設定
type TeaConfig struct {
	DefaultTemperature float64            `json:"default_temperature"`
	DefaultTime        int                 `json:"default_time"`
	TeaTypes          map[string]TeaType  `json:"tea_types"`
}

// TeaType お茶の種類設定
type TeaType struct {
	Name        string  `json:"name"`
	MinTemp     float64 `json:"min_temp"`
	MaxTemp     float64 `json:"max_temp"`
	MinTime     int     `json:"min_time"`
	MaxTime     int     `json:"max_time"`
	Description string  `json:"description"`
}

// LoadConfig 設定ファイルを読み込む
func LoadConfig(configPath string) (*Config, error) {
	// デフォルト設定
	config := &Config{
		Server: ServerConfig{
			Port:         "8080",
			Host:         "localhost",
			ReadTimeout:  30,
			WriteTimeout: 30,
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "teabrewsim",
			Password: "password",
			DBName:   "teabrewsim",
			SSLMode:  "disable",
		},
		Logging: LoggingConfig{
			Level:  "info",
			Format: "json",
			Output: "stdout",
		},
		Tea: TeaConfig{
			DefaultTemperature: 85.0,
			DefaultTime:        120,
			TeaTypes: map[string]TeaType{
				"green": {
					Name:        "緑茶",
					MinTemp:     70.0,
					MaxTemp:     85.0,
					MinTime:     60,
					MaxTime:     180,
					Description: "日本茶の代表的なお茶",
				},
				"black": {
					Name:        "紅茶",
					MinTemp:     90.0,
					MaxTemp:     100.0,
					MinTime:     120,
					MaxTime:     300,
					Description: "発酵させたお茶",
				},
				"oolong": {
					Name:        "烏龍茶",
					MinTemp:     85.0,
					MaxTemp:     95.0,
					MinTime:     90,
					MaxTime:     240,
					Description: "半発酵のお茶",
				},
			},
		},
	}

	// 設定ファイルが存在する場合は読み込み
	if _, err := os.Stat(configPath); err == nil {
		data, err := os.ReadFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("設定ファイルの読み込みに失敗: %v", err)
		}

		if err := json.Unmarshal(data, config); err != nil {
			return nil, fmt.Errorf("設定ファイルの解析に失敗: %v", err)
		}
	}

	return config, nil
}

// SaveConfig 設定をファイルに保存
func SaveConfig(config *Config, configPath string) error {
	// ディレクトリが存在しない場合は作成
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("ディレクトリの作成に失敗: %v", err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("設定のシリアライズに失敗: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("設定ファイルの書き込みに失敗: %v", err)
	}

	return nil
}

// GetConfigPath 設定ファイルのパスを取得
func GetConfigPath() string {
	// 環境変数から設定ファイルパスを取得
	if path := os.Getenv("TEABREWSIM_CONFIG"); path != "" {
		return path
	}
	
	// デフォルトパス
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "./config.json"
	}
	
	return filepath.Join(homeDir, ".teabrewsim", "config.json")
}
