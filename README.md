# Simple REST API

This is a very basic REST API built using Go. The API has three endpoints:

- `/register`: Used for registering a new user
- `/login`: Used for user authentication
- `/*`: Used for catching all other requests (returns 404)

# How to Run
```sh
git clone https://github.com/kolindes/simpleGoRestApi.git
```
```sh
cd simpleGoRestApi
```
```sh
go run cmd/main.go
```

## Usage

To use this API, send HTTP requests to the relevant endpoints.

### Register

To register a new user, send a POST request to `/register` with the following data:

- `username` (string)
- `password` (string)
- `email` (string)

```bash
curl --request POST \
  --url http://localhost:8080/register \
  --header 'Content-Type: application/json' \
  --data '{"username":"user","email":"myemail@example.com","password":"password"}'
```

If the registration is successful, the API will respond with a JSON object containing the following data:

- `access_token` (string)

### Login

To log in, send a POST request to `/login` with the following data:

- `username` (string)
- `password` (string)

If the login is successful, the API will respond with a JSON object containing the following data:

- `access_token` (string)

```bash
curl --request POST \
  --url http://localhost:8080/login \
  --header 'Content-Type: application/json' \
  --data '{
	"username": "user",
	"password":"password"
}'
```

If the login is unsuccessful, the API will respond with a 401 Unauthorized error.

## License

This project is licensed under the terms of the MIT license.
