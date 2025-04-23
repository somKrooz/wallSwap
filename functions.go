package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func GetWallpaperPath() (string, error) {
	path := GetConfigFile()
	if DirExists(path, false) {
		Config := string(GetConfigFile())
		path, err := os.ReadFile(Config)
		if err != nil {
			return "", fmt.Errorf("krooz error: %w", err)
		}
		parts := strings.Split(string(path), "=")
		return strings.TrimSpace(parts[1]), nil
	}
	return "", fmt.Errorf("krooz error: failed to exec")
}

func RandomFromFile() string {
	var Images []string
	path, err := GetWallpaperPath()
	if err != nil {
		return ""
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return ""
	}
	ImageValidation := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
	}

	for _, f := range files {
		ext := strings.ToLower(filepath.Ext(f.Name()))
		if ImageValidation[ext] {
			Images = append(Images, f.Name())
		}
	}
	randImg := Images[rand.Intn(len(Images))]
	return filepath.Join(path, randImg)
}
