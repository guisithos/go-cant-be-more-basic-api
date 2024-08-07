Fast approach to demonstrate how to quickly set up a basic API using Go, Gonic (Gin), PostgreSQL, and Docker. Can't be more basic.
Features

    Basic: The API provides essential CRUD (Create, Read, Update, Delete) operations for products, the core of a RESTful API.
    Fast Setup:  Docker, to be quickly and consistently. 
    Efficient: Speed and minimalist.
    Scalable: PostgreSQL, robust and obviously scalable. 

Requirements

    Docker
    Docker Compose

# Clone the Repository

> sh
> git clone https://github.com/gusiithos/go-cant-be-more-basic.git
> cd go-cant-be-more-basic-api

# Environment Variables

## Create a .env file in the project root and configure your environment variables:

env
```
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=db
DB_PORT=5432
```
# Docker Setup

Build and run the Docker containers:
> sh
> docker-compose up --build

The API will be accessible at http://localhost:3030.
