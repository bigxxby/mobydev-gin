# ozinshe-backend-app v1.0

This project is a RESTful movie streaming platform that allows users to browse, search, and watch movies. It also provides features for administrators to manage movies, categories, and user profiles and etc.

## Technologies Used

- Backend: GoLang (Gin framework)
- Database: PostgreSQL
- Authentication: JWT

## Running the Application

To run the application, execute the following commands in your terminal:

```bash
docker-compose build
docker-compose up

    OR

cd .
go run ./server
```

## Project Structure

- **Internal**: Internal components of the application.
  - **Database**: Files for interacting with the database.
  - **Utils**: Utility functions such as authentication, encryption, random number generation, and validation.

## Packages

- **pkg/middleware**: Middleware components such as authentication.
- **pkg/routes**: API routes divided by application entities like age, categories, episodes, etc.

## HTTP Handlers

HTTP request handlers for various entities (e.g., age, categories, episodes, etc.) are located in the `pkg/routes/api` package, organized by request methods (GET, POST, PUT, DELETE) and corresponding entities.

## API Documentation

For detailed information on available endpoints and requests, please refer to the API swagger documentation at [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html) when launching the app.

## Admin Credentials

Use the following credentials to log in as an admin:

- **Email:** big@example.com
- **Password:** Aa12345678#

## Features in Progress

### Frontend

- **Status:** NOT DONE

### Unit-tests

- **Status:** NOT DONE

## General Notes:

You must include a valid authentication token in the request headers if its a route that require admin privilege or user auth.
Admin routes are indicated where only administrators have access. Admin token saved in local storage after login.
