# Task Manager API Documentation

## POSTMAN API DOC
https://documenter.getpostman.com/view/49835401/2sB3WsNJk7

## Overview

The Task Manager API is a RESTful service that allows you to manage tasks through a simple HTTP interface. It provides endpoints for creating, reading, updating, and deleting tasks (CRUD operations). The service now persists data to MongoDB. The Go service uses a small data layer in `data/task_service.go` which expects a MongoDB client to be connected via `ConnectDb(dbUri string)` before handling requests.

## Task Model

A Task in the system consists of the following fields:

| Field       | Type       | Description                                     |
|-------------|------------|-------------------------------------------------|
| id          | string     | Unique identifier for the task                  |
| title       | string     | Short title or name of the task                 |
| description | string     | Detailed description of what the task involves  |
| due_date    | string     | The date by which the task should be completed (RFC3339 string) |
| status      | string     | The current status of the task (e.g., "pending", "completed") |

Notes:
- The service stores tasks in MongoDB in the `a2sv` database, `tasks` collection.
- `due_date` is serialized as an RFC3339 timestamp string in JSON (e.g. "2025-11-06T00:00:00Z").

## Pre-seeded Data / Seeding

This implementation does not include in-memory pre-seeded data. Data is persisted in MongoDB, and the database will retain tasks across service restarts. If you want example data you can seed the database by creating tasks via the API (examples below) or using a MongoDB client.

Example curl to add a task (seed):

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Task 1","description":"First task","due_date":"2025-11-06T00:00:00Z","status":"Pending"}'
```

## Base URL

```
http://localhost:8080
```

## Authentication

This API uses JWT-based authentication for protected endpoints. Two public endpoints are available to create users and obtain a token:

- `POST /register` — Register a new user (first registered user becomes `admin`, subsequent users are `regular`).
- `POST /login` — Login with username and password, returns a JWT token.

Login request example:

```json
{
  "username": "alice",
  "password": "secret"
}
```

Login success response:

```json
{
  "token": "<jwt_token_here>"
}
```

Use the returned token in the `Authorization` header for protected requests:

```
Authorization: Bearer <jwt_token_here>
```

Notes on roles and access:
- The first registered user is automatically assigned the `admin` role (see `data.AddUser`).
- Endpoints that modify data (`POST /tasks`, `PUT /tasks/:id`, `DELETE /tasks/:id`) require the authenticated user to have `role: "admin"`.
- Read endpoints (`GET /tasks`, `GET /tasks/:id`) require any valid JWT token.

The middleware stores claim values such as `username` and `role` in the token; handlers expect `username` (or `user_id`) where appropriate.

## Endpoints

### 1. Get All Tasks

Retrieves a list of all tasks.

**URL**: `/tasks`
**Method**: `GET`
**Auth**: Required (any authenticated user)
**Response Format**:
  ```json
  [
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "due_date": "string",
      "status": "string"
    }
  ]
  ```
**Example Response**:
  ```json
  [
    {
      "id": "1",
      "title": "Complete Task 5",
      "description": "Complete Task Manager API and push work to github",
      "due_date": "2025-11-06T00:00:00Z",
      "status": "completed"
    }
  ]
  ```

### 2. Get Task by ID

Retrieves a specific task by its ID.

**URL**: `/tasks/:id`
**Method**: `GET`
**URL Parameters**:  `id`
**Auth**: Required (any authenticated user)

**Success Response**:
  ```json
  {
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "string",
    "status": "string"
  }
  ```
**Error Response** (when not found):
  ```json
  {
    "message": "task Not found"
  }
  ```

### 3. Add New Task

Creates a new task.

**URL**: `/tasks`
**Method**: `POST`
**Auth**: Required (admin only)
**Request Body**:
  ```json
  {
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "string",
    "status": "string"
  }
  ```
**Success Response**:
  ```json
  {
    "message": "Task Added Successfully"
  }
  ```
**Error Responses**:
  - Missing required fields (service error message): "invalid Request" (this originates from `data.AddTask`).
  - If a task with the same id already exists the service will return an error (handler-specific message).

### 4. Update Task

Updates an existing task.

**URL**: `/tasks/:id`
**Method**: `PUT`
**URL Parameters**: `id`
**Auth**: Required (admin only)
  
**Request Body**:
  ```json
  {
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "string",
    "status": "string"
  }
  ```
**Success Response**:
  ```json
  {
    "message": "Task Updated Successfully!"
  }
  ```
**Error Responses**:
  - If the `id` in the URL does not match the `id` in the body: service returns "Invalid Request".
  - If no task matches the given id: service returns "task with given id not found".

### 5. Delete Task

Deletes a task by its ID.

**URL**: `/tasks/:id`
**Method**: `DELETE`
**URL Parameters**:  `id`
**Auth**: Required (admin only)

**Success Response**:
  ```json
  {
    "message": "Task Deleted Successfully!"
  }
  ```
**Error Response**:
  ```json
  {
    "error": "task with given id not found"
  }
  ```

## Notes on behavior and errors
- The messages quoted above reflect strings returned by the data layer in `data/task_service.go`. HTTP handlers may wrap these into JSON responses with a `message` or `error` field; check the controller for exact HTTP status codes and body shape.
- Ensure `ConnectDb(dbUri)` is called at startup (see `docs/mongodb.md`) so the data layer can access the `tasks` collection.
  
# MongoDB Setup and Seeding

This project persists tasks in MongoDB. This document explains how to run MongoDB locally (or with Docker), set the connection URI, ensure the application connects at startup, and seed example tasks.

## Database and collection
- Database: `a2sv`
- Collection: `tasks`

The data layer (`data/task_service.go`) expects a connected MongoDB client and uses `client.Database("a2sv").Collection("tasks")`.

Set the DB URI your app uses. Example URIs:

- Local (no auth): `mongodb://localhost:27017`
- Atlas or remote (example): `mongodb+srv://<user>:<pass>@cluster0.example.mongodb.net`

**How to provide it to the app**

In main.go call ConnectDb from task_service and pass it your database URI like this:

```
if err := data.ConnectDb("mongodb://localhost:27017"); err != nil {
  fmt.Println("Could not connect to database")
  fmt.Println(err)
}
```

## Seeding example tasks

You can seed tasks either by using the API endpoints using curl:

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Task 1","description":"First task","due_date":"2025-11-06T00:00:00Z","status":"Pending"}'

curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":"2","title":"Task 2","description":"Second task","due_date":"2025-11-07T00:00:00Z","status":"In Progress"}'

curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":"3","title":"Task 3","description":"Third task","due_date":"2025-11-08T00:00:00Z","status":"Completed"}'
```