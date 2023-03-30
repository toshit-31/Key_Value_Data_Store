# Key Value Data Store

Key value data store written in `Go`

## Run the server

- navigate to the root folder of the project in the terminal
- run `go get .` to install all dependencies
- run `go run .` to start the server
___

## Rest API

method: `POST`

endpoint: `/query`

body:
```json
{
    "query" : "GET abc"
}
```

Response Example:
```json
{
    "value": "123"
}
```

Error Response Example:
```json
{
    "error": "KeyNotFound"
}
```