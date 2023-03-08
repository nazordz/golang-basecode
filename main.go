package main

import (
	"github.com/devgoorita/golang-basecode/pkg"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	server := InitializedServer()

	err := server.Run(":" + pkg.GodotEnv("PORT"))
	if err != nil {
		log.Error(err)
	}
}
