# Go Gin Hexagonal Architecture Example

This project is an example of hexagonal architecture implementation in Golang based on a simple CRUD application. The application is built using the Gin web framework and demonstrates how to structure a Go project following the principles of hexagonal architecture.

## Project Structure

The project is organized into several layers, each with a specific responsibility. Below is a visual interpretation of the project structure:

```
├── cmd
│   └── main.go               # Entry point of the application
├── config
│   └── config.go             # Configuration loading and validation
├── internal
│   ├── domain                # Domain entities
│   │   └── album.go
│   ├── handlers              # HTTP handlers
│   │   ├── handlers.go
│   │   └── handlers_test.go
│   ├── infrastructure
│   │   └── persistence       # Database persistence layer
│   │       └── db.go
│   ├── repositories          # Repository interfaces and implementations
│   │   └── album_repository.go
│   ├── routers               # HTTP routers
│   │   └── routers.go
│   └── services              # Business logic services
│       └── album_service.go
├── pkg
│   └── logger                # Logger interface and implementation
│       └── logger.go
├── .env                      # Environment configuration file
├── go.mod                    # Go module file
└── go.sum                    # Go dependencies file
```

## Hexagonal Architecture

Hexagonal architecture, also known as ports and adapters architecture, is a design pattern that aims to decouple the core logic of an application from its external dependencies. This allows for easier testing, maintenance, and scalability.

### Layers

1. **Domain Layer**: Contains the core business logic and domain entities. In this project, the `domain` package contains the `Album` entity.
2. **Service Layer**: Implements the business logic and interacts with the repository layer. The `services` package contains the `AlbumService`.
3. **Repository Layer**: Provides an interface for data access and implements the persistence logic. The `repositories` package contains the `AlbumRepository`.
4. **Handler Layer**: Handles HTTP requests and maps them to service calls. The `handlers` package contains the `AlbumHandler`.
5. **Infrastructure Layer**: Contains the implementation details for external dependencies such as databases. The `infrastructure/persistence` package contains the database connection logic.
6. **Configuration Layer**: Manages application configuration. The `config` package contains the configuration loading and validation logic.
7. **Logger Layer**: Provides a logging interface and implementation. The `pkg/logger` package contains the logger implementation.

## Running the Application

1. Clone the repository:
    ```sh
    git clone https://github.com/szymonsitko/go_gin_hexagonal.git
    cd go_gin_hexagonal
    ```

2. Create a `.env` file in the root directory with the following content:
    ```
    DB_USER=root
    DB_PASSWORD=password
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=albums
    PORT=8080
    ```

3. Run the application:
    ```sh
    go run cmd/main.go --env-path .env
    ```

4. The application will start and listen on the port specified in the `.env` file.

## Testing

To run the tests, use the following scripts:
```sh
    ./test_unit.sh
```

```sh
    ./test_integration.sh
```

These testing scripts are using CLI interface for running unit and integration tests. This CLI utility has been compiled and attached to this repository
within cmd/test folder (main.go/main binary).

## Conclusion

This project demonstrates how to implement hexagonal architecture in a Go application using the Gin web framework. By following this architecture, the application is more modular, testable, and maintainable.

Feel free to explore the code and modify it to suit your needs. Contributions and feedback are welcome!