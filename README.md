# OpenPachiSlot-Server

OpenPachiSlot-Server is a robust software solution integrating a Go-based web server, a Vue.js front-end, and a PostgreSQL database. It's designed for efficient management and tracking of items via a RESTful API. The system employs Docker for streamlined environment setup, the Gin framework for handling HTTP requests, and Vue.js for a dynamic user interface.

## Requirements

- Docker and Docker Compose
- Go (for backend development)
- Node.js and npm (for front-end Vue.js development)

## Installation and Setup

1. **Clone the Repository:**

    Begin by cloning the project repository to your local machine.

2. **Environment Variables:**

    Configure environment variables in a `.env` file. Here's an example:

    ```
    POSTGRES_HOST=postgres
    POSTGRES_USER=YOUR_USERNAME
    POSTGRES_PASSWORD=Y0UR_PASSWORD
    POSTGRES_PORT=5432
    POSTGRES_DB=DATABASE_NAME
    ```

3. **Docker Compose Configuration:**

    The `docker-compose.yml` outlines four services:

    - `vue-app`: The Vue.js application for the front-end, accessible at `http://vue-app:8080`.
    - `go-app`: The Go server for the backend, set up at `http://go-app:8080`.
    - `postgres`: A PostgreSQL database.
    - `nginx`: An Nginx server acting as a reverse proxy, routing traffic to the appropriate service.

    To initiate these services, execute:
    ```
    docker-compose up --build
    ```

    **Nginx Configuration:**

    The Nginx server is configured via `nginx.conf` to route:
    
    - All `/` requests to the Vue.js application.
    - All `/api` requests to the Go server, stripping the `/api` prefix.

4. **Running the Application:**

    After running the Docker Compose command, the Vue.js app will be available on port 80 of the host machine, and the Go server will handle API requests at `/api`.

## API Endpoints

- **POST /register:** Registers a new item with a name and card ID.
- **POST /add:** Increments the count of an item based on its ID and card ID.
- **POST /data:** Retrieves item data using its card ID.

## Development Notes

- The server uses `air` for auto-reloading on code changes.
- Database tables are created automatically if they don't exist.
- The project is in development; not all features are fully implemented or tested.

## Contributing

Contributions, suggestions, and feedback are welcomed as this project is a work-in-progress.