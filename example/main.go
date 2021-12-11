package main

import (
	"encoding/json"
    "fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	Version   = "v0.1.0"
	BuildTime = "latest"
	VCSTag    = "latest"
)

func main() {
	execName, _ := os.Executable()
	versionMap := map[string]string{
		"Version":   Version,
		"BuildTime": BuildTime,
		"VCSTag":    VCSTag,
		"ExecName":  filepath.Base(execName),
	}
	versionString, _ := json.Marshal(versionMap)
	log.Println(string(versionString))

    log.Println(fmt.Sprintf("%s / %s / %s / %s", filepath.Base(execName), Version, BuildTime, VCSTag))

	log.Println("Version:", Version)
	log.Println("BuildTime:", BuildTime)
	log.Println("VCSTag:", VCSTag)
	log.Println("ExecName:", filepath.Base(execName))
}
