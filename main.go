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

	"gopkg.in/yaml.v3"
)

type Config struct {
	LogFile       string   `yaml:"logfile"`
	Filters       []string `yaml:"filters"`
	SleepInterval string   `yaml:"sleep_interval"`
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	cfg, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	file, err := os.Open(cfg.LogFile)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	stat, _ := file.Stat()
	offset := stat.Size()
	file.Seek(offset, 0)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	filterCounts := make(map[string]int)

	// Parse interval
	interval, _ := time.ParseDuration(cfg.SleepInterval)

	fmt.Printf("Watching %s ...\n", cfg.LogFile)

	for {
		select {
		case <-sig:
			fmt.Println("\nShutting down...")
			for f, count := range filterCounts {
				fmt.Printf("%s: %d\n", f, count)
			}
			return
		default:
			reader := bufio.NewReader(file)
			line, err := reader.ReadString('\n')
			if err == nil {
				line = strings.TrimSpace(line)
				matched := false
				for _, f := range cfg.Filters {
					if strings.Contains(line, f) {
						fmt.Printf("[%s] %s\n", f, line)
						filterCounts[f]++
						matched = true
						break
					}
				}
				if !matched {
					fmt.Println(line)
				}
			} else {
				time.Sleep(interval)
			}
		}
	}
}
