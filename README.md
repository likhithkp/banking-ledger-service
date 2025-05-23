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

## Tech Stack

- **Golang** for backend service  
- **PostgreSQL** for account balances (ACID)  
- **MongoDB** for transaction logs  
- **Kafka** for asynchronous transaction queuing  
- **Docker Compose** for container orchestration  
- **migrate/migrate** for DB migrations  

---

## Getting Started

### Prerequisites

- Go 1.20+  
- Docker & Docker Compose  

### Run the Service

1. Clone the repo  
   ```bash
   git clone https://github.com/likhithkp/banking-ledger-service.git
   cd banking-ledger-service
   ```

2. Start all services with Docker Compose  
   ```bash
   docker-compose up --build -d
   ```

3. Ensure PostgreSQL, MongoDB, Kafka are healthy. Run DB migrations if needed:  
   ```bash
   docker compose run --rm migrate
   ```

4. API will be available at  
   ```
   http://localhost:3001
   ```

### Directory Structure

```
.
├── config/             # Environment config
├── db/
│   ├── mongo/          # Mongo connection
│   └── psql/           # Postgres connection + migrations
├── handlers/           # HTTP handlers
├── services/           # Kafka consumer/producer
├── routers/            # Routing logic
├── shared/             # Shared models & utils
├── Dockerfile
├── docker-compose.yml
└── .env
```

---

## API Endpoints

- `POST /accounts` — Create new account  
- `POST /accounts/{id}/transactions` — Submit deposit/withdraw  
- `GET /accounts/{id}` — Fetch account info  
- `GET /accounts/{id}/transactions` — Get account's transaction history  
- `GET /health` — Health check  

---

## Kafka Topics

- `transactions` — All deposit/withdraw requests  
- Producer sends to `localhost:9092` or `kafka:9092` in Docker  
- Consumer listens in background on app boot

---

## MongoDB Access

To open Mongo shell inside container:

```bash
docker exec -it mongodb mongosh -u admin -p password123 --authenticationDatabase admin
```

---

## Troubleshooting

- Make sure Kafka is addressed as `kafka:9092` **inside Docker**, not `localhost:9092`
- If DB migrations fail, check that folder `db/psql/migrations` exists with `.up.sql` and `.down.sql` files
- Clean up orphans with:  
  ```bash
  docker compose down --remove-orphans
  ```

---

## License

MIT License © 2025 likhithkp
