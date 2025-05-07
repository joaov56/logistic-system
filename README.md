# Logistics System

A simple logistics system built with Go using Domain-Driven Design (DDD) principles.

## Architecture

The system follows a clean architecture approach with the following layers:

- Domain Layer: Contains the core business logic and entities
- Application Layer: Implements use cases and orchestrates the domain
- Infrastructure Layer: Handles external concerns (database, etc.)
- Interface Layer: Manages HTTP endpoints

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- PostgreSQL (if running locally)

## Getting Started

1. Clone the repository:

```bash
git clone <repository-url>
cd logistic-system
```

2. Install dependencies:

```bash
go mod download
```

3. Run with Docker Compose:

```bash
docker-compose up --build
```

The application will be available at `http://localhost:8080`

## API Endpoints

### Deliveries

- `POST /deliveries` - Create a new delivery
- `GET /deliveries/{id}` - Get a delivery by ID
- `PUT /deliveries/{id}/status` - Update delivery status
- `DELETE /deliveries/{id}` - Delete a delivery
- `GET /deliveries` - List all deliveries (with optional filtering)

### Example Requests

Create a delivery:

```bash
curl -X POST http://localhost:8080/deliveries \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": "123",
    "customer_id": "456",
    "address": "123 Main St"
  }'
```

Update delivery status:

```bash
curl -X PUT http://localhost:8080/deliveries/{id}/status \
  -H "Content-Type: application/json" \
  -d '{
    "status": "IN_TRANSIT"
  }'
```

## Development

### Running Tests

```bash
go test ./...
```

### Database Migrations

The database schema is automatically created when the application starts.

## License

MIT
