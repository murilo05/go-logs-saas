# Log Ingest SaaS (Go) 🚀

> ⚠️ Work in progress – follow the dev journey in the DevLog!

This project is a **log ingestion service** built in **Go**, following **Clean Architecture** principles.  
It’s designed to handle high-scale logging for multiple applications, supporting structured logs, validation, and queuing for async processing.  

---
## 🛠 Tech Stack

**Language:** Go  
**Framework:** Gin (HTTP), gRPC planned  
**Architecture:** Clean Architecture  
**Queue:** AWS SQS  
**Database:** PostgreSQL (metadata), MongoDB/DynamoDB (logs)  
**Containerization:** Docker  
**Orchestration:** Kubernetes  
---

## ⚡ Current Progress

**Day #1:** Project setup, folder structure, main.go, environment variables, logger initialization, basic DI setup (Repository, UseCase, Handler).  
**Day #2:** HTTP server setup, request DTOs, handler for **POST /logs**, input validation, structured logging, and error/success response handling.  
UseCase and Repository implementation will be completed in the next days.  

---

## 📌 Next Steps

Implement **UseCase logic**  
Connect **Repository** for persistence  
Publish logs to **SQS** for async processing  
Add **unit tests** and error handling improvements  
gRPC endpoints  

---