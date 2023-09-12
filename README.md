## Prerequisites
### make sure you have the following installed:
`golang >= 1.20`

`mongodb`

## Run server locally
At the root of this repo is a file `sample.env` this contains all the necessary environment variables required to run this. 
rename `sample.env` to `app.env` and modify the default values of the variables (if you wish to) or leave it as is.

### run server
```bash
go run main.go
```

by default, it runs on port :8000

## Testing API Endpoints using `curl`

<b style="color:yellow">CREATE</b> a new user
```bash
# Create a new user
# localhost:8000/api
# request body should contain a field called `name` this is required.

# Example request:

curl -X POST 'localhost:8000/api' \
--header 'Content-Type: application/json' \
--data '{
    "name": "John Doe"
}'

# Example response:
{
  "data": {
    "id": "650084150bc493d01f85144d",
    "name": "John Doe",
    "created_at": "2023-09-12T15:30:29.398Z"
    "updated_at": "2023-09-12T15:30:29.398Z"
  },
  "message": "user creation successful",
  "status": 200
}
```


<b style="color:green">GET</b> user by id
```bash
# Retrieve user by id
# localhost:8000/api/:id
# `id` url parmeter refers to the user id.

# Example request:
curl -X GET 'localhost:8000/api/650084150bc493d01f85144d'

# Example response:
{
  "data": {
    "id": "650084150bc493d01f85144d",
    "name": "John Doe",
    "created_at": "2023-09-12T15:30:29.398Z"
    "updated_at": "2023-09-12T15:30:29.398Z"
  },
  "message": "user retrieval successful",
  "status": 200
}
```

<b style="color:blue">UPDATE</b> user by id
```bash
# Update user data by id
# localhost:8000/api/:id
# `id` url parmeter refers to the user id.

# Example request:
curl -X PUT 'localhost:8000/api/650084150bc493d01f85144d' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Johnny Doe"
}'

# Example response:
{
  "message": "user update successful",
  "status": 200
}
```

<b style="color:red">DELETE</b> user by id
```bash
# Delete user data by id
# localhost:8000/api/:id
# `id` url parmeter refers to the user id.

# Example request:
curl -X DELETE 'localhost:8000/api/650084150bc493d01f85144d'

# Example response:
{
  "message": "user delete successful",
  "status": 200
}
```
## ER DIAGRAM
![HNGx STAGE 2 API ER DIAGRAM](https://github.com/noornee/hngx_stage-2/assets/71889751/4669291c-f34c-4aac-85d7-cc80cd1034ac)


# TODO:
- [x] CRUD Endpoints
- [x] UML/ER DIAGRAM
- [x] Deploy on render
- [x] Postman test docs
