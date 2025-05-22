
# banking-ledger-service

**A reliable, scalable Golang backend service managing bank accounts and transactions with ACID consistency, asynchronous processing, and comprehensive testing.**

---

## Features

- Create bank accounts with initial balances  
- Deposit and withdraw funds asynchronously via Kafka message queue  
- Maintain ACID-compliant balances in PostgreSQL  
- Store detailed transaction logs in MongoDB  
- Horizontal scalability with Kafka-backed transaction processing  
- RESTful API endpoints for external interaction  
- Comprehensive unit, integration, and feature tests

---

## Getting Started

### Prerequisites

- Go 1.20+  
- Docker & Docker Compose  
- Kafka & Zookeeper (via Docker Compose)  
- PostgreSQL  
- MongoDB  

### Running the Service

1. Clone the repo  
   ```bash
   git clone https://github.com/likhithkp/banking-ledger-service.git
   cd banking-ledger-service
   ```

2. Start all services with Docker Compose  
   ```bash
   docker-compose up -d
   ```

3. The API will be available at `http://localhost:3001`

---

## API Endpoints

- `POST /accounts` — Create a new account  
- `POST /accounts/{id}/transactions` — Submit deposit/withdraw transaction  
- `GET /accounts/{id}` — Get account details  
- `GET /accounts/{id}/transactions` — Get transaction history  

---

## Architecture Overview

- **API Server:** Handles HTTP requests, validates and publishes transactions to Kafka  
- **Transaction Processor:** Background worker consuming Kafka messages to update DBs  
- **PostgreSQL:** Stores account balances with ACID guarantees  
- **MongoDB:** Stores immutable transaction logs (ledger)  
- **Kafka:** Manages asynchronous transaction request queuing and delivery  

---

## Testing

- Unit tests for services and handlers  
- Integration tests with real DB and Kafka  
- Feature tests covering end-to-end flows  

Run tests with:  
```bash
go test ./...
```

---

## License

MIT License © 2025 likhithkp
