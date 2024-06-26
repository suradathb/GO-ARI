# GO-ARI
-- Install
cmd : go install github.com/cosmtrek/air@latest

-- Documentation Structure Project
go-swagger-api/
├── .air.toml
├── main.go
├── controllers/
│   └── book_controller.go
├── models/
│   └── book.go
├── routes/
│   └── book_routes.go
└── docs/
    └── ... (Swagger generated docs)

# Go Swagger API Example

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Setup and Installation](#setup-and-installation)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Adding a New API Endpoint](#adding-a-new-api-endpoint)

## Prerequisites

- Go (version 1.16 or higher)
- Gin Framework
- Swaggo for Swagger integration

## Project Structure
swag init
air
