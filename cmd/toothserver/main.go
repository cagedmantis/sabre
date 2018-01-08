package main

import (
	envflag "github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting tooth server...")

	var binDir string
	envflag.StringVar(&binDir, "toothserver_bin_dir", "/data/toothserver", "path to bin directory")
	envflag.Parse()

	log.WithFields(log.Fields{
		"bin.directory": binDir,
	}).Info("config settings")

	log.Info("exiting tooth server")
}
