## How to Run the Project

### Setup
To set up the project, install the necessary dependencies for both the frontend and backend.

### Frontend
To install the frontend dependencies, use the following command:
```sh
npm install
```

To run the frontend, use the following command:
```sh
npm run start
```

### Backend
To install the backend dependencies, navigate to the backend directory and use the following commands:
```sh
cd backend
go get ./...
go mod download
```

To run the backend, use the following command:
```sh
go run .
```
### Database Configuration

To configure the database in `main.go`, follow these steps:

1. Open the `main.go` file located in the `backend` directory.
2. Add the necessary database connection logic. For example:
    ```go
    package main

    import (
        "database/sql"
        _ "github.com/lib/pq"
        "log"
    )

    func main() {
       	dsn := "host=localhost user=postgres password=password dbname=test port=5432 sslmode=disable"
    }
    ```

3. Replace `postgres` and `test` with your actual database username and database name.