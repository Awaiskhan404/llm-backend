```

llm-backend/
├── cmd/gin-server/          # Main entry point for the application
│   ├── main.go              # Initializes and runs the server
│
├── configs/                 # Configuration files (environment variables)
│   ├── .env                 # Main environment file
│   ├── .env.example         # Example environment file for reference
│
├── internal/
│   ├── bootstrap/           # Application bootstrap (dependency injection)
│   │   ├── bootstrap.go
│   │   ├── routes.go        # Centralized route registration
│   │
│   ├── middlewares/         # Middleware implementations
│       ├── auth_middleware.go
│       ├── cors_middleware.go
│       ├── errors_middleware.go
│       ├── middlewares.go   # Middleware initializer
│
├── pkg/                     # Core application business logic
│   ├── auth/                # Authentication module
│   │   ├── auth.go
│   │   ├── auth_controller.go
│   │   ├── auth_routes.go
│   │   ├── auth_service.go
│   │
│   ├── users/               # Users module
│   │   ├── users.go
│   │   ├── users_controller.go
│   │   ├── users_routes.go
│   │   ├── users_service.go
│   │   ├── users_repository.go
│   │
│   ├── vector/              # Vector module (Generated via `make pkg name=vector`)
│   │   ├── vector.go
│   │   ├── vector_controller.go
│   │   ├── vector_routes.go
│   │   ├── vector_service.go
│   │   ├── vector_repository.go
│   │
│   ├── common/              # Common utilities
│   │   ├── hasher.go
│   │   ├── validation.go
│   │
│   ├── interfaces/          # Interface definitions
│   │   ├── auth_service_interface.go
│   │
│   ├── lib/                 # Core libraries (Logger, Router, Database)
│       ├── database.go
│       ├── lib.go
│       ├── logger.go
│       ├── router.go
│
├── scripts/templates/       # Templates for `make pkg` command
│   ├── controller.template
│   ├── model.template
│   ├── module.template
│   ├── repository.template
│   ├── routes.template
│   ├── service.template
│
├── sql/                     # Database migration scripts
│   ├── create_query_functions.sql
│   ├── create_users_table.sql
│
├── test/                    # Testing files
│   ├── integration/         # Integration tests
│   │   ├── signup_test.go
│   │
│   ├── mocks/               # Mock files for testing
│       ├── auth_service_mock.go
│       ├── logger_mock.go
│       ├── users_service_mock.go
│
├── .gitignore               # Git ignore file
├── go.mod                   # Go modules file
├── go.sum                   # Go modules checksum file
├── LICENSE                  # License file
├── Makefile                 # Automation commands
└── README.md                # Documentation (This file)

```