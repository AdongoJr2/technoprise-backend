# Technoprise Blog Backend API

This repository contains the backend API application for a Technoprise blog application. It provides functionalities for managing
blog posts,
including creation, retrieval (with pagination and search), updating, and deletion.

## Table of Contents

* [Features](#features)
* [Tech Stack](#tech-stack)
* [Getting Started](#getting-started)
    * [1. Clone the Repository](#1-clone-the-repository)
    * [2. Environment Setup](#2-environment-setup)
    * [3. Install Go Dependencies](#3-install-go-dependencies)
    * [4. Install Ent ORM CLI](#4-install-ent-orm-cli)
    * [5. Generate Ent Code](#5-generate-ent-code)
    * [6. Run the Application](#6-run-the-application)
* [API Endpoints](#api-endpoints)
* [Contributing](#contributing)
* [License](#license)

## Features

* **Blog Post Management:** Create, retrieve (single by slug, all with pagination/search), update, and delete blog
  posts.
* **Unique Slugs:** Automatically generates unique, URL-friendly slugs for blog posts based on their titles.
* **Pagination:** Supports paginated retrieval of blog posts.
* **Search Functionality:** Allows searching blog posts by title or content.
* **RESTful API:** Provides a clean and well-structured API for frontend consumption.
* **Image Storage:** Blog Images are stored on the server at the moment. For production environment,
cloud storage services (e.g Amazon S3 or Google Cloud Storage can be considered)
* **Database Migrations:** Automatic schema creation/update using Ent ORM on application startup (suitable for
  development).

## Tech Stack

* **Go:** The primary programming language.
    * `go 1.24.4`
* **Echo:** High-performance, minimalist Go web framework.
    * `github.com/labstack/echo/v4 v4.13.4`
* **PostgreSQL:** Robust relational database for storing blog data.
    * `github.com/lib/pq v1.10.9` (PostgreSQL driver)
* **Ent ORM:** Type-safe ORM for Go, simplifying database interactions and schema management.
    * `entgo.io/ent v0.14.4`
* **Godotenv:** For loading environment variables from a `.env` file.
    * `github.com/joho/godotenv v1.5.1`

Before you begin, ensure you have the following installed:

* **Go:** [Download and Install Go](https://golang.org/doc/install) (version 1.18 or higher recommended, specifically
  `1.24.4` as per `go.mod`).
* **PostgreSQL:** [Download and Install PostgreSQL](https://www.postgresql.org/download/). Ensure your PostgreSQL server
  is running.

## Getting Started

Follow these steps to get the backend server up and running on your local machine.

### 1. Clone the Repository

```bash
git clone https://github.com/AdongoJr2/technoprise-backend.git
cd technoprise-backend
```

### 2. Environment Setup

Create a `.env` file in the root of the technoprise-backend directory. Populate the file with contents based on the [
`.env.example`](.env.example) file on the project's root directory. (Replace the values with your own)

### 3. Install Go Dependencies
Navigate to the project root directory and install the required Go modules:
```bash
go mod tidy
```

### 4. Install Ent ORM CLI
If you haven't already, install the Ent ORM command-line tool:
```bash
go install entgo.io/ent/cmd/ent@latest
```

### 5. Generate Ent Code
Ent ORM generates Go code based on your schema definitions. You need to run this command after any changes to your ent/schema files:
```bash
go generate ./ent
```

### 6. Run the Application

```bash
go run main.go
```
The server will start on the port specified in your `.env` file

## API Endpoints
All API endpoints are prefixed with `/api/v1`.
### Blog Posts
* `POST /api/v1/posts`
  * **Description**: Creates a new blog post
  * **Request body (multipart/form-data):**
  ```bash
  # example using curl
  
  curl -X POST http://localhost:1234/api/v1/posts \
  -F "title=My New Blog Post" \
  -F "excerpt=This is a brief summary of my new post." \
  -F "content=This is the full content of my new blog post, with more details." \
  -F "image=@/path/to/your/image.jpg"
  ```
  * **Response (JSON):** The created blog post object.
* `GET /api/v1/posts`
  * **Description**: Retrieves a list of blog posts with pagination and optional search.
  * **Query Parameters:**
    * `page` (optional, int): The page number to retrieve (default: 1).
    * `limit` (optional, int): The number of posts per page (default: 10).
    * `search` (optional, string): A keyword to search in `title` or `content` or `excerpt`
  * **Example:** `GET /api/v1/posts?page=1&limit=5&search=go`
  * **Response (JSON):** The blog post object.
* `GET /api/v1/posts/:slug`
  * **Description**: Retrieves a single blog post by its unique slug.
  * **Path Parameters:**
    * `slug` (string): The unique slug of the blog post.
  * **Example:** `GET /api/v1/posts/my-first-blog-post`
  * **Response (JSON):** The blog post object.

## Contributing

Feel free to fork the repository, make improvements, and submit pull requests.

## License

This project is licensed under the MIT License.
