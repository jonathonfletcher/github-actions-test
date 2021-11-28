package main

import (
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

	log.Println("Version:", Version)
	log.Println("BuildTime:", BuildTime)
	log.Println("VCSTag:", VCSTag)
	log.Println("execName:", filepath.Base(execName))
}
