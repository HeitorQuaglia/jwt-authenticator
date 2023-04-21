# JWT Manager Microservice

This microservice is built with Go and provides two main functionalities: generating and validating JSON Web Tokens (JWTs) using the jwt-go library. The microservice exposes an HTTP API that allows clients to generate JWTs with a username and secret key, as well as validate existing JWTs.

## API endpoints

POST /generate-token: Generates a new JWT with a specified username and secret key. Returns the JWT as a JSON object.

POST /validate-token: Validates an existing JWT and returns a success message if the token is valid and the request contains a matching username.

## Project structure
jwt/generator.go: Contains the GenerateJWTToken function, which generates a new JWT with a specified username and secret key.
jwt/validator.go: Contains the ValidateJWTToken function, which validates an existing JWT and returns the username contained in the token if it is valid.
server/server.go: Contains the HTTP server implementation and the API endpoints for generating and validating JWTs.

## Dependencies
This microservice relies on the following third-party libraries:


## How to run
To run the microservice locally, you will need to have Go installed on your machine. Once Go is installed, you can clone the repository and run the following command in the root directory:

```
go run ./server/server.go
```
This will start the HTTP server on localhost:8000. You can then use a tool like curl or Postman to make requests to the API endpoints.

## How to test

To run the unit tests for the microservice, you can run the following command in the root directory:

```
go test ./...
```

This will run all the tests for the generator, validator, and server components.
