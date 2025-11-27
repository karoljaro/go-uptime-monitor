# Go Uptime Monitor 🚀

> [!NOTE]
> This project is an educational exercise – I'm learning Go while building a real-time uptime monitoring system.

[![Go](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/yourusername/go-uptime-monitor)](https://github.com/yourusername/go-uptime-monitor/issues)
[![GitHub stars](https://img.shields.io/github/stars/yourusername/go-uptime-monitor?style=social)](https://github.com/yourusername/go-uptime-monitor/stargazers)

Go Uptime Monitor is a real-time monitoring and alerting backend written in Go.
It tracks availability and response times of websites and APIs using standard HTTP/1.1 requests and provides alerting when monitored targets go down.

Features

- Monitor multiple URLs concurrently using Go goroutines
- Store monitoring results and statistics (in-memory or SQLite)
- Alerting system for downtime events
- RESTful API for managing targets, results, alerts, and statistics
- Clean Architecture implementation for maintainable and scalable code
- Simple HTTP/1.1 support (no HTTP/3/QUIC yet)
- Designed for easy expansion to HTTP/2 and HTTP/3

Getting Started

Prerequisites

- Go 1.21+
- Optional: SQLite for persistent storage

Installation

git clone https://github.com/yourusername/go-uptime-monitor.git
cd go-uptime-monitor
go mod tidy
go run cmd/server/main.go

Run Tests

go test ./...

Project Structure

go-uptime-monitor/
├─ cmd/server/           # Entry point for HTTP server
├─ internal/
│   ├─ domain/           # Entities and interfaces (Target, Result, Alert)
│   ├─ usecase/          # Business logic (monitoring, alerts, stats)
│   ├─ infrastructure/   # HTTP client, storage (in-memory / SQLite)
│   └─ interface/        # HTTP handlers
├─ go.mod
└─ README.md

API Endpoints

Method | Endpoint | Description
------ | -------- | -----------
POST   | /targets | Add a new target to monitor
GET    | /targets | List all monitored targets
GET    | /results/{id} | Get monitoring results
GET    | /alerts | View active alerts
GET    | /stats/{id} | Get uptime statistics
GET    | /ping | Health check

Example Target JSON

{
  "url": "https://api.github.com",
  "interval": 10
}

Usage Example

Add a target:

curl -X POST http://localhost:8080/targets \
-H "Content-Type: application/json" \
-d '{"url": "https://api.github.com", "interval": 10}'

Get results:

curl http://localhost:8080/results/target-id

Future Features

- HTTP/2 and HTTP/3 (QUIC) support
- Persistent storage in SQLite/PostgreSQL
- Web dashboard (React/TypeScript) for visualization
- User authentication (JWT)
- Alert integrations: email, Slack, webhook

License

MIT License © 2025 [Your Name]
