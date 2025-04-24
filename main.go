package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"syscall"
	"unsafe"
)

const (
	SPI_SETDESKWALLPAPER = 0x0014
	SPIF_UPDATEINIFILE   = 0x01
	SPIF_SENDCHANGE      = 0x02
)

func SetWallpaper(Path string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	systemParametersInfo := user32.NewProc("SystemParametersInfoW")
	ret, _, err := systemParametersInfo.Call(
		SPI_SETDESKWALLPAPER,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(Path))),
		SPIF_UPDATEINIFILE|SPIF_SENDCHANGE,
	)

	if ret == 0 {
		log.Fatal("Krooz Fatel: ", err)
	}
}

func main() {
	channel := make(chan bool)
	urlPtr := flag.String("url", "", "URL to process")
	pathPtr := flag.String("path", "", "Path to process")
	initPtr := flag.Bool("r", false, "Initialize application")
	envPtr := flag.Bool("env", false, "Get Current Path")
	setEnvPtr := flag.String("set", "", "Set Current Path")

	flag.Parse()

	if *urlPtr != "" {
		EnsureMainPath()

		URL, err := url.ParseRequestURI(*urlPtr)
		if err != nil {
			log.Fatal("This is not a valid URL...")
			return
		}
		go LoadingScreen(channel)
		PATH := Downloader(URL.String())
		channel <- true
		cls()
		if PATH != "" {
			SetWallpaper(PATH)
		}
	}

	if *pathPtr != "" {
		if Exists(*pathPtr) {
			SetWallpaper(*pathPtr)
		} else {
			log.Fatal("Bad Path:")
		}
	}

	if *initPtr {
		EnsureMainPath()

		path := RandomFromFile()
		if path != "" {
			SetWallpaper(path)
		}
	}

	if *envPtr {
		EnsureMainPath()
		path, err := GetWallpaperPath()
		if err != nil {
			fmt.Println("Bad Param:", err)
		} else {
			fmt.Println("path:", path)
		}

	}
	if *setEnvPtr != "" {
		EnsureMainPath()
		if SetPath(*setEnvPtr) {
			fmt.Println("Path Changed To: ", *setEnvPtr)
		}
	}

	if flag.NFlag() == 0 {
		fmt.Println("No flags provided. Use --help for usage.")
	}
}
