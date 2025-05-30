# Ollama UI

A simple web interface to view models installed via [Ollama](https://ollama.ai). 

Built with **Go** + **HTMX** for fast, dynamic updates without JavaScript.

![Screenshot](screenshots/screenshot.png)

## Features

- ✅ List all installed Ollama models
- 🚀 Fast loading with minimal dependencies
- 🧼 Clean, readable HTML with HTMX-powered updates
- 📦 Easy to run locally or containerize

## Requirements

- Go 1.20+
- [Ollama](https://ollama.ai/download)  running locally

## How to Run

```bash
# Start Ollama in one terminal
ollama serve

# In another terminal
go run main.go
