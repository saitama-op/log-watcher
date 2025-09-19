package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <logfile>")
		return
	}

	logFile := os.Args[1]
	file, err := os.Open(logFile)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Move to the end of file
	stat, _ := file.Stat()
	offset := stat.Size()
	file.Seek(offset, 0)

	// Setup graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	errorCount, warnCount := 0, 0

	fmt.Printf("Watching %s ...\n", logFile)

	for {
		select {
		case <-sig:
			fmt.Println("\nShutting down...")
			fmt.Printf("Summary: %d errors, %d warnings\n", errorCount, warnCount)
			return
		default:
			reader := bufio.NewReader(file)
			line, err := reader.ReadString('\n')
			if err == nil {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "ERROR") {
					fmt.Printf("[ERROR] %s\n", line)
					errorCount++
				} else if strings.Contains(line, "WARNING") {
					fmt.Printf("[WARN]  %s\n", line)
					warnCount++
				} else {
					fmt.Println(line)
				}
			} else {
				time.Sleep(1 * time.Second) // wait before checking again
			}
		}
	}
}
