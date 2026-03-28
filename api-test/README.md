# API Testing Guide 

This directory contains HTTP request files used to manually test the API endpoints during development.

## Recommended Extension
To execute these tests directly within VS Code, we recommend installing:
- **REST Client** (by Huachao Mao)

## How to Test
1. Ensure the main server is running: `go run main.go`.
2. Open any `.http` file in this folder.
3. Click the **"Send Request"** link that appears above the HTTP methods (GET, POST, etc.).

## Authentication & Tokens
Several routes (such as creating events or registrations) are protected and require a valid JWT token in the `Authorization` header.

**Suggested Test Workflow:**
1. Use `create-user.http` to register an account.
2. Use `login.http` to authenticate and receive a token.
3. Copy the generated token and paste it into the `Authorization` field of the other test files.

> **Note:** Tokens are valid for 2 hours.