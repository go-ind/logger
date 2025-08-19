package main

import "logger"

func main() {
	log := logger.New("info", "text")
	log.Info("hello")
	log.Warn("world")
	log.Error("boom")

	jsonLog := logger.New("info", "json")
	jsonLog.Info("structured")
}
