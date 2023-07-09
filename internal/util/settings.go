package util

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Settings struct {
	AutoImport         bool `json:"autoImport"`
	ShowNativeTitleBar bool `json:"showNativeTitleBar"`
	path               string
}

func InitializeSettings() Settings {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln(err)
	}

	settingsPath := filepath.FromSlash(configDir + "/geemo/settings.json")

	// If settings file doesn't exist, create it
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		return createNewSettingsFile(configDir)
	}

	settings := Settings{
		path: settingsPath,
	}
	file, err := os.Open(settingsPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.NewDecoder(file).Decode(&settings)
	if err != nil {
		log.Fatalln(err)
	}

	return settings
}

func createNewSettingsFile(configDir string) Settings {
	// First create the config directory if it doesn't exist
	path := filepath.FromSlash(configDir + "/geemo")

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	// Then create the settings file
	path = filepath.FromSlash(path + "/settings.json")

	file, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	showNativeTitleBar := true
	switch userOS := runtime.GOOS; userOS {
	case "windows":
		showNativeTitleBar = false
	}

	// Write default settings to file
	settings := Settings{
		AutoImport:         false,
		ShowNativeTitleBar: showNativeTitleBar,
		path:               path,
	}
	settingsJson, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = file.Write(settingsJson)
	if err != nil {
		log.Fatalln(err)
	}

	return settings
}

func (s *Settings) Save() {
	// os.Create() truncates the file if it already exists
	file, err := os.Create(s.path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	settingsJson, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = file.Write(settingsJson)
	if err != nil {
		log.Fatalln(err)
	}
}
