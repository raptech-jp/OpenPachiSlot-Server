# OpenPachiSlot-Server

OpenPachiSlot-Server is an in-development software solution integrating a Go-based web server, a Vue.js front-end, and a PostgreSQL database. It's designed for managing and tracking items using a RESTful API, employing Docker for environment setup, the Gin framework for HTTP requests, and Vue.js for a dynamic user interface.

## Requirements

- Docker and Docker Compose
- Go (for development)
- Node.js and npm (for Vue.js development)


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
    VUE_PORT=8081
    ```

3. **Docker Compose Setup:**

    The `docker-compose.yml` file configures three services:

    - `vue-app`: A Vue.js application for the front-end interface.
    - `go`: The main Go server for the backend.
    - `postgres`: A PostgreSQL database.

    To start all
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