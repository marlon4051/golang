# Golang Task Management API

## Getting Started

### Prerequisites

- **Go**: Version 1.23.2  
  Download from: [Go Downloads](https://go.dev/dl/)

- **MySQL**: Version 8.0.40  
  Download from: [MySQL Installer](https://dev.mysql.com/downloads/installer/)

## Cloning Repository
Follow these steps to set up and run the project locally:

Clone the repository:

```sh
git clone https://github.com/marlon4051/task-api
```
Navigate to the project directory:
```sh
cd task-api
```

### Installation Steps

1. **Install Go**: Follow the instructions on the Go website to install Go on your machine.

2. **Install MySQL**: Follow the instructions on the MySQL website to install MySQL.

3. **Database Setup**:
   After installing MySQL, you need to run a script to create the database and the necessary tables. Make sure to execute the provided SQL script for setting up the `task_manager_db`. The scrip is the next one:
   
    ```sh
    CREATE DATABASE task_manager_db;

    USE task_manager_db;

    CREATE TABLE users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        userName VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE TABLE tasks (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT NOT NULL,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        status ENUM('pending', 'working', 'completed') DEFAULT 'pending',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE -- Delete all task if we delete user
    );
    ```

4. **Fetch Dependencies**:  
    After setting up the database, navigate to the project directory and run:
   ```sh
   go mod tidy
   ```
5. **Configure Environment Variables**:
   Verify the .env file in the project root. You need to provide the necessary database connection details. The file should look like this:
    ```sh
    DB_USERNAME=
    DB_PASSWORD=
    DB_HOST=localhost
    DB_PORT=
    DB_NAME=task_manager_db
    JWT_SECRET=
    ```

### Running the Application
To start the application, use the following command:

```sh
go run main.go
```

### Running the Frontend
To use the complete application, ensure the frontend is running simultaneously. You can find instructions to run the frontend in the following repository: Angular Task UI.
[Angular Task UI](https://github.com/marlon4051/task-angular-ui)

With both the backend and frontend running at the same time, you will be able to use the application locally.

### Features
- User Authentication
- CRUD operations for Tasks
- JWT for secure access