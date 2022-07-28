# Basic Go Website

A simple TODO app to learn GO http stuff.

To run the website type:
```
go run main.go
```

In the browser searchbar:
```
http://localhost:9091/todo
```

Port is specified in main.go:
```go
log.Fatal(http.ListenAndServe(":9091", mux))
```