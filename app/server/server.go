package main

import (
	"log"
	"os"

	"github.com/k-ueki/sublog/app/server/controller"
)

func main() {
	if err := controller.FetchBlogs(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
