# nouveau-microservices-go

"New Venture" is our current research or brain storming topic , so we are going to build a new software with cutting edge technology and cloud native.
In this project, I have built backend platform by using microservice architecture with REST APIs where Postgres database is used as persistence data.
CRUD operations/ routes are implemented as middleware "handler". Model package contain mainly three fields: Name, Domain and Revenue.
All the SQL queries are running in the handler.go for Create, Retrieve, Update and Delete Venture.

In Golang, There are several libraries are imported: http requests are handled through "gorilla Mux" , and env file is loaded though "godotenv" library.

This project is intended to write through Golang.
The APIs have been tested through Postman.


Next step, I'll build front end app with Angular framework. And, later everything will be running in cloud.


