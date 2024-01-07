# OpenPachiSlot-Server

OpenPachiSlot-Server is an advanced software solution, designed for efficient management and tracking of items. It integrates a Go-based web server with a PostgreSQL database and a Vue.js front-end, offering a comprehensive RESTful API service. The project leverages Docker for simplified environment setup, the Gin framework for efficient HTTP request handling, and Vue.js for a responsive user interface.

## Key Features

- **Robust Item Tracking:** Manage and track items efficiently via a RESTful API.
- **Modern Technology Stack:** Go, Vue.js, PostgreSQL, Docker, and Nginx.
- **Easy Setup:** Docker and Docker Compose for straightforward deployment.
- **Flexible API Endpoints:** Facilitate various operations like item registration, data retrieval, and count updates.

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

3. **Docker Compose:**
   Use `docker-compose.yml` to set up the environment:
   ```bash
   docker-compose up --build
   ```

4. **Access the Application:**
   - Vue.js frontend: `http://localhost`
   - Go server API: `http://localhost/api`

## API Documentation

- **POST /user:** Register a new user with a name and card ID.
- **PUT /user/:uuid:** Update user data based on UUID.
- **GET /user/:uuid:** Retrieve user data using UUID.
- **GET /all-items:** Get a list of all items.

## Development Notes

- The server uses `air` for auto-reloading on code changes.
- Database tables are created automatically if they don't exist.
- The project is in development; not all features are fully implemented or tested.

## Contributing

Contributions, suggestions, and feedback are welcomed as this project is a work-in-progress.