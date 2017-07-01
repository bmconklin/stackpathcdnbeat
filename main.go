package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/bmconklin/stackpathcdnbeat/beater"
)

func main() {
	err := beat.Run("stackpathcdnbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
