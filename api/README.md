# API Application

This repository contrains the source code for an API application built in Go.

## Folder Structure
- `cmd`:  Contains the main application entry point and any related CLI code.
    - `main.go`: The entry point of the API application.
- `pkg`: Holds the reusable and shareable packages or libraries of this application.
    - `handlers`: Contains the HTTP request handlers responsible for processing incoming requests and generating responses.
    - `middleware`: Contains the middleware code for handling authentication, logging, and other common functionalities.
    - `models`: Defines the data models or structs used by your application.
    - `repositories`: Includes the database setup and data access code.
    - `routers`: Contains the code for setting up the router and defining the API routes.
- `config`: Holds the configuration-related code or files.
- `go.mod`: The Go module file that defines the module and its dependencies.
