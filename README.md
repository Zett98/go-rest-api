# go-rest-api


simple rest api in go.

localhost:8080/users -- POST request requires email and password fields

localhost:8080/sessions -- POST request requires email and password fields

localhost:8080/private/whoami -- GET REQUEST -- returns info about user requires session cookie to be present

availible commands:

make build;./apiserver -- run the server

make test -- run tests
