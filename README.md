Book Store API with GO
============================

> This Book Store allows you to get, create, update and delete books
> in a database

## About the repository
> This repository is designed as follows:

    bookStore
    ├── app
    │    ├── cmd
    │    │    └── main.go
    │    └── pkg
    │         ├── config
    │         │     └── app.go
    │         ├── controllers
    │         │     └── book.controller.go
    │         ├── models
    │         │     └── book.go
    │         ├── routes
    │         │     └── app.go
    │         └── utils
    │               └── app.go
    ├── docker-compose.yml
    └── README.md

# Start Project
Create a .env file with the next env variables.
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=postgres
DB_HOST=db

Then, run the next docker command:

 ```
 docker-compose up --build
 ```

## Endpoint Calls
> I'll show you how to make a call to an endpoint using curl.
> If you want to format the JSON response in your terminal, 
> you can use "jq" else just remove "| jq" from the curl.
> Here is how to install it in your computer.

- Linux: 
 ```
 sudo apt-get update && sudo apt-get install jq
 ```

- Mac:
 ```
 brew install jq
 ```

- Windows:
 ```
 Don't use that
 ```

### Home Endpoint
 ```
curl -X GET localhost:4000/
 ```

### Get All Books
 ```
curl -X GET localhost:4000/book/ | jq
 ```

### Create Book
 ```
curl -X POST -H "Content-Type: application/json" -d '{
"name": "Carl",
"author": "Sagan",
"publication": "The Universe"
}' localhost:4000/book/ | jq
 ```

### Get Book by ID
 ```
curl -X GET localhost:4000/book/1 | jq
curl -X GET localhost:4000/book/2 | jq
curl -X GET localhost:4000/book/3 | jq
 ```

### Delete Book by ID
 ```
curl -X DELETE localhost:4000/book/1 | jq
curl -X DELETE localhost:4000/book/2 | jq
curl -X DELETE localhost:4000/book/3 | jq
 ```

### Update Book by ID
 ```
curl -X PUT -H "Content-Type: application/json" -d '{
"name": "Carl",
"author": "Sagan",
"publication": "The Universe"
}' localhost:4000/book/3 | jq
 ```