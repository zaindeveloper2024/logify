package main

import (
	"github.com/zaindeveloper2024/logify/logify"
)

func main() {
	log := logify.New()
	log.Debug("Hello world")
	log.Info("Hello world")
}