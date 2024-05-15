package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

func appDataPath(appName string) (string, error) {
	var appDataPath string

	switch runtime.GOOS {
	case "windows":
		// For Windows, use %APPDATA% for Roaming or %LOCALAPPDATA% for Local
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			return "", fmt.Errorf("LOCALAPPDATA environment variable is not set")
		}
		appDataPath = filepath.Join(localAppData, appName)

	case "darwin":
		// For macOS, use ~/Library/Application Support
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		appDataPath = filepath.Join(homeDir, "Library", "Application Support", appName)

	case "linux":
		// For Linux, use ~/.local/share
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		appDataPath = filepath.Join(homeDir, ".local", "share", appName)

	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	err := os.MkdirAll(appDataPath, 0755)
	if err != nil {
		return "", err
	}

	return appDataPath, nil
}

func database(dbname string) (*sql.DB, error) {
	storePath, err := appDataPath("blaze-assistant")
	if err != nil {
		return nil, err
	}

	return sql.Open("sqlite3", filepath.Join(storePath, dbname))
}
