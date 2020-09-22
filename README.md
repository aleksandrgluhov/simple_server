# Simple web server
## Source code structure:
  * __main.go__ - main package file. Contains main fuction and binds all the fuctionality toghether.
  * __server.go__ - Simple web server implementation.
  * __types.go__ - custom data types and structures declaration file
  * __config.go__ - application configuration management.
  * __util.go__ - common procedures and functions.

## Build:
```bash
make deps
make build
```

## Run
```bash
./simple_server
```

## Docker build
```bash
make deps
make build-prod
docker build -t agluhov/simple_server:latest .
docker run -it --rm --name simple_server -p 8080:8080 agluhov/simple_server:latest
```
