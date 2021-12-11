package main

import (
	"encoding/json"
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
	//ctx := context.Background()
	execName, _ := os.Executable()
	versionMap := map[string]string{
		"Version":   Version,
		"BuildTime": BuildTime,
		"VCSTag":    VCSTag,
		"ExecName":  filepath.Base(execName),
	}
	versionString, _ := json.Marshal(versionMap)
	log.Println(string(versionString))

	log.Println("Version:", Version)
	log.Println("BuildTime:", BuildTime)
	log.Println("VCSTag:", VCSTag)
	log.Println("ExecName:", filepath.Base(execName))
}
