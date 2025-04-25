package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetWallpaperPath() (string, error) {
	EnsureMainPath()

	path := GetConfigFile()
	if DirExists(path, false) {
		Config := string(GetConfigFile())
		path, err := os.ReadFile(Config)
		if err != nil {
			return "", fmt.Errorf("krooz error: %w", err)
		}
		if len(path) <= 0 {
			return "", fmt.Errorf("krooz error: path is empty")
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
		fmt.Println("Bad Parm: invalid Directory name")
		os.Exit(0)
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

func Downloader(url string) string {
	EnsureMainPath()

	if strings.Contains(url, "github.com") {
		url = strings.Replace(url, "github.com", "raw.githubusercontent.com", 1)
		url = strings.Replace(url, "/blob/", "/", 1)
	}
	fmt.Println(url)
	sep := strings.Split(url, "/")
	ext := filepath.Ext(sep[len(sep)-1])

	res, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	Dir := GetDownloadDirectory()
	path := filepath.Join(Dir, "Wallpaper"+ext)
	_, err = os.Stat(path)

	if err == nil {
		err = os.Remove(path)
		if err != nil {
			return ""
		}
	}

	err = os.WriteFile(path, body, 0644)
	if err != nil {
		return ""
	}
	return path
}
