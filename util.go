package main

import (
	"log"
	"os"
	"runtime"
)

func getMinecraftPath() string {
	rt := runtime.GOOS
	switch rt {
	case "windows":
		return os.Getenv("APPDATA") + "\\.minecraft"
	case "darwin":
		return os.Getenv("HOME") + "/Library/Application Support/minecraft"
	case "linux":
		return os.Getenv("HOME") + "/.minecraft" //TODO: Test this
	default:
		log.Fatal("Unsupported OS\n")
		return ""
	}
}
