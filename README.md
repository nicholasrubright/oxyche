# Oxyche CLI 💽

A CLI tool for managing a proxy cache server. Built with Go and Docker, this tool lets you easily start and manage a caching proxy server from the command line.

## Overview

This project provides a proxy server with caching capabilities:
- Built using Go's Gin framework for high performance
- Runs in Docker for easy deployment
- Currently uses in-memory cache (Redis support coming soon)
- Managed through a simple CLI interface

## Technologies Used

- Go 1.23
- Docker
- Gin (HTTP framework)
- Cobra (CLI framework)

## Project Structure

```
├── cmd/          # CLI tool code
├── internal/     # Internal logic code
├── server/       # Proxy server code
├────── Dockerfile    
└────── config.yml    
```

## Getting Started

1. Clone the repository
```bash
git clone https://github.com/nicholasrubright/oxyche
cd oxyche
```

2. Install the Executable
```bash
cd oxyche
go install oxyche
```

3. Run the proxy server
```bash
oxyche start --port <port> --origin <origin>
```

## Basic Commands

```bash
# Start the proxy server
oxyche start

# Check server status
oxyche status

# Stop the server
oxyche stop
```

## Configuration

The server can be configured through `server/config.yml`:

```yaml
server:
  port: 8080
  origin: http://localhost
```

## Coming Soon

- Redis cache support
- More configuration options
- Performance metrics

## Development

Requirements:
- Go 1.23 or higher
- Docker
- Make (optional)
