# Task Manager API Documentation

## Overview

The Task Manager API is a RESTful service that allows you to manage tasks through a simple HTTP interface. It provides endpoints for creating, reading, updating, and deleting tasks (CRUD operations). The API is built using Go and the Gin web framework.

## Task Model

A Task in the system consists of the following fields:

| Field       | Type   | Description                                     |
|-------------|--------|-------------------------------------------------|
| id          | string | Unique identifier for the task                  |
| title       | string | Short title or name of the task                |
| description | string | Detailed description of what the task involves  |

## Pre-seeded Data

The API comes with the following pre-seeded task:

```json
{
    "id": "1",
    "title": "Complete Task 5",
    "description": "Complete Task Manager API and push work to github"
}
```

## Base URL

```
http://localhost:8080
```

## Endpoints

### 1. Get All Tasks

Retrieves a list of all tasks.

**URL**: `/tasks`
**Method**: `GET`
**Response Format**:
  ```json
  [
    {
      "id": "string",
      "title": "string",
      "description": "string"
    }
  ]
  ```
**Example Response**:
  ```json
  [
    {
      "id": "1",
      "title": "Complete Task 5",
      "description": "Complete Task Manager API and push work to github"
    }
  ]
  ```

### 2. Get Task by ID

Retrieves a specific task by its ID.

**URL**: `/tasks/:id`
**Method**: `GET`
**URL Parameters**:  `id`

**Success Response**:
  ```json
  {
    "id": "string",
    "title": "string",
    "description": "string"
  }
  ```
**Error Response**:
  ```json
  {
    "message": "task Not Found"
  }
  ```

### 3. Add New Task

Creates a new task.

**URL**: `/tasks`
**Method**: `POST`
**Request Body**:
  ```json
  {
    "id": "string",
    "title": "string",
    "description": "string"
  }
  ```
**Success Response**:
  ```json
  {
    "message": "Task Added Successfully"
  }
  ```
**Error Responses**:
  ```json
  {
    "error": "invalid Request"
  }
  ```
  or
  ```json
  {
    "error": "task Already Exists"
  }
  ```

### 4. Update Task

Updates an existing task.

**URL**: `/tasks/:id`
**Method**: `PUT`
**URL Parameters**: `id`
  
**Request Body**:
  ```json
  {
    "title": "string",
    "description": "string"
  }
  ```
**Success Response**:
  ```json
  {
    "message": "Task Updated Successfully!"
  }
  ```
**Error Response**:
  ```json
  {
    "message": "task Not Found"
  }
  ```

### 5. Delete Task

Deletes a task by its ID.

**URL**: `/tasks/:id`
**Method**: `DELETE`
**URL Parameters**:  `id`

**Success Response**:
  ```json
  {
    "message": "Task Deleted Successfully!"
  }
  ```
**Error Response**:
  ```json
  {
    "error": "task Not Found"
  }
  ```