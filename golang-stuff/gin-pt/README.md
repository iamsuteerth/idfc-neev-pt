# Getting Started with Gin Framework
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
        - `Controller` to manage the handles as found [here](/golang-stuff/gin-pt/controllers/video-controller.go)
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

## Data Binding and Validation
- **Never trust user input, Never trust user input and Never trust user input!!!**
- One of the Bind functions we've taken a look at is the `BindJSON()` and we can bind a lot of other things such as content, query, uri etc.
    - This is a shortcut for the `MustBindWith(obj, binding.JSON)`
    - In case of an error, it aborts with a 4xx status code and returns plaintext.
- For more flexibility to send the right headers with the content type that we want for the response, use `ShouldBindJSON()`
- Binding options in practice:
    ```go
    type Video struct {
        Title       string `json:"title" xml:"title" form:"title" "validate":"email" binding:"required"`
        Description string `json:"description"`
        URL         string `json:"url"`
    }
    ```
- Take a look at this complete section:
    ```go
    type Video struct {
        Title       string `json:"title" binding:"min=2,max=15"`
        Description string `json:"description" binding:"max=40"`
        URL         string `json:"url" binding:"required,url"`
        Author      Person `json:"author" binding:"required"`
    }
    ```
    - Here the title can be of max 15 characters and minimum 2
    - The URL is required and it's validated by telling that it's expected to be a url
    - You need to explictly validate the inputs by using `err = validate.Struct(video)` or something similar, binding does it automatically while it's trying to unmarshal the JSON.
    - Custom validators can also be used by specififying them in `validate:"is-cool"` and it's working can be found [here](/golang-stuff/gin-pt/validators/validators.go)

## HTML, Templates and Multi-Route Grouping
- Gin provides a way to use static files such as html and css. They can be served like so:
    ```go
    server.Static("/css", "./templates/css")
    server.LoadHTMLGlob("./templates/*.html")
    ```
    - A feature to organise our API endpoints is the ability to group a set of endpoints. So like the existing GET and POST methods are grouped into one and the remaining are then there can be a group for viewing these files.
        ```go
        apiRoutes := server.Group("/api")
        {
            apiRoutes.GET("/test", func(context *gin.Context) {
                context.JSON(200, gin.H{
                    "message": "Hello World!",
                })
            })
        }
        ```
    - A new function is made for viewing the videos which is made using the template html and css files.
- **NOTE** - `Context` . *The pointer \*gin. Context gives you access to the HTTP request as well as Gin's framework functions, such as String , which lets you write a string-type HTTP response. Use the router's method GET(path, handler) to associate the route / with the handler function homePage you defined above.*
- The variables can be passed on like so:
    ```go
    func (vc *videoController) ShowALl(context *gin.Context) {
        videos := vc.service.FindAll()
        // Store all the variables to be used within the templates
        data := gin.H{
            "title":  "Video Page",
            "videos": videos,
        }
        // Pass this to template
        context.HTML(http.StatusOK, "index.html", data)
    }
    ```
- The can then be accessed in the html files through `{{ .name_of_variable}}`
- The range keyword is used inside the html to loop through the videos like so `{{range .videos}}`, here the videos is the name of the variable we passed while making the map in the `ShowAll()` function.