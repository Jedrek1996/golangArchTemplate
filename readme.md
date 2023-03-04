Clean Architecture is a design philosophy for software development that emphasizes separation of concerns and modularity. It aims to make software more maintainable, testable, and scalable. In Go, clean architecture can be implemented using a combination of packages and interfaces. Here is an example implementation of clean architecture in Go:


First, let's define the layers of our application:

Controller layer: handles HTTP requests and responses.
Service layer: contains the business logic of the application.
Repository layer: contains the data access logic.


cmd/
  template/
    main.go
internal/
  controller/
    user_controller.go
  data/
    db.go
    repository.go
  interface/
    user_interface.go
  model/
    models.go
  service/
    user_service.go


1. Data Layer: The data layer contains code related to data storage and retrieval. It includes a database connection, data models, and a repository interface and implementation.

internal/data/db.go
internal/data/user_repository.go

2. Service Layer: The service layer contains business logic related to the application. It includes a service interface and implementation.

internal/service/user_service.go

3. Controller Layer: The controller layer contains code related to handling HTTP requests and responses. It includes a controller interface and implementation.

internal/controller/user_controller.go

This main function initializes the database connection, data layer, service layer, and controller layer, and sets up the HTTP routes using the Gorilla mux router. When a request comes in, the appropriate controller function is called, which in turn calls the appropriate service function, which may call the appropriate data layer function. The response flows back up through the layers in the opposite direction.


1. The main function in cmd/template/main.go creates a new instance of the Server struct and calls its ListenAndServe method to start the server.
2. The Server struct in internal/server/server.go has an instance of the Router struct and its ListenAndServe method sets up the router to handle incoming HTTP requests.
3. The Router struct in internal/router/router.go uses the http package's NewServeMux function to create a new request multiplexer and sets up routes for handling requests to the /users endpoint.
4. The UserController struct in internal/controller/user_controller.go has methods for handling incoming HTTP requests to the /users endpoint. These methods call methods on the UserService interface to retrieve or update user data.
5. The UserService interface in internal/service/user_service.go defines the methods that the UserController needs to interact with user data. It has dependencies on the UserRepository and Logger interfaces.
6. The UserServiceImpl struct in internal/service/user_service_impl.go implements the UserService interface and has an instance of the UserRepository interface. It uses the repository to interact with user data and logs events using the Logger interface.
7. The UserRepository interface in internal/data/user_repository.go defines the methods for interacting with the database to perform CRUD operations on user data.
8. The UserRepositoryImpl struct in internal/data/user_repository.go implements the UserRepository interface and has an instance of the DB struct, which represents a connection to the database. It uses this connection to interact with the database to perform CRUD operations on user data.
9. So the flow of the application goes from the main function to the Server, then to the Router, which directs incoming requests to the appropriate UserController method. The UserController then uses the UserService to interact with user data, which in turn uses the UserRepository to interact with the database.