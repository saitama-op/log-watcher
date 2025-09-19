# 📝 Log File Watcher (Go)

A simple **log file monitoring tool** written in Go.  
It works like `tail -f`, but with extra powers: highlights keywords (e.g., **ERROR**, **WARNING**), shows stats, and supports YAML configuration.  

---

## 🚀 Features
- Real-time log monitoring  
- Keyword filtering (configurable in `config.yaml`)  
- Highlights `ERROR` and `WARNING` lines  
- Graceful shutdown with summary stats  
- Configurable log file path & refresh interval  

---

## 📂 Project Structure
```
log-watcher/
├── main.go        # Main program
├── config.yaml    # Configuration file
├── go.mod
└── README.md
```

---

## ⚙️ Configuration (`config.yaml`)
```yaml
logfile: "/var/log/syslog"   # Path to the log file
filters:
  - "ERROR"
  - "WARNING"
sleep_interval: 1s           # Interval to check for new logs
```

---

## 🛠 Installation & Usage

### Clone the repo
```bash
git clone https://github.com/saitama-op/log-watcher.git
cd log-watcher
```

### Install dependencies
```bash
go mod tidy
```

### Run the program
```bash
go run main.go
```

It will read settings from `config.yaml` automatically.  

---

## 📊 Example Output
```
Watching /var/log/syslog ...
Normal log line
[ERROR] Disk full on /dev/sda1
[WARNING] High memory usage
```

On exit (`Ctrl+C`):
```
Shutting down...
ERROR: 3
WARNING: 2
```

---

## 📦 Build Executable
```bash
go build -o logwatcher
./logwatcher
```

---

## 🧑‍💻 Future Improvements
- Handle log rotation  
- Export stats as Prometheus metrics  
- JSON/Color output for better readability  
- Docker support  

---

## 📜 License
MIT License – free to use and modify.  
