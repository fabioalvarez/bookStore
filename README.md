# Connect to MySql from terminal
    docker exec -it <container_id_or_name> bash
    mysql -u root -p
    SHOW DATABASES;

# Clean Docker Cache
docker system prune -a

# Endpoint Calls

I'll show you how to make a call to an endpoint using curl.
If you want to format the JSON response in your terminal, 
you can use "jq" else just remove "| jq" from the curl.
Here is how to install it in your computer.

1. Linux: 
    ```
    sudo apt-get update && sudo apt-get install jq
    ```
   
2. Mac:
    ```
    brew install jq
    ```

3. Windows:
    ```
    Don't use that
    ```

### Home Endpoint
    curl -X GET localhost:4000/

### Get All Books
    curl -X GET localhost:4000/book/ | jq

### Create Book Endpoint
    curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Carl",
    "author": "Sagan",
    "publication": "The Universe"
    }' localhost:4000/book/ | jq

### Get Book by ID
    curl -X GET localhost:4000/book/1 | jq
    curl -X GET localhost:4000/book/2 | jq
    curl -X GET localhost:4000/book/3 | jq

### Delete Book by ID
    curl -X DELETE localhost:4000/book/1 | jq
    curl -X DELETE localhost:4000/book/2 | jq
    curl -X DELETE localhost:4000/book/3 | jq

### Update Book by ID
    curl -X PUT -H "Content-Type: application/json" -d '{
    "name": "Carl",
    "author": "Sagan",
    "publication": "The Universe"
    }' localhost:4000/book/3 | jq