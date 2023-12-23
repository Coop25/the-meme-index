# Go File Storage with PostgreSQL

This is a simple Go application that allows you to store files in a PostgreSQL database. The application provides basic functionalities to upload files, retrieve files by ID, retrieve a random file, and search files by tags.

## Prerequisites

Before running the application, make sure you have the following installed:

- Go (https://golang.org/doc/install)
- PostgreSQL

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/your-repository.git
    cd your-repository
    ```

2. Set up your PostgreSQL database and update the connection string in `main.go`:

    ```go
    const (
        dbDriver = "postgres"
        dbSource = "user=username dbname=mydatabase sslmode=disable" // replace with your PostgreSQL connection string
    )
    ```

3. Run the application:

    ```bash
    go run client/main.go
    ```

The server will start on port 8080 by default. You can customize the port in `main.go`.

## API Endpoints

- `POST /upload`: Upload a file to the database.
- `GET /files/{id}`: Retrieve a file by ID.
- `GET /random`: Retrieve a random file.
- `GET /search?tags=tag1,tag2`: Search files by tags.

## Usage Examples

### Upload a File

```bash
curl -X POST -F "file=@/path/to/your/file" -F "tags=tag1,tag2" http://localhost:8080/upload
