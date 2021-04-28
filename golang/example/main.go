package main

import log "github.com/kevinmalaverr/logger/golang"

func main() {
	log.InitWriteLog()

	log.Success("message")
	log.Info("message")
	log.Warn("message")
	log.Error("message")
}
