# NBA Stats Microservice

This repository contains a modular backend system for NBA player statistics. It includes:

- **Go-based REST API** (`/backend-go/`) for querying player data and stats
- **Python scraper** (`/scripts/`) to ingest and update data from the [BallDontLie API](https://www.balldontlie.io) into MongoDB

---

## Purpose

This project is the backend foundation for a future **multi-sport analytics platform**. It is designed to be split into scalable microservices with clean data ingestion and querying layers.

---

## Structure

- [`/backend-go`](./backend-go): Lightweight Go HTTP API for consuming NBA data
- [`/scripts`](./scripts): Python scripts for scraping and populating the MongoDB database

---

## Usage

For setup, usage, and endpoint/API details, refer to the `README.md` inside each subdirectory:

- 📘 [`backend-go/README.md`](./backend-go/README.MD)
- 📘 [`scripts/README.md`](./scripts/README.md)
