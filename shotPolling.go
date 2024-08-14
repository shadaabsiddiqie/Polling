package main

import (
	"os"
	"time"
)

func ShotPolling() {
	for {
		DataToWrite, _ = os.ReadFile("data/file.txt")
		time.Sleep(time.Second)
	}
}
