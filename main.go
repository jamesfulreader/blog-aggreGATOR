package main

import (
	"fmt"
	"log"

	"github.com/jamesfulreader/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("read config : %+v\n", cfg)

	err = cfg.SetUser("james")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("read config again: %+v\n", cfg)
}
