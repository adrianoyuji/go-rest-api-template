# GO REST API Template
Golang API template with Gin, GORM (PostgresSQL), and JWT authentication.

## Quickstart
1. Copy `.env.example` to `.env` and fill values.
2. `go mod tidy`
3. `go run ./cmd/server`

This template includes example controllers, services, repositories, and middleware to jumpstart development.

## Features

- RESTful API structure using [Gin](https://github.com/gin-gonic/gin)
- Database integration with [GORM](https://gorm.io/) (PostgreSQL)
- JWT-based authentication and authorization
- Modular architecture: controllers, services, repositories, middleware
- Environment-based configuration
- Example CRUD endpoints

## Project Structure

```
.
├── cmd/
│   └── server/         # Entry point
├── config/             # Configuration loading
├── controllers/        # HTTP handlers
├── middleware/         # Custom middleware (e.g., JWT)
├── models/             # GORM models
├── repositories/       # Data access logic
├── routes/             # Route definitions
├── services/           # Business logic
├── utils/              # Utility functions
├── .env.example        # Example environment variables
└── README.md
```

## Environment Variables

Copy `.env.example` to `.env` and set the following:

- `PORT`
- `DB_URL`
- `JWT_SECRET`

## License

MIT