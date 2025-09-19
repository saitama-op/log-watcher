# Log File Watcher (Go)

A simple log file monitoring tool written in Go.  
It works like `tail -f` but highlights **ERROR** and **WARNING** lines.

## 🚀 Features
- Real-time log monitoring
- Detects and highlights errors/warnings
- Graceful shutdown with summary stats

## 🛠 Usage
```bash
go run main.go /var/log/syslog
