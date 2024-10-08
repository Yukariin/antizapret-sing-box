package main

import (
	"log"

	"github.com/Yukariin/antizapret-sing-box/geosite_antizapret"
)

func main() {
	generator := geosite_antizapret.NewGenerator()

	if err := generator.GenerateAndWrite(); err != nil {
		log.Fatalf("Failed to generate and write data: %v", err)
	}
}
