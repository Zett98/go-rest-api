# go-rest-api

## Golang simple rest api with authentication
incorporates user auth and session cookies

# How it works
```
.
├── cmd
│   └── apiserver        
│       └── main.go      
├── config
│    └──  apiServer.toml  
│
├──internal
│    └── app               
│        ├── apiserver
│        │     ├──  apiserver.go
│        │     ├──  config.go
│        │     ├──  responsewWriter.go
│        │     ├──  server_internal_test.go
│        │     └──  server.go
│        ├── model
│        │     ├──  testing.go
│        │     ├──  user_test.go
│        │     ├──  user.go
│        │     └──  validations.go
│        └── store
│             ├──  sqlstore
│             │       ├──  store_test.go
│             │       ├──  store.go
│             │       ├──  testing.go
│             │       ├──  userRepository_test.go
│             │       └──  userRepository.go
│             ├──  teststore
│             │       ├──  store.go
│             │       ├──  userrepository_test.go.go
│             │       └──  userrepository.go
│             ├──  errors.go
│             ├──  repository.go
│             └──  store.go        

```

# Getting started

## Install the Golang
https://golang.org/doc/install

# Endpoints

### localhost:8080/users
 POST request, requires email and password fields

### localhost:8080/sessions 
POST request, requires email and password fields

### localhost:8080/private/whoami
GET request, returns info about user requires session cookie to be present

# Availible commands:

### Start
```
 -> make build
 -> ./apiserver 
``` 
### Testing
```
-> make test 
```
