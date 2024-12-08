# Go Authentication API

A robust authentication API built with Go, featuring JWT-based authentication, user management, and PostgreSQL database integration.

## Features

- User registration and login
- JWT-based authentication with access and refresh tokens
- PostgreSQL database integration using GORM
- Environment-based configuration
- Secure password hashing
- Clean architecture and modular design

## Tech Stack

- [Go](https://golang.org/) - Programming language
- [Gin](https://gin-gonic.com/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [PostgreSQL](https://www.postgresql.org/) - Database
- [JWT-Go](https://github.com/golang-jwt/jwt) - JWT implementation

## Getting Started

### Prerequisites

- Go 1.19 or higher
- PostgreSQL
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-auth.git
cd go-auth
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
```
Edit the `.env` file with your configuration.

### Running the Application

```bash
go run main.go
```

The server will start on the configured port (default: 3700).

## API Documentation

Complete API documentation is available on Postman:
[View API Documentation](https://documenter.getpostman.com/view/40197379/2sAYBd6T17)

### Quick API Reference

- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login user
- `POST /api/v1/auth/refresh` - Refresh access token

For detailed request/response examples and testing, please refer to the Postman documentation.

## Environment Variables

- `PORT` - Server port (default: 3700)
- `DB_URI` - PostgreSQL connection string
- `JWT_SECRET` - Secret key for JWT access tokens
- `REFRESH_SECRET` - Secret key for JWT refresh tokens

## Project Structure

```
go-auth/
├── config/             # Configuration files
├── database/          # Database migrations and setup
├── src/
│   ├── controllers/   # Request handlers
│   ├── models/        # Data models
│   ├── routes/        # API routes
│   └── utils/         # Utility functions
├── .env.example       # Example environment variables
├── .gitignore
├── go.mod
├── go.sum
├── main.go           # Application entry point
└── README.md
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
