# OpenPachiSlot-Server

OpenPachiSlot-Server is an in-development software solution integrating a Go-based web server with a PostgreSQL database. It is designed to manage and track items using a RESTful API, leveraging the power of Docker for environment setup and Gin framework for handling HTTP requests.

## Requirements

- Docker and Docker Compose
- Go (for development)

## Installation and Setup

1. **Clone the Repository:**

    Clone the project repository to your local machine.

2. **Environment Variables:**

    The project utilizes environment variables for configuration, stored in a `.env` file. A sample configuration is:

    ```
    POSTGRES_HOST=postgres
    POSTGRES_USER=YOUR_USERNAME
    POSTGRES_PASSWORD=Y0UR_PASSWORD
    POSTGRES_PORT=5432
    POSTGRES_DB=DATABASE_NAME
    GO_PORT=8080
    ```

3. **Docker Compose:**

    The `docker-compose.yml` file is set up to create two services:

    - `go`: The main Go server.
    - `postgres`: A PostgreSQL database.

    To start the services, run:

    ```
    docker-compose up --build
    ```

## API Endpoints

- **POST /register:** Register a new item with a name and card ID.
- **POST /add:** Increment the count of an item based on its ID and card ID.
- **POST /data:** Retrieve item data using its card ID.

## Development Notes

- The server is configured to auto-reload on code changes using `air`.
- Database tables are automatically created if they do not exist.
- The project is still in development, and not all features are fully implemented or tested.

## Contributing

As this is a work-in-progress, contributions, suggestions, and feedback are highly appreciated.