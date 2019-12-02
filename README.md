# go-rest-api


simple rest api in go.

localhost:8080/users -- POST request requires email and password fields

localhost:8080/sessions -- POST request requires email and password fields

localhost:8080/private/whoami -- GET REQUEST -- returns info about user requires session cookie to be present


requirements go,postgres. 

to run the server type this in project folder - make build;./apiserver

to run tests - make test
