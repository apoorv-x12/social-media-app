# Social Media App Backend

## Overview

This repository contains the backend of a social media application, designed with a scalable and clean architecture. The application is built using Go and implements the repository pattern for efficient data management, ensuring separation of concerns and ease of maintenance. It also incorporates various production-level features such as rate limiting, authentication, and structured logging.

## Features

- **Scalable Architecture**: Designed to handle a growing user base and increased load.
- **Repository Pattern**: Facilitates the separation of data access logic from business logic, making the application easier to test and maintain.
- **Rate Limiting**: Protects the API from abuse by limiting the number of requests a user can make within a specified time frame.
- **Authentication**: Implements secure user authentication with JWT (JSON Web Tokens).
- **Structured Logging**: Provides clear, structured logs for better monitoring and debugging.
- **Clean Code Practices**: Follows best practices for Go development to maintain high code quality.

## Technologies Used

- **Go**: The primary programming language.
- **Redis**: Used for rate limiting and caching.
- **PostgreSQL**: The database used to store user data and application content.
- **Docker**: For containerization and ease of deployment.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.17 or higher)
- [Docker](https://www.docker.com/get-started)
- PostgreSQL (install locally or use a Docker container)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/apoorv-x12/social-media-app-backend.git
   cd social-media-app-backend
