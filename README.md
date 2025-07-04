# EGP Backend

A Go-based gRPC backend service for the EGP (Ebisu-GP - Ebisu Beer Grand Prix) system, built with Clean Architecture principles.

## Architecture

- **Model Layer**: Database models using GORM with PostgreSQL
- **Repository Layer**: Data access layer with interfaces  
- **UseCase Layer**: Business logic implementation
- **Router Layer**: gRPC server with Protocol Buffers

## Technologies

- Go 1.23.4
- PostgreSQL with PostGIS extensions
- GORM v2 ORM
- gRPC with Protocol Buffers
- Docker Compose for development
- Nginx reverse proxy

## Quick Start

```bash
# Start all services
make up

# Run database migrations
make db-migrate-up

# Test the API
make grpcurl-shops
```

## Development Commands

### Docker Environment
```bash
make up          # Start all services
make down        # Stop all services
make logs-api    # View API logs
make logs-db     # View database logs
make logs-web    # View nginx logs
```

### Database Operations
```bash
# Create new migration
make db-migrate-create name=your_migration_name

# Run migrations
make db-migrate-up

# Rollback migrations
make db-migrate-down
```

### Application
```bash
make run         # Run the application locally
go run main.go   # Alternative way to run
```

## gRPC Service

The service implements `EgpService` with endpoints for:
- `GetShops` - Retrieve shop listings
- `GetShop` - Get individual shop details
- `GetShopsTotal` - Get total shop count
- `AddStamp` - Add user stamp
- `DeleteStamp` - Remove user stamp

All endpoints require API key authentication via metadata.

## Environment Setup

Required environment variables (typically in `.env` file):
- `DB_USER`, `DB_PASS`, `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_SCHEMA`
- `API_PORT`, `API_KEY`
- `LOCATION_NAME` (timezone)

## Database Schema

The database uses PostgreSQL with a schema named `egp`. Key tables include:
- `events` - Event information
- `shops` - Shop details with location and time data
- `stamps` - User stamp collection
- `categories` - Shop categories
- `beer_types` - Beer type classifications
- `config` - Application configuration

## Ports

- gRPC API: 50051
- Nginx reverse proxy: 8080
