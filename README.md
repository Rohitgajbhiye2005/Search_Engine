# Mini Search Engine (Go)

A concurrent web crawler and search engine built using Go. The system crawls developer documentation pages, indexes their content, and allows users to search through the indexed dataset.

This project was built to understand how real search engines work internally — including crawling, indexing, and fast query retrieval.

## 🚀 Live Demo

**Live Application:** [Click here to use Mini Search Engine](https://search-engine-iktp.onrender.com/)

## Features

- Concurrent web crawler using **goroutines**
- Worker pool for parallel page fetching
- URL discovery and deduplication
- Crawls and indexes **10,000+ pages**
- Content stored in **PostgreSQL**
- Fast search using **PostgreSQL Full-Text Search (`tsvector` + GIN index)**
- REST API built with **Gin**
- Frontend built with HTML, CSS, and JavaScript
- Go backend serves the frontend files

## Architecture

Crawler → Content Extraction → Database Storage → Search Index → REST API → Web Frontend

## Tech Stack

- Go (Golang)
- Gin Web Framework
- HTML, CSS, JavaScript
- PostgreSQL
- Goroutines & Channels
- Worker Pool Pattern
- PostgreSQL Full-Text Search

## Running the Project Locally

### 1. Clone the repository

```bash
git clone https://github.com/Rohitgajbhiye2005/Mini_Search_Engine.git
cd Mini_Search_Engine
```

### 2. Configure the database

Make sure PostgreSQL is installed and running. Configure the database connection using the environment variables or configuration method expected by your application.

### 3. Start the backend and frontend

```bash
go run main.go
```

The Go server serves the frontend files and API endpoints.

Open your browser and visit:

http://localhost:8080

**Note:** Ensure that your frontend files are in the location expected by your Go server and that your database is configured correctly before starting the application.

### 4. Start crawling pages

With the application running, open another terminal and execute:

```bash
curl -X POST http://localhost:8080/crawl \
  -H "Content-Type: application/json" \
  -d '{
    "seeds": ["https://go.dev"],
    "max_pages": 10000
  }'
```

This sends a request to the crawler API to begin crawling from the specified seed URL, subject to your implementation's limits and configuration.

## Example Search Queries

- `channel`
- `goroutine`
- `concurrency`

## Repository

[View Mini Search Engine on GitHub](https://github.com/Rohitgajbhiye2005/Mini_Search_Engine)

