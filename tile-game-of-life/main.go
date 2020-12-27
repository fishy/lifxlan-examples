package main

import (
	"context"
	"log"

	"go.yhsif.com/lifxlan"
)

func checkContextError(err error) bool {
	return err != nil && err != context.Canceled && err != context.DeadlineExceeded
}

func main() {
	td := findDevice(lifxlan.Target(target))
	if td == nil {
		log.Fatal("No matching tile device found.")
	}
	log.Printf("Found %v", td)
	draw(td)
}
