package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func cls() {
	fmt.Print("\r\033[K")
}
func GetDownloadDirectory() string {
	Temp := os.TempDir()
	return filepath.Join(Temp, "Krooz")
}

func GetConfigFile() string {
	Temp := os.TempDir()
	return filepath.Join(Temp, "Krooz", "krooz.txt")
}

func DirExists(Path string, isdir bool) bool {
	if _, err := os.Stat(Path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		if isdir {
			err := os.Mkdir(Path, os.ModeDevice)
			return err == nil
		} else {
			err := os.WriteFile(Path, []byte(""), 0644)
			return err == nil
		}

	}
	return false
}

func Exists(Path string) bool {
	if _, err := os.Stat(Path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

func Downloader(url string) string {
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

	// Check if the file exists and overwrite it
	_, err = os.Stat(path)
	if err == nil { // File exists
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

func LoadingScreen(pchan <-chan bool) {
	spinner := []string{"", "k", "kr", "kro", "kroo", "krooz", "krooz.", "krooz.."}
	i := 0

	for {
		select {
		case <-pchan:
			cls()
			return
		default:
			fmt.Printf("\r\033[K%s", spinner[i%len(spinner)])
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}
}
