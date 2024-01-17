# OpenPachiSlot-Server

OpenPachiSlot-Server is an advanced software solution, designed for efficient management and tracking of items. It integrates a Go-based web server with a PostgreSQL database and a Vue.js front-end, offering a comprehensive RESTful API service. The project leverages Docker for simplified environment setup, the Gin framework for efficient HTTP request handling, and Vue.js for a responsive user interface.

## Key Features

- **Robust Item Tracking:** Manage and track items efficiently via a RESTful API.
- **Modern Technology Stack:** Go, Vue.js, PostgreSQL, Docker, and Nginx.
- **Easy Setup:** Docker and Docker Compose for straightforward deployment.
- **Flexible API Endpoints:** Facilitate various operations like item registration, data retrieval, and count updates.
- **Multilingual Frontend Support:** We welcome contributions for frontend multilingual support.

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
    ADMIN_ID=admin
    ADMIN_PASSWORD=Password01
    ```

3. **Docker Compose:**
   Use `docker-compose.yml` to set up the environment:
   ```bash
   docker-compose up --build
   ```

4. **Access the Application:**
   - Vue.js frontend: `http://localhost`
   - Go server API: `http://localhost/api`

## Authentication

The `/api/login` endpoint is used for authentication. Users should authenticate using their "id" and "password". Upon successful authentication, a JWT (JSON Web Token) is issued. This token must be included in the request header as `Authorization: token` for subsequent authenticated requests.

## API Endpoints

- `POST /user:**` (registerUser): Register a new user with a name and card ID.

Protected endpoints include:

- `PUT /user/:uuid` (updateUser): Update user data.
- `GET /user/:uuid` (getUserData): Retrieve data of a specific user.
- `GET /all-items` (getAllItems): Retrieve data of all items.

These endpoints require authentication using the JWT obtained from the `/api/login` endpoint.

## Development Notes

- The server uses `air` for auto-reloading on code changes.
- Database tables are created automatically if they don't exist.
- The project is in development; not all features are fully implemented or tested.

## Contributing

Contributions, suggestions, and feedback are welcomed as this project is a work-in-progress.