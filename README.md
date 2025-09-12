# Fitness Backend

## Overview

Fitness Backend is a Go-based application designed to manage fitness routines, exercises, sessions, and user data. It provides a robust API for handling CRUD operations on fitness-related data.

## Features

- User management
- Exercise tracking
- Routine creation and management
- Session logging
- Workout planning

## Project Structure

```
docker-compose.yaml
go.mod
go.sum
main.go
internal/
	config/
		database/
		handlers/
		models/
		routes/
```

## Getting Started

### Prerequisites

- Go 1.20+
- Docker

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/the-tunnster/fitness-backend.git
   ```
2. Navigate to the project directory:
   ```bash
   cd fitness-backend
   ```
3. Start the database container using Docker Compose:
   ```bash
   docker-compose up -d
   ```
4. Build the app executable:
	```bash
	go build
	```
5. Start the web-server:
	```bash
	./fitness-tracker
	```

## License

This project is licensed under the CC by NC License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or features you'd like to add.

## Contact

For inquiries, please contact the repository owner at `the-tunnster@example.com`.
