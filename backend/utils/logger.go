package utils

import (
	"log"
	"time"
)

func LogRequest(method, path string, status int) {
	log.Printf("[%s] %s %s - %d", time.Now().Format(time.RFC3339), method, path, status)
}
