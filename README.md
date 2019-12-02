# go-rest-api


## Golang simple rest api with authentication

## Endpoints

### localhost:8080/users
 POST request, requires email and password fields

### localhost:8080/sessions 
POST request, requires email and password fields

### localhost:8080/private/whoami
GET request, returns info about user requires session cookie to be present

## Availible commands:

### Start
```
 -> make build
 -> ./apiserver 
``` 
### Testing
```
-> make test 
```
