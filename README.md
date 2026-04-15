# Event Booking API 

A robust REST API developed in **Go (Golang)** for managing event bookings.

## Features
- **User Management:** Secure Sign-up and Login system.
- **Event CRUD:** Full capabilities to Create, Read, Update, and Delete events.
- **Registration System:** Users can register for events and cancel their registrations.
- **Security:** Protected routes using **JWT** (JSON Web Tokens) and password hashing with **Bcrypt**.
- **Middleware:** Custom authentication middleware to verify user identity.

## Tech Stack
- **Language:** Go (Golang)
- **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database:** SQLite (for local development)
- **Authentication:** JWT (JSON Web Tokens)
- **Security:** Bcrypt for password encryption


## Prerequisites
- Go 1.25.0 or higher installed

## Installation and Running
1. Clone the repository:
   ```bash
    git clone [https://github.com/artur-sres/golang-event-booking-api.git](https://github.com/artur-sres/golang-event-booking-api.git)

2. Navigate to the project directory:
   ```bash
    cd golang-event-booking-api

3. Install dependencies:
   ```bash
    go mod tidy

4. Run the Application:
   ```bash
    go run main.go

The API will be avaliable at http://localhost:8080

## Main Endpoints
| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| POST | `/signup` | Create a new user account | No |
| POST | `/login` | Authenticate and get a JWT token | No |
| GET | `/events` | List all available events | No |
| POST | `/events` | Create a new event | **Yes** |
| PUT | `/events/:id` | Update an existing event | **Yes** |
| DELETE | `/events/:id` | Delete an event | **Yes** |
| POST | `/events/:id/register` | Register for an event | **Yes** |
| DELETE | `/events/:id/register` | Cancel a registration | **Yes** |

