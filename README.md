# Archive API

Archive API is a RESTful API built with the Go programming language and the Echo framework. It provides functionality to manage a file archive, including saving files, searching within files, adding captions and descriptions, assigning categories, editing files, deleting files, sorting and filtering files, user authentication, generating public links, and more.

## Prerequisites

Before running the Archive API, ensure you have the following installed:

- Go (version X.X.X)
- Echo framework (version X.X.X)
- Database system: PostgreSQL

## Installation

Clone the repository:

```shell
git clone https://github.com/BaseMax/ArchiveAPIGo.git
```

Navigate to the project directory:

```shell
cd ArchiveAPIGo
```

Install dependencies:

```shell
go mod download
```

Set up the database:

- Create a new database in your chosen database system.
- Update the database connection details in `config/config.go` or it's better to be inside `.env`.

Build and run the application:

```shell
go run main.go
```

The Archive API should now be running on http://localhost:8080.

## API Endpoints

The following API endpoints are available:

- `POST /login`: User login with username and password.
- `POST /register`: User registration with username, email, and password.
- `GET /files`: Get a list of all files with optional sorting and filtering parameters.
- `GET /files/{id}`: Get details of a specific file by its ID.
- `POST /files`: Upload a new file with optional caption, description, and category assignment.
- `PUT /files/{id}`: Update the details of a specific file by its ID.
- `DELETE /files/{id}`: Delete a specific file by its ID.
- `GET /files/{id}/public-link`: Generate a public link for a specific file by its ID.
- `GET /categories`: Get a list of all categories.
- `GET /categories/{id}`: Get details of a specific category by its ID.
- `POST /categories`: Create a new category.
- `PUT /categories/{id}`: Update the details of a specific category by its ID.
- `DELETE /categories/{id}`: Delete a specific category by its ID.
- `GET /search?q={keyword}`: Search for files based on a keyword.
- `GET /users/{id}`: Get details of a specific user by their ID.
- `PUT /users/{id}`: Update the details of a specific user by their ID.
- `DELETE /users/{id}`: Delete a specific user by their ID.

## Authentication

Authentication is required for most API endpoints. To authenticate, include the Authorization header in your requests with a valid JWT token:

```
Authorization: Bearer {token}
```

To obtain a JWT token, make a POST request to /login with valid credentials.

## License

This project is licensed under the GPL-3.0 License.

**Authors:**

- Amirhosein
- Max Base

Copyright 2023, Max Base
