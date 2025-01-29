# Getting Started
The first step is to initalize a module and get the gin framework.

```go
go mod init iamsuteerth/golang-stuff/gin-pt
go get github.com/gin-gonic/gin
```

## Server Basics
- The `gin.Default()` has two middlewares, `Logger()` and `Recovery()`
    - The latter returns a 500 HTTP code.
- Creating our first endpoint
    ```go
    server.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Ok!",
		})
	})
    ```
- Next up is creating an entity, a service and a controller.
    - It is important to add JSON serialization while creating structs and can be done like so
        ```go
        type Video struct {
            Title       string `json:"title"`
            Description string `json:"description"`
            URL         string `json:"url"`
        }
        ```
    - The reason is that it corresponds the data we provide in JSON to the field it's associated with.
    ---
    - A general format is:
        - `Struct` for the object you're handling such as a video as found [here](/golang-stuff/gin-pt/entity/video.go)
        - `Service` to interact with the object you created as found [here](/golang-stuff/gin-pt//service/video-service.go)
        - `Controller` to manage the handles as found [here](/golang-stuff/gin-pt/controller/video-controller.go)
    - It is better to follow this format while creating a controller or a service.
        - Interface 
        - Private struct which will implement that interface
        - Constructor returning this struct's address
        - Methods for the struct where the receiver is passed by reference

## Middlewares: Logging, Auth and Debugging
- The middleware `Logger()` and `Recovery()` are there in the `Default()` but it's the same as:
    ```go
    server := gin.New()
    server.Use(gin.Recovery())
    server.Use(gin.Logger())
    ```
- Let's say we want to customize a middleware such as logger which can be found [here](/golang-stuff/gin-pt/middlewares/logger.go)
    - The return type expected for middlewares is `HandlerFunc`
    ---
- On the subject of logging, it can be done on `stdout` as well as a `file`
    ```go
    func setupLogOutput() {
        file, _ := os.Create("./logs/gin.log")
        gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
    }
    ```
    - To setup logging in both file and stdout.
    - We can also do it on one of the outputs.
    - We are basically telling gin's deafault writer what to do.
    ---
- A very basic authentication setup can be done using the `BasicAuth()` which again expects a `HandlerFunc`. An implementation can be found [here](/golang-stuff/gin-pt/middlewares/basic-auth.go)
- The next step is to add this to our server.
- The username and password are encoded in `Base64` in the headers.

    ---
- Last but not the least is debugging. We can set it up with [github.com/tpkeeper/gin-dump](github.com/tpkeeper/gin-dump)