package main

import (
	"os"
	"time"
)

var prevFileData []byte

var DataToWrite []byte = []byte("default value")

func LongPolling() {
	for {
		latestFileData, _ := os.ReadFile("data/file.txt")
		if string(prevFileData) != string(latestFileData) {
			DataToWrite = latestFileData
		}
		time.Sleep(time.Second)
	}
}
