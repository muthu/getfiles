**Note**: The below consist of topics I had to lookup/learn about during the implementation of this project.

### References:

- [reference1](https://rickyanto.com/understanding-go-standard-http-libraries-servemux-handler-handle-and-handlefunc/)
- [reference2](https://www.integralist.co.uk/posts/understanding-golangs-func-type/)
- [reference3](https://restfulapi.net/)
- [reference4](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types)

### ServeHTTP
This function is needed to handle http requests of any type. The function declaration looks like this: `ServeHTTP (ResponseWriter, *Request)`

### ServeMux

Think of a ServeMux or mux (short form) as a multiplexer or router. Based on the URL pattern it routes requests coming in to certain endpoints to its respective handlers function.
http.HandleFunc uses a default servemux which is created as part of the package implementation. 

### Handler

Handler is a interface that define ServeHTTP function that will be called by ServeMux to process http request and send http response.

```
type Handler interface {
    ServeHTTP (ResponseWriter, *Request)
}
```

### HandlerFunc

HandlerFunc is the actual function a user-defined type that implements a handler interface and it makes any normal function as an HTTP handler. It means if I have a normal function with two parameters ResponseWriter, *Request then that can be used as an handler to serve a request.

### HandleFunc

A wrapper for the `HandlerFunc`. Refer to reference2 for more information.

Say you create a new multiplexer : `mux := new http.NewServeMux()`.

To add a route to this multiplexer you need to follow this syntax : `mux.Handle("/endpoint", http.HandlerFunc(func))`. Note that this function should be of type `func f (w ResponseWriter, r *Request)`

Why?
The second argument of `mux.Handle` expects a Handler. Handler is nothing but a interface whose definition is given above. So we know that `http.HandlerFunc` implements the handler interface. 
`http.HandlerFunc` is defined in the package as:

```
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f (w, r)
}
```
So now to add a new route we can use : `mux.HandleFunc("/endpoint", functionName)`

