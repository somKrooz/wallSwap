package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func SetPath(path string) bool {
	Path := GetConfigFile()
	data := fmt.Sprint("path = ", path)
	err := os.WriteFile(Path, []byte(data), 0644)
	return err == nil

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
