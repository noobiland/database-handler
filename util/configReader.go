package util

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	DBPath       string  `yaml:"db_path"`
	BackupPath   string  `yaml:"backup_path"`
	InitPath     string  `yaml:"init_path"`
	InitDataPath *string `yaml:"init_data_path"` // use pointer to detect `null`
}

type Config struct {
	Databases map[string]DBConfig `yaml:"databases"`
}

func GetSupportedDbs() []string {
	var cfg Config
	if err := yaml.Unmarshal(GetConfigs(), &cfg); err != nil {
		panic(err)
	}
	result := []string{}
	for name, db := range cfg.Databases {
		// TODO: Proper Verbose
		if false {
			slog.Info(fmt.Sprintf("Database: %s", name))
			slog.Info(fmt.Sprintf("	DB Path: %s", db.DBPath))
			slog.Info(fmt.Sprintf("	Backup Path: %s", db.BackupPath))
			slog.Info(fmt.Sprintf("	Init Path: %s", db.InitPath))
			if db.InitDataPath != nil {
				slog.Info(fmt.Sprintf("	Init Data Path: %s", *db.InitDataPath))
			} else {
				slog.Info("	Init Data Path: null")
			}
		}
		result = append(result, name)
	}
	return result
}

func GetConfigs() []byte {
	data, err := os.ReadFile("./configs/config.yaml")
	if err != nil {
		panic(err)
	}
	return data
}

func CheckDbIsSupported(candidate string) bool {
	for _, supported_db := range GetSupportedDbs() {
		if supported_db == candidate {
			slog.Info(fmt.Sprintf("%s database is supported", candidate))
			return true
		}
	}
	slog.Error(fmt.Sprintf("%s database is NOT supported", candidate))
	return false
}
