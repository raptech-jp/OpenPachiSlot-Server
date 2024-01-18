### Database-Related Functions

| Function Name         | Parameters   | Return Type   | Description |
|-----------------------|--------------|---------------|-------------|
| `initializeDatabase`  | None         | None          | Reads database connection information from environment variables and initializes the connection to the PostgreSQL database. On successful connection, it calls `createTable`. |
| `createTable`         | None         | None          | Creates the `items` table, which stores item IDs, names, counts, and card IDs. |

### Server and Routing Configuration

| Function Name   | Parameters   | Return Type    | Description |
|-----------------|--------------|----------------|-------------|
| `main`          | None         | None           | Initializes the database and sets up routing, then starts the server. |
| `setupRouter`   | None         | `*gin.Engine`  | Sets up HTTP routing using the `gin` framework. Defines endpoints for login, user registration, and item retrieval. |

### Authentication and Security

| Function Name     | Parameters       | Return Type         | Description |
|-------------------|------------------|---------------------|-------------|
| `login`           | `*gin.Context`   | `void`              | Processes login and issues JWT. |
| `authMiddleware`  | None             | `gin.HandlerFunc`   | Provides a middleware to validate JWTs. |

### User and Item Management

| Function Name    | Parameters       | Return Type | Description |
|------------------|------------------|-------------|-------------|
| `registerUser`   | `*gin.Context`   | `void`      | Registers a new user and saves it to the database. |
| `updateUser`     | `*gin.Context`   | `void`      | Updates information for a specified user. |
| `getUserData`    | `*gin.Context`   | `void`      | Retrieves item data based on a specified user ID. |
| `getAllItems`    | `*gin.Context`   | `void`      | Retrieves data for all items. |

### Other Functions

| Function Name    | Parameters   | Return Type         | Description |
|------------------|--------------|---------------------|-------------|
| `corsMiddleware` | None         | `gin.HandlerFunc`   | Provides middleware for CORS settings. |