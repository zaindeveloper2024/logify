package main

import (
	"github.com/zaindeveloper2024/logify/logify"
)

func main() {
	log := logify.New()
	log.SetLevel(logify.Debug)
	log.DebugF("Hello %s", "debug")
	log.Info("Hello info")
	log.SetLevel(logify.Info)
	log.DebugF("Hello %s", "debug")
}
